package main

import (
	"sync"

	"github.com/fanliao/go-concurrentMap"
	"github.com/orcaman/concurrent-map"
	"github.com/tidwall/shardmap"
)

// ConcurrentMap is the interface all of our
// concurrently accessible maps should conform to
type ConcurrentMap interface {
	Get(string) (string, bool)
	Set(string, string)
}

type UnshardedSingleMutexMap struct {
	sync.Mutex
	m map[string]string
}

func (m *UnshardedSingleMutexMap) Get(key string) (string, bool) {
	m.Lock()
	val, exists := m.m[key]
	m.Unlock()
	return val, exists
}

func (m *UnshardedSingleMutexMap) Set(key, value string) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

func NewUnshardedSingleMutexMap() *UnshardedSingleMutexMap {
	return &UnshardedSingleMutexMap{
		m: make(map[string]string),
	}
}

type UnshardedSingleRWMutexMap struct {
	sync.RWMutex
	m map[string]string
}

func (m *UnshardedSingleRWMutexMap) Get(key string) (string, bool) {
	m.RLock()
	val, exists := m.m[key]
	m.RUnlock()
	return val, exists
}

func (m *UnshardedSingleRWMutexMap) Set(key, value string) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

func NewUnshardedSingleRWMutexMap() *UnshardedSingleRWMutexMap {
	return &UnshardedSingleRWMutexMap{
		m: make(map[string]string),
	}
}

type mutexShard struct {
	sync.Mutex
	m map[string]string
}

type ShardedMultiMutexMap struct {
	shards []*mutexShard
}

func (m *ShardedMultiMutexMap) Get(key string) (string, bool) {
	shardIdx := getShardIndex(key, len(m.shards))
	shard := m.shards[shardIdx]
	shard.Lock()
	val, exists := shard.m[key]
	shard.Unlock()
	return val, exists
}

func (m *ShardedMultiMutexMap) Set(key, value string) {
	shardIdx := getShardIndex(key, len(m.shards))
	shard := m.shards[shardIdx]
	shard.Lock()
	shard.m[key] = value
	shard.Unlock()
}

func NewShardedMultiMutexMap(shardCount int) *ShardedMultiMutexMap {
	shards := make([]*mutexShard, shardCount, shardCount)
	for i := 0; i < shardCount; i++ {
		shards[i] = &mutexShard{
			m: make(map[string]string),
		}
	}

	return &ShardedMultiMutexMap{
		shards: shards,
	}
}

type rwmutexShard struct {
	sync.RWMutex
	m map[string]string
}

type ShardedMultiRWMutexMap struct {
	shards []*rwmutexShard
}

func (m *ShardedMultiRWMutexMap) Get(key string) (string, bool) {
	shardIdx := getShardIndex(key, len(m.shards))
	shard := m.shards[shardIdx]
	shard.RLock()
	val, exists := shard.m[key]
	shard.RUnlock()
	return val, exists
}

func (m *ShardedMultiRWMutexMap) Set(key, value string) {
	shardIdx := getShardIndex(key, len(m.shards))
	shard := m.shards[shardIdx]
	shard.Lock()
	shard.m[key] = value
	shard.Unlock()
}

func NewShardedMultiRWMutexMap(shardCount int) *ShardedMultiRWMutexMap {
	shards := make([]*rwmutexShard, shardCount, shardCount)
	for i := 0; i < shardCount; i++ {
		shards[i] = &rwmutexShard{
			m: make(map[string]string),
		}
	}

	return &ShardedMultiRWMutexMap{
		shards: shards,
	}
}

// ShardedMultiSegregatedRWMutexMap is a sharded
// map that, unlike other types, keeps shard RW Mutexes
// segragated from the actual data shards
type ShardedMultiSegregatedRWMutexMap struct {
	shards []map[string]string
	mu     []sync.RWMutex
}

func (m *ShardedMultiSegregatedRWMutexMap) Get(key string) (string, bool) {
	shardIdx := getShardIndex(key, len(m.shards))
	shard := m.shards[shardIdx]
	m.mu[shardIdx].RLock()
	val, exists := shard[key]
	m.mu[shardIdx].RUnlock()
	return val, exists
}

func (m *ShardedMultiSegregatedRWMutexMap) Set(key, value string) {
	shardIdx := getShardIndex(key, len(m.shards))
	shard := m.shards[shardIdx]
	m.mu[shardIdx].Lock()
	shard[key] = value
	m.mu[shardIdx].Unlock()
}

func NewShardedMultiSegragatedRWMutexMap(shardCount int) *ShardedMultiSegregatedRWMutexMap {
	shards := make([]map[string]string, shardCount, shardCount)
	for i := 0; i < shardCount; i++ {
		shards[i] = make(map[string]string)
	}

	return &ShardedMultiSegregatedRWMutexMap{
		shards: shards,
		mu:     make([]sync.RWMutex, shardCount),
	}
}

type OrcamanLibrary struct {
	internal cmap.ConcurrentMap
}

func (m *OrcamanLibrary) Get(key string) (string, bool) {
	val, exists := m.internal.Get(key)
	if exists == false || val == nil {
		return "", false
	}
	return val.(string), exists
}

func (m *OrcamanLibrary) Set(key, value string) {
	m.internal.Set(key, value)
}

func NewOrcamanLibrary() *OrcamanLibrary {
	return &OrcamanLibrary{
		internal: cmap.New(),
	}
}

type FanliaoLibrary struct {
	internal *concurrent.ConcurrentMap
}

func (m *FanliaoLibrary) Get(key string) (string, bool) {
	val, err := m.internal.Get(key)
	if err != nil || val == nil {
		return "", false
	}
	return val.(string), true
}

func (m *FanliaoLibrary) Set(key, value string) {
	m.internal.Put(key, value)
}

func NewFanliaoLibrary() *FanliaoLibrary {
	return &FanliaoLibrary{
		internal: concurrent.NewConcurrentMap(),
	}
}

type TidwallLibrary struct {
	internal *shardmap.Map
}

func (m *TidwallLibrary) Get(key string) (string, bool) {
	val, exists := m.internal.Get(key)
	if !exists || val == nil {
		return "", false
	}
	return val.(string), true
}

func (m *TidwallLibrary) Set(key, value string) {
	m.internal.Set(key, value)
}

func NewTidwallLibrary() *TidwallLibrary {
	return &TidwallLibrary{
		internal: &shardmap.Map{},
	}
}

type SyncMap struct {
	sync.Map
}

func (m *SyncMap) Get(key string) (string, bool) {
	val, exists := m.Load(key)
	if exists == false || val == nil {
		return "", false
	}
	return val.(string), exists
}

func (m *SyncMap) Set(key, value string) {
	m.Store(key, value)
}

func NewSyncMap() *SyncMap {
	return &SyncMap{}
}

// getShardIndex returns the index of the shard that
// the key belongs in
func getShardIndex(key string, shards int) uint {
	return uint(djbHash(key)) % uint(shards)
}

// fnv32 deterministically generates a 32 bit number
// for a given string.
// https://github.com/orcaman/concurrent-map
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// DjbHash is for sharding the map.
// according to internets, this is fastest hashing algorithm ever made.
// we dont need security, we need distribution which this provides for us.
// https://github.com/zutto/shardedmap/blob/master/ShardedMap.go
func djbHash(key string) uint32 {
	var hash uint32 = 5381 //magic constant, apparently this hash fewest collisions possible.

	for _, chr := range key {
		hash = ((hash << 5) + hash) + uint32(chr)
	}
	return hash
}