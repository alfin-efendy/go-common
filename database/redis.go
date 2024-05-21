package database

import (
	"context"
	"fmt"
	"time"

	"github.com/alfin87aa/go-common/configs"
	"github.com/alfin87aa/go-common/logger"
	"github.com/alfin87aa/go-common/servers/restapi"
	"github.com/redis/go-redis/v9"
)

func (d *database) initRedis(ctx context.Context) {
	config := configs.Configs.DB.Redis

	for id, redisConfig := range config {
		var redisClient *redis.Client

		switch redisConfig.Mode {
		case "single":
			redisClient = initSingleMode(ctx, redisConfig)

		case "sentinel":
			redisClient = initSentinelMode(ctx, redisConfig)
		default:
			logger.Fatalf(ctx, fmt.Errorf("redis mode %s is not supported", redisConfig.Mode), "❌ Failed unsupported redis mode")
			return
		}

		logger.Infoln(ctx, "✅ Redis client connected")

		restapi.AddChecker("redis-"+id, func(ctx context.Context) error {
			if _, err := redisClient.Ping(ctx).Result(); err != nil {
				return err
			}
			return nil
		})

		d.redisManager[id] = redisClient
	}
}

func initSingleMode(ctx context.Context, config *configs.Redis) *redis.Client {
	option := &redis.Options{
		Addr: config.Address,
	}

	if config.Username != nil {
		option.Username = *config.Username
	}
	if config.Password != nil {
		option.Password = *config.Password
	}
	if config.DB != nil {
		option.DB = *config.DB
	}
	if config.MinRetryBackoff != nil {
		option.MinRetryBackoff = time.Duration(*config.MinRetryBackoff) * time.Minute
	}
	if config.MaxRetryBackoff != nil {
		option.MaxRetryBackoff = time.Duration(*config.MaxRetryBackoff) * time.Minute
	}
	if config.DialTimeout != nil {
		option.DialTimeout = time.Duration(*config.DialTimeout) * time.Minute
	}
	if config.ReadTimeout != nil {
		option.ReadTimeout = time.Duration(*config.ReadTimeout) * time.Minute
	}
	if config.WriteTimeout != nil {
		option.WriteTimeout = time.Duration(*config.WriteTimeout) * time.Minute
	}
	if config.PoolFIFO != nil {
		option.PoolFIFO = *config.PoolFIFO
	}
	if config.PoolSize != nil {
		option.PoolSize = *config.PoolSize
	}
	if config.PoolTimeout != nil {
		option.PoolTimeout = time.Duration(*config.PoolTimeout) * time.Minute
	}
	if config.MinIdleConns != nil {
		option.MinIdleConns = *config.MinIdleConns
	}
	if config.MaxIdleConns != nil {
		option.MaxIdleConns = *config.MaxIdleConns
	}

	RedisClient := redis.NewClient(option)

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		fmt.Println("❌ Redis client failed to connect...")

		logger.Fatal(ctx, err, "❌ Redis client failed to connect")
	}
	return RedisClient
}

func initSentinelMode(ctx context.Context, config *configs.Redis) *redis.Client {
	option := &redis.FailoverOptions{
		MasterName:    config.MasterName,
		SentinelAddrs: config.SentinelAddress,
	}

	if config.Username != nil {
		option.Username = *config.Username
	}
	if config.Password != nil {
		option.Password = *config.Password
	}
	if config.DB != nil {
		option.DB = *config.DB
	}
	if config.MinRetryBackoff != nil {
		option.MinRetryBackoff = time.Duration(*config.MinRetryBackoff) * time.Minute
	}
	if config.MaxRetryBackoff != nil {
		option.MaxRetryBackoff = time.Duration(*config.MaxRetryBackoff) * time.Minute
	}
	if config.DialTimeout != nil {
		option.DialTimeout = time.Duration(*config.DialTimeout) * time.Minute
	}
	if config.ReadTimeout != nil {
		option.ReadTimeout = time.Duration(*config.ReadTimeout) * time.Minute
	}
	if config.WriteTimeout != nil {
		option.WriteTimeout = time.Duration(*config.WriteTimeout) * time.Minute
	}
	if config.PoolFIFO != nil {
		option.PoolFIFO = *config.PoolFIFO
	}
	if config.PoolSize != nil {
		option.PoolSize = *config.PoolSize
	}
	if config.PoolTimeout != nil {
		option.PoolTimeout = time.Duration(*config.PoolTimeout) * time.Minute
	}
	if config.MinIdleConns != nil {
		option.MinIdleConns = *config.MinIdleConns
	}
	if config.MaxIdleConns != nil {
		option.MaxIdleConns = *config.MaxIdleConns
	}

	RedisClient := redis.NewFailoverClient(option)

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		fmt.Println("❌ Redis client failed to connect...")

		logger.Fatal(ctx, err, "❌ Redis client failed to connect")
	}

	return RedisClient
}

func (d *database) redisConnection(ctx context.Context, id string) redis.UniversalClient {
	if client, ok := d.redisManager[id]; ok {
		return client
	}

	logger.Fatalf(ctx, fmt.Errorf("redis client %s not found", id), "❌ Failed to get redis client")
	return nil
}
