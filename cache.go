package TinyCache

import (
	"math"
	"sync"
)

const (
	TYPE_FIFO = "fifo"
	TYPE_LRU  = "lru"
	TYPE_LRUK = "lruk"
	TYPE_LFU  = "lfu"
	TYPE_ARC  = "arc"
	TYPE_2Q   = "2q"
)

// the interface of Cache
type Cache interface {
	// Set or Update the key-value pair
	Set(key string, value Value)

	// Get returns the value for the specific key-value pair
	Get(key string) (Value, error)

	// Remove removes the specific key from the cache if the key is present
	Remove(key string) error

	// Purge is used to completely clear the cache
	Purge()

	// Size returns the space used by the cache
	Size() int64

	// MaxSize returns the maxSize of the cache
	MaxSize() int64
}

type mainCache struct {
	mu        sync.RWMutex
	cache     Cache
	cacheType string
	OnEvicted func(key string, value Value)
}

type Option struct {
	cacheType string
	maxBytes  int64
	k         int
	OnEvicted func(key string, value Value)
}

var DefaultOption = Option{
	cacheType: "lru",
	maxBytes:  int64(math.MaxInt32),
	k:         DefaultLRUK,
	OnEvicted: nil,
}

type ModOption func(option *Option)

func New(opts ...Option) *mainCache {
	return NewByOption(DefaultOption.cacheType, DefaultOption.maxBytes)
}

func NewByOption(cacheType string, maxBytes int64, opts ...Option) *mainCache {
	option := DefaultOption
	option.cacheType = cacheType

	var cache Cache
	switch cacheType {
	case "lru":
		cache = NewLRUCache(maxBytes)
	case "fifo":
		cache = NewFIFOCache(maxBytes)
	case "lfu":
		cache = NewLFUCache(maxBytes)
	case "lruk":
		cache = NewLRUKCache(option.k, maxBytes)
	default:
		cache = NewLRUCache(maxBytes)
	}

	return &mainCache{
		cache:     cache,
		cacheType: cacheType,
	}
}

func (m *mainCache) Set(key string, value Value) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache.Set(key, value)
}

func (m *mainCache) Get(key string) (Value, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.cache.Get(key)
}

func (m *mainCache) Remove(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.cache.Remove(key)
}

func (m *mainCache) Purge() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache.Purge()
}

func (m *mainCache) Size() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.cache.Size()
}

func (m *mainCache) MaxSize() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.cache.MaxSize()
}
