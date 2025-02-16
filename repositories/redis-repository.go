package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	GetCache(context.Context, string, interface{}) error
	SetCache(context.Context, string, interface{}, time.Duration) error
	DeleteCache(context.Context, string) error
	GetKeys(context.Context, string) ([]string, error)
}

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) RedisRepository {
	return &redisRepository{client: client}
}

func (r *redisRepository) GetCache(ctx context.Context, key string, dest interface{}) error {
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

func (r *redisRepository) SetCache(ctx context.Context, key string, data interface{}, ttl time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, jsonData, ttl).Err()
}

func (r *redisRepository) DeleteCache(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *redisRepository) GetKeys(ctx context.Context, pattern string) ([]string, error) {
	var keys []string
	var cursor uint64
	const batchSize = 50

	for {
		var batch []string
		var err error
		batch, cursor, err = r.client.Scan(ctx, cursor, pattern, batchSize).Result()
		if err != nil {
			return nil, err
		}

		keys = append(keys, batch...)
		if cursor == 0 {
			break
		}
	}

	return keys, nil
}
