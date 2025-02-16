package configs

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func NewRedis(env *Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.RedisHost,
		Password: env.RedisPass,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Gagal terhubung ke Redis: %v", err)
	}

	return client
}
