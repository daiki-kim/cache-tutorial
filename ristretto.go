package main

import (
	"fmt"

	"github.com/dgraph-io/ristretto"
)

func main() {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})

	cache.Set("key", "value", 0)
	cache.Wait()

	if value, found := cache.Get("key"); found {
		fmt.Println("fFound value:", value)
	}
}
