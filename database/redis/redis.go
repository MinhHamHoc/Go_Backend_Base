package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

var AddressUrl = "localhost:6379"

func connectRedis(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     AddressUrl,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	redisClient = client
}

func HSet(ctx context.Context, key, val string) {
	err := redisClient.Set(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func HGet(ctx context.Context, key string) string {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val
}

func HGetAll(ctx context.Context, key string) []string {
	keys := []string{}

	iter := redisClient.Scan(ctx, 0, key, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	return keys
}
