package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewRedis(address string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return &RedisClient{
		Client: client,
	}, nil
}

func (r *RedisClient) HSet(key, value string) error {
	return r.Client.HSet(Ctx, key, value).Err()
}

func (r *RedisClient) HDel(key string, fields ...string) error {
	return r.Client.HDel(Ctx, key, fields...).Err()
}

func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	cmd := r.Client.HGetAll(Ctx, key)
	if err := cmd.Err(); err != nil {
		return map[string]string{}, err
	}

	return cmd.Val(), nil
}

func (r *RedisClient) HGet(key, field string) (string, error) {
	cmd := r.Client.HGet(Ctx, key, field)
	if err := cmd.Err(); err != nil {
		return "", err
	}

	return cmd.Val(), nil
}
