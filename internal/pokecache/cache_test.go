package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestCacheMiss(t *testing.T) {
  cache := NewCache(5 * time.Second)

  val, ok := cache.Get("non-existent-key")
  if ok {
    t.Errorf("expected cache miss, but key was found")
    return
  } 

  if val != nil {
    t.Errorf("expected nil value on cache miss")
    return
  }
}

func TestOverwriteValue(t *testing.T) {
  cache := NewCache(5 * time.Second)

  cache.Add("test-key", []byte("Original value"))
  cache.Add("test-key", []byte("Updated value"))

  val, ok := cache.Get("test-key")
  if !ok {
    t.Errorf("expected to find key")
    return
  }
  if string(val) != "Updated value" {
    t.Errorf("expected 'Updated value' \n got %s", string(val))
    return
  }
}

func TestMultipleKeys(t *testing.T) {
    cache := NewCache(5 * time.Second)
    
    // Add multiple different keys
    cache.Add("key1", []byte("value1"))
    cache.Add("key2", []byte("value2"))
    cache.Add("key3", []byte("value3"))
    
    // Make sure each key returns its correct value
    val1, ok1 := cache.Get("key1")
    val2, ok2 := cache.Get("key2")
    val3, ok3 := cache.Get("key3")
    
    if !ok1 || !ok2 || !ok3 {
        t.Errorf("expected to find all keys")
        return
    }
    
    if string(val1) != "value1" || string(val2) != "value2" || string(val3) != "value3" {
        t.Errorf("values don't match expected")
        return
    }
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
	t.Errorf("expected to not find key")
		return
	}
}
