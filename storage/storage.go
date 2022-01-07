package storage

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
)

var cm cache.CacheInterface
var rist *ristretto.Cache

func Init() error {
	rStore, err := newRistStore()
	if err != nil {
		return err
	}

	sqlitePath := os.Getenv("DB_PATH")
	if len(sqlitePath) == 0 {
		sqlitePath = "data.db"
	}
	sStore, err := newSqliteStore(sqlitePath)
	if err != nil {
		return err
	}
	fmt.Println(sStore.GetType())

	cm = cache.NewChain(
		cache.New(rStore),
		cache.New(sStore),
	)
	return nil
}

func newRistStore() (store.StoreInterface, error) {
	// https://github.com/dgraph-io/ristretto#example
	// https://github.com/dgraph-io/ristretto#config
	rCache, err := ristretto.NewCache(&ristretto.Config{
		// 10x the number of items you expect to keep in the cache when full.
		NumCounters: 1e7,
		// Here cost == len(value). Thus MaxCost is max memory size.
		MaxCost:     1 << 29, // 512 MB
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	rist = rCache

	rStore := store.NewRistretto(rCache, nil)
	return rStore, nil
}

func flush() {
	if rist != nil {
		// Must Wait(), otherwise subsequent Get() may fail.
		// https://github.com/eko/gocache/issues/47#issuecomment-648831087
		// https://github.com/dgraph-io/ristretto#example
		rist.Wait()
	}
}

func Set(table, key string, value []byte) error {
	ck := cacheKey(table, key)
	err := cm.Set(ck, value, &store.Options{Cost: int64(len(value))})
	flush()
	return err
}

func Get(table, key string) ([]byte, bool, error) {
	ck := cacheKey(table, key)
	value, err := cm.Get(ck)
	if err != nil {
		if isNotFoundError(err) {
			return nil, false, nil
		} else {
			return nil, false, err
		}
	}

	switch value.(type) {
	case []byte:
		return value.([]byte), true, nil
	default:
		return nil, false, errors.New("bad type")
	}
}

func Exists(table, key string) bool {
	_, ok, _ := Get(table, key)
	return ok
}

func cacheKey(table, key string) string {
	return table + ":" + key
}

func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	switch err.Error() {
	// https://github.com/eko/gocache/blob/v1.2.0/store/ristretto.go#L49
	case "Value not found in Ristretto store":
	case "Value not found in SQLite store":
		return true
	}
	return false
}
