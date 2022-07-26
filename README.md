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
Benchmarked with Go 1.18.4 on Linux via WSL 2.0 on Windows 11 using AMD Ryzen 7 Microsoft Surface (R) Edition, 2000 Mhz, 8 Core(s), 16 Logical Processor(s)


```
BenchmarkUnshardedSingleMutex/Get-16                             5882634               202.1 ns/op
BenchmarkUnshardedSingleMutex/Set-16                             1217817               957.8 ns/op
BenchmarkUnshardedSingleMutex/Mix-16                             3033954               365.4 ns/op
BenchmarkShardedMultiMutexMap/Get-16                             5544064               220.7 ns/op
BenchmarkShardedMultiMutexMap/Set-16                             1459060               821.0 ns/op
BenchmarkShardedMultiMutexMap/Mix-16                             4260982               310.5 ns/op
BenchmarkShardedMultiRWMutexMap/Get-16                           5651881               215.3 ns/op
BenchmarkShardedMultiRWMutexMap/Set-16                           1725718               856.0 ns/op
BenchmarkShardedMultiRWMutexMap/Mix-16                           4174183               287.5 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Get-16                 5642370               228.5 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Set-16                 1463014               685.6 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Mix-16                 4235272               297.1 ns/op
BenchmarkOrcamanLibrary/Get-16                                   5270935               235.8 ns/op
BenchmarkOrcamanLibrary/Set-16                                   1667658               970.9 ns/op
BenchmarkOrcamanLibrary/Mix-16                                   3427770               309.8 ns/op
BenchmarkFanLiaoLibrary/Get-16                                   4777159               284.4 ns/op
BenchmarkFanLiaoLibrary/Set-16                                   1213696               975.2 ns/op
BenchmarkFanLiaoLibrary/Mix-16                                   3996474               331.8 ns/op
BenchmarkTidwallLibrary/Get-16                                   5335886               237.5 ns/op
BenchmarkTidwallLibrary/Set-16                                   1741177               802.7 ns/op
BenchmarkTidwallLibrary/Mix-16                                   4004839               305.5 ns/op
BenchmarkSyncMap/Get-16                                          5635993               223.2 ns/op
BenchmarkSyncMap/Set-16                                          1000000              1623 ns/op
BenchmarkSyncMap/Mix-16                                          4149564               334.8 ns/op
```