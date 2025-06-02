package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/marcosvfn/url-shortener/internal/domain/url"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(addr string) (*RedisRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisRepository{client: client}, nil
}

func (r *RedisRepository) Save(url *url.URL) error {
	return r.client.Set(context.Background(), url.ShortCode, url.OriginalURL, 0).Err()
}

func (r *RedisRepository) FindByShortCode(shortCode string) (*url.URL, error) {
	originalURL, err := r.client.Get(context.Background(), shortCode).Result()
	if err == redis.Nil {
		return nil, errors.New("URL not found")
	}
	if err != nil {
		return nil, err
	}
	return &url.URL{OriginalURL: originalURL, ShortCode: shortCode}, nil
}
