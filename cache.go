package ggcache

import (
	"fmt"
	"sync"
	"time"
)

// Cacher is an interface used for performing caching operations.
// Applications can implement this interface to integrate different caching managers.
type Cacher interface {
	// Get returns the value associated with the specified key.
	// If the key is not found or an error occurs, an error object is returned.
	Get(key []byte) ([]byte, error)

	// Set adds the value associated with the specified key to the cache with the specified expiration time.
	// If the duration is zero, the cache is held indefinitely.
	Set(key []byte, value []byte, expiration time.Duration) error

	// Has checks whether the specified key exists in the cache.
	Has(key []byte) bool

	// Delete removes the specified key from the cache.
	// If the key is not found, an error object is returned.
	Delete(key []byte) error
}

// Cache is a simple in-memory cache implementation.
// It utilizes a sync.RWMutex for concurrent read and write safety.
// The cache stores data as byte slices, using string keys for retrieval.
type Cache struct {
	// lock is a sync.RWMutex to ensure concurrent read and write safety.
	lock sync.RWMutex

	// data is a map that stores byte slices with string keys for retrieval.
	data map[string][]byte
}

// New creates and returns a new instance of the Cache with initialized internal data.
// The Cache is an in-memory cache implementation using a sync.RWMutex for concurrency safety.
// The internal data is represented as a map with string keys and byte slice values.
func New() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

// Get retrieves the value associated with the specified key from the cache.
// It acquires a read lock to ensure concurrent safety during retrieval.
// If the key is not found, an error is returned indicating the absence of the key.
// The retrieved value and a nil error are returned if the key is present in the cache.
func (c *Cache) Get(key []byte) ([]byte, error) {
	// Acquire a read lock to ensure concurrent safety during retrieval.
	c.lock.RLock()
	defer c.lock.RUnlock()

	// Convert the byte slice key to a string for map lookup.
	keyStr := string(key)

	// Retrieve the value associated with the key from the internal data map.
	val, ok := c.data[keyStr]
	if !ok {
		// Return an error if the key is not found.
		return nil, fmt.Errorf("key (%s) not found", keyStr)
	}

	// Return the retrieved value and a nil error if the key is present in the cache.
	return val, nil
}

// Set adds or updates the cache with the specified key-value pair.
// It acquires a write lock to ensure concurrent safety during insertion.
// If the time-to-live (TTL) duration is greater than zero, a goroutine is launched to remove the entry after the specified duration.
// The key-value pair is stored in the cache, and if a TTL is set, the entry is automatically deleted after the specified duration.
// The method returns nil, indicating a successful operation.
func (c *Cache) Set(key, value []byte, ttl time.Duration) error {
	// Acquire a write lock to ensure concurrent safety during insertion.
	c.lock.Lock()
	defer c.lock.Unlock()

	// Convert the byte slice key to a string for map storage.
	keyStr := string(key)

	// Add or update the cache with the specified key-value pair.
	c.data[keyStr] = value

	// If TTL is greater than zero, launch a goroutine to remove the entry after the specified duration.
	if ttl > 0 {
		go func() {
			<-time.After(ttl)
			c.lock.Lock()
			defer c.lock.Unlock()
			delete(c.data, keyStr)
		}()
	}

	// Return nil, indicating a successful operation.
	return nil
}

// Has checks if the specified key exists in the cache.
// It acquires a read lock to ensure concurrent safety during the lookup.
// The method returns true if the key is found in the cache, and false otherwise.
func (c *Cache) Has(key []byte) bool {
	// Acquire a read lock to ensure concurrent safety during the lookup.
	c.lock.RLock()
	defer c.lock.RUnlock()

	// Check if the key exists in the cache.
	_, ok := c.data[string(key)]

	// Return true if the key is found, and false otherwise.
	return ok
}

// Delete removes the specified key from the cache.
// It acquires a write lock to ensure concurrent safety during deletion.
// The method returns nil, indicating a successful deletion.
func (c *Cache) Delete(key []byte) error {
	// Acquire a write lock to ensure concurrent safety during deletion.
	c.lock.Lock()
	defer c.lock.Unlock()

	// Remove the specified key from the cache.
	delete(c.data, string(key))

	// Return nil, indicating a successful deletion.
	return nil
}
