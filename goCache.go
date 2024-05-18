package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)

	c.Set("key", "value", cache.DefaultExpiration)

	if x, found := c.Get("key"); found {
		fmt.Println("Found value:", x)
	}
}
