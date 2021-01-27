package fifo

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
)

func TestInmemory_Put_And_GetN(t *testing.T) {
	cache := NewCache(2)

	type A struct {
		B int
		C string
	}

	a1 := &A{
		B: 1,
		C: "key1",
	}
	a2 := &A{
		B: 2,
		C: "key2",
	}
	a3 := &A{
		B: 3,
		C: "key3",
	}

	cache.Put("key1", a1)
	cache.Put("key2", a2)
	cache.Put("key3", a3)

	assert.Nil(t, cache.GetN("key1", 1))
	assert.Equal(t, []interface{}{a2}, cache.GetN("key2", 1))
	assert.Equal(t, []interface{}{a3}, cache.GetN("key3", 1))

	cache.Put("key3", a1)
	cache.Put("key3", a2)

	assert.Equal(t, []interface{}{a1, a2}, cache.GetN("key3", 2))
	assert.Equal(t, []interface{}{a1, a2}, cache.GetN("key3", 99999999))
}

func TestInmemory_GetAll(t *testing.T) {
	cache := NewCache(2)

	type A struct {
		B int
		C string
	}

	a2 := &A{
		B: 2,
		C: "key2",
	}
	a3 := &A{
		B: 3,
		C: "key3",
	}

	cache.Put("key2", a2)
	cache.Put("key3", a2)
	cache.Put("key3", a3)

	assert.Subset(t, []interface{}{a2, a2, a3}, cache.GetAll(2))
}

func BenchmarkInmemory_Put_Into_One_Key(b *testing.B) {
	key := "key1"
	var val interface{} = 10

	cache := NewCache(2)

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Put(key, val)
		}
	})
}

func BenchmarkInmemory_Put(b *testing.B) {
	var val interface{} = 10

	cache := NewCache(2)

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Put(strconv.Itoa(rand.Intn(4)), val)
		}
	})
}

func BenchmarkInmemory_GetN_From_One_Key(b *testing.B) {
	key := "key1"
	var val interface{} = 10

	cache := NewCache(2)
	cache.Put(key, val)

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			v := cache.GetN(key, 1)
			_ = v
		}
	})
}
