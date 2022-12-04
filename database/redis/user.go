package redis

var UserRedisAccessIdPrefix = "user_access"

type UserAccessIDRedisRepository struct {
	redisClient *RedisClient
}
