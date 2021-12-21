package storage

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if err := Init(); err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func assertGetEq(t *testing.T, b KvBackend, table, key string, value interface{}) {
	val, ok, err := b.Get(table, key)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, val, value)
}

func assertGetNo(t *testing.T, b KvBackend, table, key string) {
	val, ok, err := b.Get(table, key)
	assert.Nil(t, err)
	assert.False(t, ok)
	assert.Nil(t, val)
}

// Test that Set() does not fail.
func TestSet(t *testing.T) {
	m1 := NewMemBackend()
	m2 := NewMemBackend()
	ts := NewCascadeKvStorage(m1, m2)

	// Different table-key-value
	assert.Nil(t, ts.Set("t1", "k1", "v1"))
	assert.Nil(t, ts.Set("t2", "k2", "v2"))
	assert.Nil(t, ts.Set("t3", "k3", "v3"))

	// Same table
	assert.Nil(t, ts.Set("t4", "k4", "v4"))
	assert.Nil(t, ts.Set("t4", "k5", "v5"))
	assert.Nil(t, ts.Set("t4", "k6", "v6"))

	// Same key
	assert.Nil(t, ts.Set("t5", "k7", "v7"))
	assert.Nil(t, ts.Set("t5", "k7", "v8"))
	assert.Nil(t, ts.Set("t5", "k7", "v9"))
}

// Test that Set() writes to all backends
func TestSetCascade(t *testing.T) {
	m1 := NewMemBackend()
	m2 := NewMemBackend()
	ts := NewCascadeKvStorage(m1, m2)

	assert.Nil(t, ts.Set("t1", "k1", "v1"))
	assertGetEq(t, ts, "t1", "k1", "v1")
	assertGetEq(t, m1, "t1", "k1", "v1")
	assertGetEq(t, m2, "t1", "k1", "v1")

	assert.Nil(t, ts.Set("t1", "k1", "v2"))
	assertGetEq(t, ts, "t1", "k1", "v2")
	assertGetEq(t, m1, "t1", "k1", "v2")
	assertGetEq(t, m2, "t1", "k1", "v2")
}

// Test that Get() can retrieve from topmost level with the data
// and that Get() fills the upper cache
func TestGetCascade(t *testing.T) {
	m1 := NewMemBackend()
	m2 := NewMemBackend()
	ts := NewCascadeKvStorage(m1, m2)

	// [m1=nil, m2=val]
	assert.Nil(t, m2.Set("t1", "k1", "v1"))
	assertGetNo(t, m1, "t1", "k1")
	assertGetEq(t, m2, "t1", "k1", "v1")

	// ts.Get() will return value anyway
	assertGetEq(t, ts, "t1", "k1", "v1")

	// After CascadeKvStorage.Get(), all levels have data.
	assertGetEq(t, m1, "t1", "k1", "v1")
	assertGetEq(t, m2, "t1", "k1", "v1")
}
