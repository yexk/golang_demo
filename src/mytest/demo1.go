package mytest

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Demo1 测试1
func Demo1() {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	redisdb.Options().DB = 1
	pong, err := redisdb.Set("a", "撒旦法打是", 0).Result()
	fmt.Println(pong, err)
}
