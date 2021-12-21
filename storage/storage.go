package storage

var kvs *CascadeKvStorage

func Init() error {
	kvs = NewCascadeKvStorage(
		NewMemBackend(),
	)
	return kvs.Connect()
}

func Set(table, key string, value interface{}) error {
	return kvs.Set(table, key, value)
}

func Get(table, key string) (interface{}, bool, error) {
	return kvs.Get(table, key)
}

type KvBackend interface {
	Connect() error
	Set(table, key string, value interface{}) error
	Get(table, key string) (interface{}, bool, error)
}

type CascadeKvStorage struct {
	Backends []KvBackend
}

func NewCascadeKvStorage(backends ...KvBackend) *CascadeKvStorage {
	c := &CascadeKvStorage{
		Backends: backends,
	}
	return c
}

func (c *CascadeKvStorage) Connect() error {
	for _, b := range c.Backends {
		if err := b.Connect(); err != nil {
			return err
		}
	}
	return nil
}

func (c *CascadeKvStorage) Set(table, key string, value interface{}) error {
	for _, b := range c.Backends {
		if err := b.Set(table, key, value); err != nil {
			return err
		}
	}
	return nil
}

func (c *CascadeKvStorage) Get(table, key string) (interface{}, bool, error) {
	result, err := c.search(table, key)
	if err != nil {
		return "", false, err
	} else if result == nil {
		return "", false, nil
	}

	for i := 0; i < result.depth; i++ {
		if err := c.Backends[i].Set(table, key, result.value); err != nil {
			return result.value, true, err
		}
	}
	return result.value, true, err
}

type searchResult struct {
	value interface{}
	depth int
}

func (c *CascadeKvStorage) search(table, key string) (*searchResult, error) {
	for depth, b := range c.Backends {
		if value, ok, err := b.Get(table, key); err != nil {
			return nil, err
		} else if ok {
			return &searchResult{value: value, depth: depth}, nil
		}
	}
	return nil, nil
}
