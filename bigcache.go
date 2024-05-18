package main

import (
	"fmt"
	"time"

	"github.com/allegro/bigcache"
)

func main() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("key", []byte("value"))
	entry, _ := cache.Get("key")
	fmt.Println("Found value:", string(entry))
}
