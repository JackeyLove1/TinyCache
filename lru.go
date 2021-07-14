package TinyCache

import (
	"container/list"
)

type lruEntry struct {
	key   string
	value Value
}

type LRUCache struct {
	maxBytes int64
	nbytes   int64 // used space
	ll       *list.List
	cache    map[string]*list.Element
}

// New is the constructor of Cache
func NewLRUCache(maxBytes int64) *LRUCache {
	return &LRUCache{
		maxBytes: maxBytes,
		nbytes:   0,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

func (lru *LRUCache) Set(key string, value Value) {
	if ele, ok := lru.cache[key]; ok {
		lru.ll.MoveToFront(ele)
		kv := ele.Value.(*lruEntry)
		lru.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := lru.ll.PushFront(&lruEntry{key: key, value: value})
		lru.cache[key] = ele
		lru.nbytes += int64(len(key)) + int64(value.Len())
	}

	for lru.maxBytes <= lru.nbytes {
		lru.removeOldest()
	}

}

func (lru *LRUCache) Get(key string) (Value, error) {
	if ele, ok := lru.cache[key]; ok {
		lru.ll.MoveToFront(ele)
		kv := ele.Value.(*lruEntry)
		return kv.value, nil
	}
	return nil, ErrKeyNotFound
}

func (lru *LRUCache) Remove(key string) error {
	if ele, ok := lru.cache[key]; ok {
		lru.ll.Remove(ele)
		kv := ele.Value.(*lruEntry)
		delete(lru.cache, kv.key)
		lru.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		return nil
	}
	return ErrKeyNotFound
}

func (lru *LRUCache) Size() int64 {
	return lru.nbytes
}

func (lru *LRUCache) MaxSize() int64 {
	return lru.maxBytes
}

func (lru *LRUCache) Purge() {
	lru.ll = nil
	lru.cache = nil
}

// Removes the oldest item
func (lru *LRUCache) removeOldest() {
	ele := lru.ll.Back()
	if ele != nil {
		lru.ll.Remove(ele)
		kv := ele.Value.(*lruEntry)
		delete(lru.cache, kv.key)
		lru.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
	}
}
