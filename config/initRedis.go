package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", Redis_Host, Redis_Port),
		DB:   Redis_Db,
	})
	_, err := RedisClient.Do(Ctx, "INFO").Result()
	if err != nil {
		Logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"DB":    RedisClient,
		}).Error("error while connecting to redis")
		log.Fatalf("Error while connecting to redis db, %v", err)
	}
	log.Println("Redis client initialized successfully")
}
