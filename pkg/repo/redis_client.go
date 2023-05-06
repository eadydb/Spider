package repo

import (
	"strings"

	"github.com/redis/go-redis/v9"
)

var RedisClient redis.UniversalClient

type RedisConfig struct {
	MasterName string
	Addrs      string
	Password   string
}

// NewDefaultConfig returns a default redis config.
func NewDefaultRedisConfig() *RedisConfig {
	return &RedisConfig{
		Addrs:    ":6379",
		Password: "",
	}
}

// NewConfig returns a redis config.
func NewRedisConfig(masterName, Addrs, Password string) *RedisConfig {
	return &RedisConfig{
		MasterName: masterName,
		Addrs:      Addrs,
		Password:   Password,
	}
}

// NewClient returns a redis client.
func (c *RedisConfig) NewRedisClient() redis.UniversalClient {
	if "" == c.MasterName {
		return redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    strings.Split(c.Addrs, ","),
			Password: c.Password,
		})
	}
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      strings.Split(c.Addrs, ","),
		Password:   c.Password,
		MasterName: c.MasterName,
	})
}
