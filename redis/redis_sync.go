package redis

import (
	"bnqkl/chain-cms/logger"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

var redisSync *redsync.Redsync

func GetRedisSync() *redsync.Redsync {
	return redisSync
}

func InitRedisSync(log *logger.Logger) {
	pool := goredis.NewPool(redisDb)
	_redisSync := redsync.New(pool)
	redisSync = _redisSync
	log.Info("init redis sync server success")
}

func formatLockKey(lockKey string) string {
	return "lock:" + lockKey
}

func DoWithLock[T any](lockKey string, task func() (T, error)) (T, error) {
	var r T
	mutex := redisSync.NewMutex(formatLockKey(lockKey), redsync.WithExpiry(time.Second*30))
	if err := mutex.Lock(); err != nil {
		return r, err
	}
	r, err := task()
	// FIXME：解锁失败目前无解，屌大的处理一下，设置一下自动过期？？
	mutex.Unlock()
	return r, err
}

func DoWithLockMulti[T any](lockKeys []string, task func() (T, error)) (T, error) {
	var r T
	mutexs, err := lockMulti(lockKeys)
	if err != nil {
		unlockMulti(mutexs)
		return r, err
	}
	r, err = task()
	unlockMulti(mutexs)
	return r, err
}

func lockMulti(lockKeys []string) ([]*redsync.Mutex, error) {
	mutexs := []*redsync.Mutex{}
	expire := time.Second * 30 * time.Duration(len(lockKeys))
	for _, lockKey := range lockKeys {
		mutex := redisSync.NewMutex(formatLockKey(lockKey), redsync.WithExpiry(expire))
		err := mutex.Lock()
		if err != nil {
			return mutexs, err
		}
		mutexs = append(mutexs, mutex)
	}
	return mutexs, nil
}

func unlockMulti(mutexs []*redsync.Mutex) {
	for _, mutex := range mutexs {
		mutex.Unlock()
	}
}
