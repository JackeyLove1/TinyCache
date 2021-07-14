package TinyCache

import (
	"log"
	"testing"
)

func TestFIFOCache_Get(t *testing.T) {
	f := NewFIFOCache(MaxCacheSize)
	k1, v1 := "test_key_1", "test_value_1"
	f.Set(k1, String(v1))

	if v, err := f.Get(k1); err != nil || string(v.(String)) != v1 {
		t.Fatalf("cache hit key = test_key_1 failed!")
	}

	k2 := "test_key_2"
	if _, err := f.Get(k2); err == nil {
		t.Fatalf("cache hit nil key = test_key_2")
	}
}

func TestFIFOCache_Remove(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	f := NewFIFOCache(int64(cap + 1))
	f.Set(k1, String(v1))
	f.Set(k2, String(v2))
	f.Set(k3, String(v3))

	if v, err := f.Get("k1"); err == nil {
		log.Println("v1: ", string(v.(String)))
		t.Fatalf("RemoveOldest Error")
	}
}
