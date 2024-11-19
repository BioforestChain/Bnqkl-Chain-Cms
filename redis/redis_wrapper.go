package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisWrapper struct {
}

var redisWrapper *RedisWrapper

func GetRedisWrapper() *RedisWrapper {
	return redisWrapper
}

func (redisWrapper *RedisWrapper) Get(key string) (string, error) {
	ctx := context.Background()
	value, err := redisDb.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		return "", nil
	case err != nil:
		return "", err
	}
	return value, nil
}

func (redisWrapper *RedisWrapper) SetWithoutTTL(key string, value any) error {
	ctx := context.Background()
	err := redisDb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (redisWrapper *RedisWrapper) SetWithTTL(key string, value any, expiration time.Duration) error {
	ctx := context.Background()
	err := redisDb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (redisWrapper *RedisWrapper) Del(keys ...string) error {
	ctx := context.Background()
	err := redisDb.Del(ctx, keys...).Err()
	if err != nil {
		return err
	}
	return nil
}

// 计数器
func (redisWrapper *RedisWrapper) IncrBy(key string, value int64) (int64, error) {
	ctx := context.Background()
	result, err := redisDb.IncrBy(ctx, key, value).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

// list 头部插入
func (redisWrapper *RedisWrapper) ListLeftPush(key string, value ...any) error {
	ctx := context.Background()
	err := redisDb.LPush(ctx, key, value).Err()
	if err != nil {
		return err
	}
	return nil
}

// list 尾部删除
func (redisWrapper *RedisWrapper) ListRightPop(key string) (string, error) {
	ctx := context.Background()
	value, err := redisDb.RPop(ctx, key).Result()
	switch {
	case err == redis.Nil:
		return "", nil
	case err != nil:
		return "", err
	}
	return value, nil
}

// 获取 list 长度
func (redisWrapper *RedisWrapper) ListLen(key string) (int64, error) {
	ctx := context.Background()
	value, err := redisDb.LLen(ctx, key).Result()
	return value, err
}

type ZSetMember = redis.Z

func (redisWrapper *RedisWrapper) ZSetAdd(key string, values ...ZSetMember) (int64, error) {
	ctx := context.Background()
	value, err := redisDb.ZAdd(ctx, key, values...).Result()
	return value, err
}

func (redisWrapper *RedisWrapper) ZSetRange(key string, start, stop int64) ([]string, error) {
	ctx := context.Background()
	values, err := redisDb.ZRange(ctx, key, start, stop).Result()
	return values, err
}

func (redisWrapper *RedisWrapper) GetAllKeys(prefix string) ([]string, error) {
	ctx := context.Background()
	_prefix := ""
	if prefix == "" {
		_prefix = "*"
	} else {
		_prefix = prefix + ":*"
	}
	iter := redisDb.Scan(ctx, 0, _prefix, 0).Iterator()
	values := []string{}
	for iter.Next(ctx) {
		values = append(values, iter.Val())
	}
	return values, iter.Err()
}
