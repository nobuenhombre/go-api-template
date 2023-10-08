package store

import (
	"fmt"
	"go-api-template/src/internal/app/api-my-domain-com/api/server/config"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	RedisClient *redis.Client
}

func NewStore(config *config.RedisConfig) *Store {
	store := new(Store)

	store.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	return store
}
