package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client
var Ctx context.Context

func ConnectRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := Cache.Ping(Ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
}
