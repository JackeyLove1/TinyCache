package TinyCache

import "testing"

func TestLRUCache_Get(t *testing.T) {
	lru := NewLRUCache(MaxCacheSize)
	k1, v1 := "test_key_1", "test_value_1"
	lru.Set(k1, String(v1))

	if v, err := lru.Get(k1); err != nil || string(v.(String)) != v1 {
		t.Fatalf("cache hit key = test_key_1 failed!")
	}

	k2 := "test_key_2"
	if _, err := lru.Get(k2); err == nil {
		t.Fatalf("cache hit nil key = test_key_2")
	}
}

func TestLRUCache_RemoveOldest(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := NewLRUCache(int64(cap + 1))
	lru.Set(k1, String(v1))
	lru.Set(k2, String(v2))
	lru.Set(k3, String(v3))

	if _, err := lru.Get("k1"); err == nil {
		t.Fatalf("RemoveOldest Error")
	}
}