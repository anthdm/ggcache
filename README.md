# ggcache

`ggcache` is a flexible and lightweight in-memory cache package for Go, suitable for both single-cache applications and distributed systems with multiple clients.

## Overview

This package provides a basic in-memory caching mechanism with support for key-value storage, expiration time, and concurrent read and write safety. It is designed to be versatile and adaptable to various use cases, from standalone applications to distributed systems where multiple clients share a common cache.

## Features

- **Concurrent Safety:** The cache utilizes a `sync.RWMutex` to ensure safe access and modification of data in a concurrent environment.

- **Expiration Time:** Cached items can have an optional expiration time, allowing automatic removal of entries after a specified duration.

- **Client-Server Architecture:** The package supports a client-server architecture, enabling multiple clients to interact with a centralized cache server.

## Installation

To use `ggcache` in your Go project, you can install it using the following:

```bash
go get github.com/anthdm/ggcache
```

## Single Cache Example
```go
package main

import (
	"fmt"
	"time"

	"github.com/anthdm/ggcache"
)

func main() {
	// Create a new cache instance
	cache := ggcache.New()

	// Set a key-value pair in the cache
	key := []byte("exampleKey")
	value := []byte("exampleValue")
	err := cache.Set(key, value, time.Second*30)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Get the value from the cache
	result, err := cache.Get(key)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", string(result))
	}

	// Check if a key exists in the cache
	if cache.Has(key) {
		fmt.Println("Key exists in the cache.")
	}

	// Delete a key from the cache
	err = cache.Delete(key)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```

## Contributing
Feel free to contribute to the development of ggcache by submitting issues or pull requests.

## License
This package is licensed under the MIT License - see the LICENSE file for details.
