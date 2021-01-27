# InMemory FIFO cache

## Test it

```
make test
```

```
BenchmarkInmemory_Put_Into_One_Key
BenchmarkInmemory_Put_Into_One_Key-8             5137932               206 ns/op             114 B/op          1 allocs/op
BenchmarkInmemory_Put
BenchmarkInmemory_Put-8                          5325735               263 ns/op             131 B/op          1 allocs/op
BenchmarkInmemory_GetN_From_One_Key
BenchmarkInmemory_GetN_From_One_Key-8           18159153                65.4 ns/op             0 B/op          0 allocs/op
```
