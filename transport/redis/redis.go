package redis

import (
	"app/config"
	"fmt"

	"github.com/go-redis/redis/v7"
)

func loadRedisConfigurationFromConfig() (*redis.Options, error) {
	host := config.C.GetString("redis.host")
	port := config.C.GetString("redis.port")
	username := config.C.GetString("redis.username")
	password := config.C.GetString("redis.password")
	database := config.C.GetInt("redis.database")
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Username: username,
		Password: password,
		DB:       database,
	}, nil
}

var cli *redis.Client

func connect(options *redis.Options) (*redis.Client, error) {
	cli = redis.NewClient(options)
	stat := cli.Ping()
	if stat.Err() != nil {
		return nil, stat.Err()
	}
	return cli, nil
}

func NewRedis() (*redis.Client, error) {
	options, err := loadRedisConfigurationFromConfig()
	if err != nil {
		return nil, err
	}
	if cli == nil {
		return connect(options)
	}
	if err := cli.Ping().Err(); err != nil {
		return connect(options)
	}
	return cli, nil
}
