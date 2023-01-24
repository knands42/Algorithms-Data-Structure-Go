package lrucache

import "testing"

func Test_CacheGettersAndSetters(t *testing.T) {
	cache := NewLRU(1)
	cache.Set("key", "value")
	value, ok := cache.Get("key")
	if !ok {
		t.Errorf("Expected value to be set")
	}
	if value != "value" {
		t.Errorf("Expected value to be 'value', got '%s'", value)
	}
}

func Test_CacheGettersAndSettersWithOverwriteData(t *testing.T) {
	cache := NewLRU(1)
	cache.Set("key", "value")
	value, ok := cache.Get("key")
	if !ok {
		t.Errorf("Expected value to be set")
	}
	if value != "value" {
		t.Errorf("Expected value to be 'value', got '%s'", value)
	}

	cache.Set("key", "new-value")
	value, ok = cache.Get("key")
	if !ok {
		t.Errorf("Expected value to be set")
	}
	if value != "new-value" {
		t.Errorf("Expected value to be 'new-value', got '%s'", value)
	}
}

func Test_LRU(t *testing.T) {
	cache := NewLRU(3)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	cache.Set("key4", "value4")

	u, ok := cache.Get("key1")
	if ok {
		t.Errorf("Expected value to be evicted")
	}
	if u != "" {
		t.Errorf("Expected value to be '', got '%s'", u)
	}
}
