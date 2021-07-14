package TinyCache

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	mc := New()
	if mc == nil || mc.cache == nil {
		t.Fatalf("New Error ... ")
	}
}

func TestMainCache_Get(t *testing.T) {
	mc := New()
	k1, v1 := "test_key_1", "test_value_1"
	mc.Set(k1, String(v1))

	if v, err := mc.Get(k1); err != nil || string(v.(String)) != v1 {
		t.Fatalf("cache hit key = test_key_1 failed!")
	}

	k2 := "test_key_2"
	if _, err := mc.Get(k2); err == nil {
		t.Fatalf("cache hit nil key = test_key_2")
	}
}

func TestMainCache_Remove(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"
	cap := int64(len(k1 + k2 + v1 + v2))
	mc := NewByOption("lru", cap)
	mc.Set(k1, String(v1))
	mc.Set(k2, String(v2))
	mc.Set(k3, String(v3))

	if v, err := mc.Get("k1"); err == nil {
		log.Printf("v: %s， nbytes: %d, maxBytes: %d\n", string(v.(String)), mc.Size(), mc.MaxSize())
		t.Fatalf("RemoveOldest Error")
	}
}