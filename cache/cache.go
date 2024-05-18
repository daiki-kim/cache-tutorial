package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
)

func main() {
	cache, _ := lru.New(128)
	cache.Add("key", "value")
	if value, ok := cache.Get("key"); ok {
		fmt.Println("Found value:", value)
	} else {
		fmt.Println("Key not found")
	}
}
