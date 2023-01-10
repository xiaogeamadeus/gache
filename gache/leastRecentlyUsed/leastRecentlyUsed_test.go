package leastRecentlyUsed

import (
	"reflect"
	"testing"
)

type String string

func (data String) Len() int {
	return len(data)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("testKey1", String("1234"))
	if value, ok := lru.Get("testKey1"); !ok || string(value.(String)) != "1234" {
		t.Fatalf("cace hit testKey1=1234 failed")
	}
	if _, ok := lru.Get("testKey2"); ok {
		t.Fatalf("cache miss testKey2 failed")
	}
}

func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "testKey1", "testKey2", "testKey3"
	v1, v2, v3 := "testValue1", "testValue2", "testValue3"
	capacity := len(k1 + k2 + v1 + v2)
	lru := New(int64(capacity), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("testKey1"); ok || lru.Len() != 2 {
		t.Fatalf("RemoveOldest testkey1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := New(int64(10), callback)
	lru.Add("key1", String("123456"))
	lru.Add("k2", String("k2"))
	lru.Add("k3", String("k3"))
	lru.Add("k4", String("k4"))

	expect := []string{"key1", "k2"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call onEvicted failed, expect keys equals to %s", expect)
	}
}