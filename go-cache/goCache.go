package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	// キャッシュデータの有効期限が5分、10分ごとにキャッシュをリフレッシュする
	c := cache.New(5*time.Minute, 10*time.Minute)

	// キャッシュにデータを保存
	c.Set("key", "value", cache.DefaultExpiration)

	// キャッシュからデータを取得
	if x, found := c.Get("key"); found {
		fmt.Println("Found value:", x)
	}
}
