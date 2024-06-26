package drivers

import (
	"github.com/daqiancode/env"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func GetRedis() *redis.Client {
	if redisClient != nil {
		return redisClient
	}
	redisClient = CreateRedis()
	return redisClient
}

func CreateRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     env.Get("REDIS_ADDR", "localhost:6379"),
		Username: env.Get("REDIS_USERNAME", ""),
		Password: env.Get("REDIS_PASSWORD", ""), // no password set
		DB:       env.GetIntMust("REDIS_DB", 0), // use default DB
	})
}
