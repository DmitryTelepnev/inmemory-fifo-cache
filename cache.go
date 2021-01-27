package inmemory_fifo_cache

//Implement cache (keep in mind - cache must be concurrent)
//Implementation must store N objects per key with FIFO eviction
type Cache interface {
	Put(key string, value interface{})
	PutAsync(key string, value interface{})
	//return n elements from key
	GetN(key string, n int) []interface{}
	//return all elements from every key
	GetAll(n int) []interface{}
}
