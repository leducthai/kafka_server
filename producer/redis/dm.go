package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	RedisConnect string
	RedisPass    string
}

func NewRedis(mc RedisConfig, rdb int) *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     mc.RedisConnect,
		Password: mc.RedisPass,
		DB:       rdb,
	})

	// Ping Redis to check the connection
	_, err := r.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err)
	}

	return r
}

type RedisManager struct {
	FirstTable  *redis.Client
	SecondTable *redis.Client
}

func NewDM(mc RedisConfig) RedisManager {
	return RedisManager{
		FirstTable:  NewRedis(mc, 0),
		SecondTable: NewRedis(mc, 1),
	}
}

func (r RedisManager) 
