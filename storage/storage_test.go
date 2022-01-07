package storage

import (
	"crypto/rand"
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

func TestSimple(t *testing.T) {
	table := "t1"
	key := "hello"
	value := []byte("world")

	err := Set(table, key, value)
	assert.Nil(t, err)

	cached, ok, err := Get(table, key)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, len(value), len(cached))
}

func TestBig(t *testing.T) {
	table := "t1"
	key := "0000000000000000000000000000000000000000000000000000000000000000"
	value := make([]byte, 10000)
	rand.Read(value)

	err := Set(table, key, value)
	assert.Nil(t, err)

	cached, ok, err := Get(table, key)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, len(value), len(cached))
}
