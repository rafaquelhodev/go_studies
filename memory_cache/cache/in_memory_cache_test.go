package cache

import (
	"testing"
)

func TestGetKey(t *testing.T) {
	data := map[string]string{
		"test": "123",
	}
	cache := InMemoryCache{Data: data}

	gotValue, gotFound := cache.get("test")
	wantValue, wantFound := "123", true
	if gotValue != wantValue {
		t.Errorf("got %q, wanted %q", gotValue, wantValue)
	}

	if gotFound != wantFound {
		t.Errorf("got %t, wanted %t", gotFound, wantFound)
	}
}

func TestGetKeyNonExistentKey(t *testing.T) {
	data := map[string]string{
		"test": "123",
	}
	cache := InMemoryCache{Data: data}

	gotValue, gotFound := cache.get("test-haha")
	wantValue, wantFound := "", false
	if gotValue != wantValue {
		t.Errorf("got %q, wanted %q", gotValue, wantValue)
	}

	if gotFound != wantFound {
		t.Errorf("got %t, wanted %t", gotFound, wantFound)
	}
}
