package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", Redis_Host, Redis_Port),
		DB:   Redis_Db,
	})
}
