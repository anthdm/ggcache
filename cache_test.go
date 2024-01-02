package ggcache

import (
	"testing"
	"time"
)

// TestCache_Get tests the Get method of the Cache.
func TestCache_Get(t *testing.T) {
	cache := New()

	// Test Case 1: Key not found
	_, err := cache.Get([]byte("nonexistent"))
	if err == nil {
		t.Error("Expected error for nonexistent key, but got nil")
	}

	// Test Case 2: Key found
	key := []byte("testKey")
	value := []byte("testValue")
	_ = cache.Set(key, value, 0)

	retrievedValue, err := cache.Get(key)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(retrievedValue) != string(value) {
		t.Errorf("Expected value %s, but got %s", value, retrievedValue)
	}
}

// TestCache_Set tests the Set method of the Cache.
func TestCache_Set(t *testing.T) {
	cache := New()

	// Test Case 1: Set with TTL
	key := []byte("testKey")
	value := []byte("testValue")
	ttl := time.Millisecond * 100
	err := cache.Set(key, value, ttl)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Wait for TTL to expire
	time.Sleep(ttl + time.Millisecond*50)

	// Check that key is not present in the cache
	if cache.Has(key) {
		t.Error("Expected key to be expired, but it's still present")
	}
}

// TestCache_Has tests the Has method of the Cache.
func TestCache_Has(t *testing.T) {
	cache := New()

	// Test Case 1: Key not present
	if cache.Has([]byte("nonexistent")) {
		t.Error("Expected false for nonexistent key, but got true")
	}

	// Test Case 2: Key present
	key := []byte("testKey")
	value := []byte("testValue")
	_ = cache.Set(key, value, 0)

	if !cache.Has(key) {
		t.Error("Expected true for existing key, but got false")
	}
}

// TestCache_Delete tests the Delete method of the Cache.
func TestCache_Delete(t *testing.T) {
	cache := New()

	// Test Case 1: Delete nonexistent key
	err := cache.Delete([]byte("nonexistent"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test Case 2: Delete existing key
	key := []byte("testKey")
	value := []byte("testValue")
	_ = cache.Set(key, value, 0)

	err = cache.Delete(key)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check that key is not present in the cache
	if cache.Has(key) {
		t.Error("Expected key to be deleted, but it's still present")
	}
}

// TestCacheIntegration tests the Cache integration by combining multiple operations.
func TestCacheIntegration(t *testing.T) {
	cache := New()

	// Test Case 1: Set and Get
	key := []byte("testKey")
	value := []byte("testValue")
	err := cache.Set(key, value, 0)
	if err != nil {
		t.Errorf("Unexpected error during Set: %v", err)
	}

	retrievedValue, err := cache.Get(key)
	if err != nil {
		t.Errorf("Unexpected error during Get: %v", err)
	}

	if string(retrievedValue) != string(value) {
		t.Errorf("Expected value %s, but got %s", value, retrievedValue)
	}

	// Test Case 2: Set with TTL
	ttl := time.Millisecond * 100
	err = cache.Set(key, value, ttl)
	if err != nil {
		t.Errorf("Unexpected error during Set with TTL: %v", err)
	}

	// Wait for TTL to expire
	time.Sleep(ttl + time.Millisecond*50)

	// Check that key is not present in the cache
	if cache.Has(key) {
		t.Error("Expected key to be expired, but it's still present")
	}

	// Test Case 3: Delete
	err = cache.Delete(key)
	if err != nil {
		t.Errorf("Unexpected error during Delete: %v", err)
	}

	// Check that key is not present in the cache
	if cache.Has(key) {
		t.Error("Expected key to be deleted, but it's still present")
	}
}
