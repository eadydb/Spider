package repo

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type KeyChecker struct {
	rdb    redis.UniversalClient
	key    string
	expire time.Duration
}

func NewKeyChecker(key string, expire time.Duration) *KeyChecker {
	return &KeyChecker{
		rdb:    RedisClient,
		key:    key,
		expire: expire,
	}
}

func (c *KeyChecker) Save(value interface{}) error {
	return c.rdb.Set(context.Background(), c.key, value, c.expire).Err()
}

func (c *KeyChecker) Get() (string, error) {
	return c.rdb.Get(context.Background(), c.key).Result()
}
