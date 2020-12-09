package redis

import (
	"app/config"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func loadRedisConfigurationFromConfig() (*redis.Options, error) {
	host := config.Config.DB.Redis.Host
	port := config.Config.DB.Redis.Port
	username := config.Config.DB.Redis.Username
	password := config.Config.DB.Redis.Password
	database := config.Config.DB.Redis.DB
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Username: username,
		Password: password,
		DB:       database,
	}, nil
}

var cli *redis.Client

func connect(ctx context.Context, options *redis.Options) (*redis.Client, error) {
	cli = redis.NewClient(options)
	stat := cli.Ping(ctx)
	if stat.Err() != nil {
		return nil, stat.Err()
	}
	return cli, nil
}

func NewRedis(ctx context.Context) (*redis.Client, error) {
	options, err := loadRedisConfigurationFromConfig()
	if err != nil {
		return nil, err
	}
	if cli == nil {
		return connect(ctx, options)
	}
	if err := cli.Ping(ctx).Err(); err != nil {
		return connect(ctx, options)
	}
	return cli, nil
}
