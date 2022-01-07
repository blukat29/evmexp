package storage

import (
	"database/sql"
	"errors"
	"time"

	"github.com/eko/gocache/store"

	_ "github.com/mattn/go-sqlite3"
)

// https://github.com/eko/gocache/tree/v1.2.0#write-your-own-custom-store

type sqliteStore struct {
	db *sql.DB
}

func newSqliteStore(path string) (store.StoreInterface, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	q := `
CREATE TABLE IF NOT EXISTS data (
	key   TEXT PRIMARY KEY,
	value TEXT
)`
	if _, err := db.Exec(q); err != nil {
		return nil, err
	}

	return &sqliteStore{db: db}, nil
}

func (s *sqliteStore) Get(key interface{}) (interface{}, error) {
	keyStr, ok := key.(string)
	if !ok {
		return nil, errors.New("key must be string")
	}

	var value []byte
	rows := s.db.QueryRow("SELECT value FROM data WHERE key = ?", keyStr)
	switch err := rows.Scan(&value); err {
	case sql.ErrNoRows:
		return nil, errors.New("Value not found in SQLite store")
	case nil:
		return value, nil
	default:
		return nil, err
	}
}

func (s *sqliteStore) GetWithTTL(key interface{}) (interface{}, time.Duration, error) {
	value, err := s.Get(key)
	return value, 0, err
}

func (s *sqliteStore) Set(key, object interface{}, options *store.Options) error {
	keyStr, ok := key.(string)
	if !ok {
		return errors.New("key must be string")
	}
	valueBytes, ok := object.([]byte)
	if !ok {
		return errors.New("value must be []byte")
	}

	_, err := s.db.Exec("INSERT OR REPLACE INTO data VALUES (?, ?)", keyStr, valueBytes)
	return err
}

func (s *sqliteStore) Delete(key interface{}) error {
	keyStr, ok := key.(string)
	if !ok {
		return errors.New("key must be string")
	}
	_, err := s.db.Exec("DELETE FROM data WHERE key = ?", keyStr)
	return err
}

func (s *sqliteStore) Invalidate(options store.InvalidateOptions) error {
	return nil
}

func (s *sqliteStore) Clear() error {
	_, err := s.db.Exec("DELETE FROM data")
	return err
}

func (s *sqliteStore) GetType() string {
	return "sqlite"
}
