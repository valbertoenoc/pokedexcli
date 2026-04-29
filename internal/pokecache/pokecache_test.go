package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	interval := 5 * time.Millisecond
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com/endpoint",
			val: []byte("test data"),
		},
		{
			key: "https://example.com/endpoint",
			val: []byte("test data 2"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			actualVal, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}

			if string(actualVal) != string(c.val) {
				t.Errorf("error getting cache value: %s != %v", string(actualVal), string(c.val))
				return
			}

		})
	}

}

func TestReapLoop(t *testing.T) {
	interval := time.Millisecond * 20
	cache := NewCache(interval)

	urlKey := "https://example.com/endpoint"
	urlKey2 := "https://example.com/endpoint2"
	testData := []byte("some-data")
	cache.Add(urlKey, testData)

	// before expiration
	val, ok := cache.Get(urlKey)
	if !ok && string(val) != string(testData) {
		t.Error("key not found before expiration")
	}

	// after expiration
	time.Sleep(time.Millisecond * 30)
	// add new entry after first expiration
	cache.Add(urlKey2, testData)
	val, ok = cache.Get(urlKey)
	if ok {
		t.Error("key found after expiration")
	}

	// check if second entry still present
	if len(cache.entries) == 0 {
		t.Errorf("expected to see second entry")
	}

}
