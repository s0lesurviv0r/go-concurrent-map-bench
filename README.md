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
goos: linux
goarch: amd64
pkg: github.com/s0lesurviv0r/go-concurrent-map-bench
cpu: AMD Ryzen 7 Microsoft Surface (R) Edition
BenchmarkUnshardedSingleMutex/Get-16                    12281431               183.0 ns/op
BenchmarkUnshardedSingleMutex/Set-16                     1000000              1254 ns/op
BenchmarkUnshardedSingleMutex/Mix-16                     3029019               413.7 ns/op
BenchmarkUnshardedSingleRWMutex/Get-16                  33020790                35.65 ns/op
BenchmarkUnshardedSingleRWMutex/Set-16                   1000000              1162 ns/op
BenchmarkUnshardedSingleRWMutex/Mix-16                  27542838                43.36 ns/op
BenchmarkShardedMultiMutexMap/Get-16                    59081289                20.27 ns/op
BenchmarkShardedMultiMutexMap/Set-16                    11379960                93.48 ns/op
BenchmarkShardedMultiMutexMap/Mix-16                    23267588                47.41 ns/op
BenchmarkShardedMultiRWMutexMap/Get-16                  59250309                20.14 ns/op
BenchmarkShardedMultiRWMutexMap/Set-16                  11466465                94.43 ns/op
BenchmarkShardedMultiRWMutexMap/Mix-16                  31438360                35.26 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Get-16                62492812                19.86 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Set-16                11220654                99.37 ns/op
BenchmarkShardedMultiSegragatedRWMutexMap/Mix-16                33407506                31.77 ns/op
BenchmarkOrcamanLibrary/Get-16                                  57665552                20.80 ns/op
BenchmarkOrcamanLibrary/Set-16                                  11423337                94.70 ns/op
BenchmarkOrcamanLibrary/Mix-16                                  28998466                35.13 ns/op
BenchmarkFanLiaoLibrary/Get-16                                  52634571                22.25 ns/op
BenchmarkFanLiaoLibrary/Set-16                                   7468802               146.2 ns/op
BenchmarkFanLiaoLibrary/Mix-16                                  27911967                42.55 ns/op
BenchmarkTidwallLibrary/Get-16                                  60208870                20.19 ns/op
BenchmarkTidwallLibrary/Set-16                                  15534781                72.84 ns/op
BenchmarkTidwallLibrary/Mix-16                                  38839478                31.34 ns/op
BenchmarkDustinxieLibrary/Get-16                                37977000                32.64 ns/op
BenchmarkDustinxieLibrary/Set-16                                 5389324               206.1 ns/op
BenchmarkDustinxieLibrary/Mix-16                                 7166394               165.0 ns/op
BenchmarkSyncMap/Get-16                                         61196125                19.48 ns/op
BenchmarkSyncMap/Set-16                                          1000000              2306 ns/op
BenchmarkSyncMap/Mix-16                                         41027619                25.55 ns/op
```
