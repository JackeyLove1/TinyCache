package TinyCache

import (
	"log"
	"testing"
)

func TestLRUKCache_Get(t *testing.T) {
	k := NewLRUKCache(DefaultLRUK, MaxCacheSize)
	k1, v1 := "test_key_1", "test_value_1"
	k.Set(k1, String(v1))

	if v, err := k.Get(k1); err != nil || string(v.(String)) != v1 {
		// log.Println("v: ", string(v.(String)))
		t.Fatalf("cache hit key = test_key_1 failed!")
	}

	k2 := "test_key_2"
	if _, err := k.Get(k2); err == nil {
		t.Fatalf("cache hit nil key = test_key_2")
	}
}

func TestLRUKCache_Remove(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	k := NewLRUKCache(DefaultLRUK, int64(cap+1))
	k.Set(k1, String(v1))
	k.Set(k2, String(v2))
	k.Set(k3, String(v3))

	if v, err := k.Get("k1"); err == nil {
		log.Println("v: ", string(v.(String)))
		t.Fatalf("RemoveOldest Error")
	}
}
