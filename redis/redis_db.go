package redis

import (
	"bnqkl/chain-cms/config"
	"bnqkl/chain-cms/logger"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisDb *redis.Client

func GetRedisBb() *redis.Client {
	return redisDb
}

func InitRedisDb(log *logger.Logger) error {
	config := config.GetConfig()
	redisConfig := config.Redis
	initRedisDb := func() error {
		rdb := redis.NewClient(&redis.Options{
			Addr:        fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
			Password:    redisConfig.Password,
			DB:          redisConfig.Db,
			DialTimeout: time.Minute,
		})
		_, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			return err
		}
		redisDb = rdb
		log.Info("init redis success")
		return nil
	}
	var err error
	for i := 0; i < 3; i++ {
		err = initRedisDb()
		if err == nil {
			return nil
		}
		log.Error(err)
		time.Sleep(time.Second)
	}
	return err
}
