package TinyCache

import (
	"container/list"
)

type fifoEntry struct {
	key   string
	value Value
}

type FIFOCache struct {
	maxBytes int64
	nbytes   int64 // used space
	ll       *list.List
	cache    map[string]*list.Element
}

func NewFIFOCache(maxBytes int64) *FIFOCache {
	return &FIFOCache{
		maxBytes: maxBytes,
		nbytes:   0,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

func (f *FIFOCache) Set(key string, value Value) {
	if ele, ok := f.cache[key]; ok {
		kv := ele.Value.(*fifoEntry)
		kv.value = value
		f.nbytes += int64(value.Len()) - int64(kv.value.Len())
	} else {
		tmp := f.ll.PushFront(&fifoEntry{key: key, value: value})
		f.cache[key] = tmp
		f.nbytes += int64(len(key)) + int64(value.Len())
	}

	for f.nbytes >= f.maxBytes {
		// log.Printf("nBytes: %d, maxBytes: %d\n", f.nbytes, f.maxBytes)
		f.RemoveLast()
	}
}

func (f *FIFOCache) Get(key string) (Value, error) {
	if ele, ok := f.cache[key]; ok {
		kv := ele.Value.(*fifoEntry)
		return kv.value, nil
	}
	return nil, ErrKeyNotFound
}

func (f *FIFOCache) Remove(key string) error {
	if ele, ok := f.cache[key]; ok {
		kv := ele.Value.(*fifoEntry)
		f.ll.Remove(ele)
		delete(f.cache, kv.key)
		f.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
	}
	return ErrKeyNotFound
}

func (f *FIFOCache) Purge() {
	f.ll = nil
	f.cache = nil
	f.nbytes = 0
}

func (f *FIFOCache) Size() int64 {
	return f.nbytes
}

func (f *FIFOCache) MaxSize() int64 {
	return f.maxBytes
}

func (f *FIFOCache) RemoveLast() {
	if f.ll == nil || f.ll.Len() == 0 {
		return
	}

	ele := f.ll.Back()
	kv := ele.Value.(*fifoEntry)
	f.ll.Remove(ele)
	delete(f.cache, kv.key)
	f.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
}
