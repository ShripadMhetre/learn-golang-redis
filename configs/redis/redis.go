package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func ConnectRedis(ctx context.Context) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
}
