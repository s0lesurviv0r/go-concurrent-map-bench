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
* github.com/cornelk/hashmap (IN PROGRESS)
* github.com/dustinxie/lockfree
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
Benchmarked with Go 1.19.3 on Ubuntu 22.04.1 LTS using AMD Ryzen 7 Microsoft Surface (R) Edition, 2000 Mhz, 8 Core(s), 16 Logical Processor(s)

```
BenchmarkUnshardedSingleMutex/Get-16                     2124543               674.6 ns/op
BenchmarkUnshardedSingleMutex/Set-16                      679298              1842 ns/op
BenchmarkUnshardedSingleMutex/Mix-16                     1373188               871.1 ns/op
BenchmarkUnshardedSingleRWMutex/Get-16                   3188382               328.2 ns/op
BenchmarkUnshardedSingleRWMutex/Set-16                   1000000              2006 ns/op
BenchmarkUnshardedSingleRWMutex/Mix-16                   1580175               780.6 ns/op
BenchmarkShardedMultiMutexMap/Get-16                     2033330               742.8 ns/op
BenchmarkShardedMultiMutexMap/Set-16                      567921              1974 ns/op
BenchmarkShardedMultiMutexMap/Mix-16                     1884939               693.2 ns/op
BenchmarkShardedMultiRWMutexMap/Get-16                   1635735               759.5 ns/op
BenchmarkShardedMultiRWMutexMap/Set-16                    682179              1663 ns/op
BenchmarkShardedMultiRWMutexMap/Mix-16                   1705411               782.6 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Get-16                 2054887               765.3 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Set-16                  988224              1571 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Mix-16                 1426081               814.1 ns/op
BenchmarkOrcamanLibrary/Get-16                                   1778530               832.0 ns/op
BenchmarkOrcamanLibrary/Set-16                                    585670              1929 ns/op
BenchmarkOrcamanLibrary/Mix-16                                   1363203               950.8 ns/op
BenchmarkFanLiaoLibrary/Get-16                                   1000000              1031 ns/op
BenchmarkFanLiaoLibrary/Set-16                                    533665              1901 ns/op
BenchmarkFanLiaoLibrary/Mix-16                                   1000000              1100 ns/op
BenchmarkTidwallLibrary/Get-16                                   1604966               856.9 ns/op
BenchmarkTidwallLibrary/Set-16                                    801894              1803 ns/op
BenchmarkTidwallLibrary/Mix-16                                   1257034               904.7 ns/op
BenchmarkDustinxieLibrary/Get-16                                 1325163               931.9 ns/op
BenchmarkDustinxieLibrary/Set-16                                  698587              1753 ns/op
BenchmarkDustinxieLibrary/Mix-16                                 1000000              1045 ns/op
BenchmarkSyncMap/Get-16                                          1772929               726.8 ns/op
BenchmarkSyncMap/Set-16                                           533874              2994 ns/op
BenchmarkSyncMap/Mix-16                                          1536703               764.5 ns/op
```