package main

import (
	"fmt"

	"github.com/coocood/freecache"
)

func main() {
	cacheSize := 100 * 1024 * 1024 // 100MB
	cache := freecache.NewCache(cacheSize)

	cache.Set([]byte("key"), []byte("value"), 60) // 60秒の有効期限
	entry, err := cache.Get([]byte("key"))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found value:", string(entry))
	}
}
