package cache

import (
	"testing"
	"time"
)

func TestNewFreeCacheClient(t *testing.T) {
	client := NewFreeCacheClient(1024)
	if client == nil {
		t.Fatal("NewFreeCacheClient() should not return nil")
	}
}

func TestFreeCacheClient_SetGet(t *testing.T) {
	client := NewFreeCacheClient(1024)

	key := []byte("testkey")
	value := []byte("testvalue")

	err := client.Set(key, value, 60)
	if err != nil {
		t.Fatalf("Set() error = %v", err)
	}

	got, err := client.Get(key)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if string(got) != string(value) {
		t.Errorf("Get() = %q, want %q", string(got), string(value))
	}
}

func TestFreeCacheClient_Get_NotFound(t *testing.T) {
	client := NewFreeCacheClient(1024)

	_, err := client.Get([]byte("nonexistent"))
	if err == nil {
		t.Error("Get() should return error for nonexistent key")
	}
}

func TestFreeCacheClient_Del(t *testing.T) {
	client := NewFreeCacheClient(1024)

	key := []byte("testkey")
	value := []byte("testvalue")

	client.Set(key, value, 60)
	client.Del("testkey")

	_, err := client.Get(key)
	if err == nil {
		t.Error("Get() should return error after Del()")
	}
}

func TestFreeCacheClient_Clear(t *testing.T) {
	client := NewFreeCacheClient(1024)

	client.Set([]byte("key1"), []byte("value1"), 60)
	client.Set([]byte("key2"), []byte("value2"), 60)

	client.Clear()

	_, err := client.Get([]byte("key1"))
	if err == nil {
		t.Error("Get() should return error after Clear()")
	}

	_, err = client.Get([]byte("key2"))
	if err == nil {
		t.Error("Get() should return error after Clear()")
	}
}

func TestFreeCacheClient_Expiry(t *testing.T) {
	client := NewFreeCacheClient(1024)

	key := []byte("testkey")
	value := []byte("testvalue")

	// Set with 1 second expiry
	client.Set(key, value, 1)

	// Should be available immediately
	got, err := client.Get(key)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if string(got) != string(value) {
		t.Errorf("Get() = %q, want %q", string(got), string(value))
	}

	// Wait for expiry
	time.Sleep(1100 * time.Millisecond)

	_, err = client.Get(key)
	if err == nil {
		t.Error("Get() should return error after expiry")
	}
}
