package storage

import "sync"

type tableKey struct {
	table string
	key   string
}

type MemBackend struct {
	m  map[tableKey]interface{}
	mu sync.RWMutex
}

func NewMemBackend() *MemBackend {
	b := &MemBackend{}
	b.m = make(map[tableKey]interface{})
	return b
}

func (b *MemBackend) Connect() error {
	return nil
}

func (b *MemBackend) Set(table, key string, value interface{}) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	tk := tableKey{table, key}
	b.m[tk] = value
	return nil
}

func (b *MemBackend) Get(table, key string) (interface{}, bool, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	tk := tableKey{table, key}
	value, ok := b.m[tk]
	return value, ok, nil
}
