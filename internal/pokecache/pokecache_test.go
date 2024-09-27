package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, cs.inputVal)

		actual, ok := cache.Get(cs.inputKey)
		if !ok {
			t.Errorf("%s not found", cs.inputKey)
			continue
		}
		if string(actual) != string(cs.inputVal) {
			t.Errorf("%s does not match %s", string(actual), string(cs.inputVal))
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))
	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get("key1")
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}
