package main

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const NUM_KEYS = 1_000_000

// testSetGet tests single thread correctness
func testSetGet(t *testing.T, m ConcurrentMap) {
	testTable := map[string]string{
		"test-key":   "test-value",
		"test-key-1": "test-value-1",
	}

	for key, value := range testTable {
		// Ensure key doesn't exist
		_, exists := m.Get(key)
		assert.False(t, exists)

		// Set key
		m.Set(key, value)

		// Ensure the key we just set exists
		// and is correct value
		val, exists := m.Get(key)
		assert.True(t, exists)
		assert.Equal(t, value, val)
	}
}

// testSetGet tests concurrent correctness, ensuring changes
// made by one goroutine are visible in another
func testConcurrentSetGet(t *testing.T, m ConcurrentMap) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Setter thread
	go func() {
		for i := 0; i < NUM_KEYS; i++ {
			m.Set(strconv.Itoa(i), strconv.Itoa(i))
		}
		wg.Done()
	}()

	// Getter thread. Delayed to ensure key has been set
	// already
	go func() {
		time.Sleep(10 * time.Millisecond)
		for i := 0; i < NUM_KEYS; i++ {
			val, exists := m.Get(strconv.Itoa(i))
			assert.True(t, exists)
			assert.Equal(t, strconv.Itoa(i), val)
		}
		wg.Done()
	}()

	wg.Wait()
}

func test(t *testing.T, m ConcurrentMap) {
	testSetGet(t, m)
	testConcurrentSetGet(t, m)
}

func TestUnshardedSingleMutex(t *testing.T) {
	test(t, NewUnshardedSingleMutexMap())
}

func TestUnshardedSingleRWMutex(t *testing.T) {
	test(t, NewUnshardedSingleRWMutexMap())
}

func TestShardedMultiMutexMap(t *testing.T) {
	test(t, NewShardedMultiMutexMap(32))
}

func TestShardedMultiRWMutexMap(t *testing.T) {
	test(t, NewShardedMultiRWMutexMap(32))
}

func TestShardedMultiSegragatedRWMutexMap(t *testing.T) {
	test(t, NewShardedMultiSegragatedRWMutexMap(32))
}

func TestOrcamanLibrary(t *testing.T) {
	test(t, NewOrcamanLibrary())
}

func TestFanliaoLibrary(t *testing.T) {
	test(t, NewFanliaoLibrary())
}

func TestTidwallLibrary(t *testing.T) {
	test(t, NewTidwallLibrary())
}

/*
func TestCornelkLibrary(t *testing.T) {
	test(t, NewCornelkLibrary())
}
*/

func TestDustinxieLibrary(t *testing.T) {
	test(t, NewDustinxieLibrary())
}

func TestSyncMap(t *testing.T) {
	test(t, NewSyncMap())
}

// benchmarkMapSets performs sets concurrently
func benchmarkMapSet(b *testing.B, m ConcurrentMap) {
	b.Run("Set", func(sb *testing.B) {
		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.Set(getRand(15), getRand(15))
			}
		})
	})
}

// benchmarkMapGet performs gets concurrently
func benchmarkMapGet(b *testing.B, m ConcurrentMap) {
	b.Run("Get", func(sb *testing.B) {
		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, _ = m.Get(getRand(15))
			}
		})
	})
}

// benchmarkMapMix tests alternating sets and gets in parallel.
// This should simulate a balanced set/get load
func benchmarkMapMix(b *testing.B, m ConcurrentMap) {
	b.Run("Mix", func(sb *testing.B) {
		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set := false
				if set {
					m.Set(getRand(15), getRand(15))
				} else {
					_, _ = m.Get(getRand(15))
				}
				set = !set
			}
		})
	})
}

func benchmark(b *testing.B, m ConcurrentMap) {
	b.ResetTimer()
	benchmarkMapGet(b, m)
	benchmarkMapSet(b, m)
	benchmarkMapMix(b, m)
}

func BenchmarkUnshardedSingleMutex(b *testing.B) {
	benchmark(b, NewUnshardedSingleMutexMap())
}

func BenchmarkUnshardedSingleRWMutex(b *testing.B) {
	benchmark(b, NewUnshardedSingleRWMutexMap())
}

func BenchmarkShardedMultiMutexMap(b *testing.B) {
	benchmark(b, NewShardedMultiMutexMap(32))
}

func BenchmarkShardedMultiRWMutexMap(b *testing.B) {
	benchmark(b, NewShardedMultiRWMutexMap(32))
}

func BenchmarkShardedMultiSegragatedRWMutexMap(b *testing.B) {
	benchmark(b, NewShardedMultiSegragatedRWMutexMap(32))
}

func BenchmarkOrcamanLibrary(b *testing.B) {
	benchmark(b, NewOrcamanLibrary())
}

func BenchmarkFanLiaoLibrary(b *testing.B) {
	benchmark(b, NewFanliaoLibrary())
}

func BenchmarkTidwallLibrary(b *testing.B) {
	benchmark(b, NewTidwallLibrary())
}

/*
func BenchmarkCornelkLibrary(b *testing.B) {
	benchmark(b, NewCornelkLibrary())
}
*/

func BenchmarkDustinxieLibrary(b *testing.B) {
	benchmark(b, NewDustinxieLibrary())
}

func BenchmarkSyncMap(b *testing.B) {
	benchmark(b, NewSyncMap())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func getRand(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
