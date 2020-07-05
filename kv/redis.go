package kv

import (
	"app/config"
	"fmt"
	"github.com/go-redis/redis/v7"
)

func setRedisDefaultConfigs() {
	config.C.Set("kv.redis.host", "localhost")
	config.C.Set("kv.redis.port", "6379")
	config.C.Set("kv.redis.database", "0")
	config.C.Set("kv.redis.ttl", 300)
}
func loadRedisConfigurationFromConfig() *redis.Options {

	host := config.C.GetString("kv.redis.host")
	port := config.C.GetString("kv.redis.port")
	username := config.C.GetString("kv.redis.username")
	password := config.C.GetString("kv.redis.password")
	database := config.C.GetInt("kv.redis.database")
	return &redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		Username: username,
		Password: password,
		DB: database,
	}
}
func NewRedis() (*redis.Client, error) {
	setRedisDefaultConfigs()
	options := loadRedisConfigurationFromConfig()
	c := redis.NewClient(options)
	stat := c.Ping()
	if stat.Err() != nil {
		return nil, stat.Err()
	}
	return c, nil
}