# Golang Concurrent Map Benchmarks
**Note: This benchmark, and these findings, haven't been reviewed by anyone else yet. Please take these results with a grain of salt. Feel free to create a Github issue if you have any observations or suggestions.**

### Overview
This repo has both correctness tests and benchmarks for a variety of different thread/goroutine safe maps.

Tests the following types of concurrent maps:
* Unsharded map with `sync.Mutex`
* Unsharded map with `sync.RWMutex`
* Sharded map with each shard having it's own `sync.Mutex`
* Sharded map with each shard having it's own `sync.RWMutex`
* Sharded map with that has a off-to-the side array of `sync.RWMutex` for the shards

Tests the following external libraries:
* github.com/fanliao/go-concurrentMap
* github.com/orcaman/concurrent-map
* github.com/tidwall/shardmap

Wherever possible, the number of shards for a map is set to 32.

### Run
* Test
    * `go test -v`
* Benchmark
    * `go test --bench=.`

### Results
```
BenchmarkUnshardedSingleMutex/Get-16                     2068736               591.3 ns/op
BenchmarkUnshardedSingleMutex/Set-16                      800686              1599 ns/op
BenchmarkUnshardedSingleMutex/Mix-16                     1423275               796.4 ns/op
BenchmarkShardedMultiMutexMap/Get-16                     1986720               652.0 ns/op
BenchmarkShardedMultiMutexMap/Set-16                      765225              1463 ns/op
BenchmarkShardedMultiMutexMap/Mix-16                     1637143               722.8 ns/op
BenchmarkShardedMultiRWMutexMap/Get-16                   1758175               624.4 ns/op
BenchmarkShardedMultiRWMutexMap/Set-16                    799548              1530 ns/op
BenchmarkShardedMultiRWMutexMap/Mix-16                   1597490               809.7 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Get-16                 1995408               623.1 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Set-16                  807332              1510 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Mix-16                 1615689               759.7 ns/op
BenchmarkOrcamanLibrary/Get-16                                   1842102               610.3 ns/op
BenchmarkOrcamanLibrary/Set-16                                    808147              1581 ns/op
BenchmarkOrcamanLibrary/Mix-16                                   1532247               805.8 ns/op
BenchmarkFanLiaoLibrary/Get-16                                   1539364               753.5 ns/op
BenchmarkFanLiaoLibrary/Set-16                                    740644              1780 ns/op
BenchmarkFanLiaoLibrary/Mix-16                                   1392890               851.8 ns/op
BenchmarkTidwallLibrary/Get-16                                   1774716               672.5 ns/op
BenchmarkTidwallLibrary/Set-16                                    884296              1609 ns/op
BenchmarkTidwallLibrary/Mix-16                                   1810134               820.8 ns/op
BenchmarkSyncMap/Get-16                                          1838002               562.9 ns/op
BenchmarkSyncMap/Set-16                                           672801              2267 ns/op
BenchmarkSyncMap/Mix-16                                          1760946               756.3 ns/op
```