package main

import (
	inmemory_fifo_cache "github.com/DmitryTelepnev/inmemory-fifo-cache"
	"github.com/DmitryTelepnev/inmemory-fifo-cache/inmemory/fifo"
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	cache := fifo.NewCache(5)

	putAsyncExample(cache)
}

func putAsyncExample(cache inmemory_fifo_cache.Cache) {
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()

			cache.PutAsync(strconv.Itoa(i%5), i)
		}(i)
	}

	wg.Wait()

	time.Sleep(1 * time.Second)

	log.Println(cache.GetAll(10), len(cache.GetAll(10)))
}
