package fifo

import (
	"sync"
)

type valueSlice []interface{}

type inmemory struct {
	sync.RWMutex

	capacity int
	channel  chan asyncPut
	storage  map[string]valueSlice
	fifoKeys []string
}

type asyncPut struct {
	key   string
	value interface{}
}

func (c *inmemory) Put(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	_, exists := c.storage[key]
	if !exists && len(c.fifoKeys) >= c.capacity {
		keyForEviction := c.fifoKeys[0]
		c.fifoKeys = c.fifoKeys[1:]

		delete(c.storage, keyForEviction)
		c.fifoKeys = append(c.fifoKeys, key)
	}

	if len(c.storage[key]) >= c.capacity {
		c.storage[key] = c.storage[key][1:]
	}

	c.storage[key] = append(c.storage[key], value)
}

func (c *inmemory) PutAsync(key string, value interface{}) {
	c.channel <- asyncPut{
		key:   key,
		value: value,
	}
}

//return n elements from key
func (c *inmemory) GetN(key string, n int) []interface{} {
	c.RLock()
	defer c.RUnlock()

	maxLen := len(c.storage[key])
	if n > maxLen {
		n = maxLen
	}
	return c.storage[key][:n]
}

//return all elements from every key
func (c *inmemory) GetAll(n int) []interface{} {
	c.RLock()
	defer c.RUnlock()

	var all []interface{}
	for _, value := range c.storage {
		l := len(value)
		if n > l {
			n = l
		}
		all = append(all, value[:n]...)
	}

	return all
}

func asyncProcessor(cache *inmemory) {
	for {
		asyncPut, isOpened := <-cache.channel
		if !isOpened {
			break
		}

		cache.Put(asyncPut.key, asyncPut.value)
	}
}

//NewCache return Cache implementation which stores n objects per key
func NewCache(n int) *inmemory {
	channel := make(chan asyncPut, n)
	storage := make(map[string]valueSlice, n)
	fifoKeys := make([]string, n)

	cache := &inmemory{
		capacity: n,
		channel:  channel,
		storage:  storage,
		fifoKeys: fifoKeys,
	}

	go asyncProcessor(cache)

	return cache
}
