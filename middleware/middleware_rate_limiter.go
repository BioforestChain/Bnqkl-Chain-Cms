package middleware

import (
	"bnqkl/chain-cms/config"
	"bnqkl/chain-cms/exception"
	"bnqkl/chain-cms/helper"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Ban struct {
	limiterMap map[string]*rate.Limiter
	locker     sync.Mutex
	limit      rate.Limit
	burst      int
	ResetTimes time.Duration
}

func NewBan() *Ban {
	rateConfig := config.GetConfig().Rate
	limit := rate.Limit(rateConfig.Limit)
	ban := &Ban{
		limiterMap: map[string]*rate.Limiter{},
		locker:     sync.Mutex{},
		limit:      limit,
		burst:      rateConfig.Burst,
		ResetTimes: time.Second * time.Duration(rateConfig.ResetTimes),
	}
	// 定期清理过期的 ip
	go func() {
		timer := time.NewTimer(ban.ResetTimes)
		for range timer.C {
			ban.locker.Lock()
			for k, v := range ban.limiterMap {
				// 限定的时间内没有请求到达
				if v.Allow() {
					delete(ban.limiterMap, k)
				}
			}
			ban.locker.Unlock()
			timer.Reset(ban.ResetTimes)
		}
	}()
	return ban
}

func (ban *Ban) getLimiter(ip string) *rate.Limiter {
	ban.locker.Lock()
	defer ban.locker.Unlock()
	if limiter, ok := ban.limiterMap[ip]; ok {
		return limiter
	}
	ban.limiterMap[ip] = rate.NewLimiter(ban.limit, ban.burst)
	return ban.limiterMap[ip]
}

func NewRateLimiterMiddleware() gin.HandlerFunc {
	ban := NewBan()
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		limiter := ban.getLimiter(ip)
		if !limiter.Allow() {
			helper.FailureResponse(ctx, exception.NewExceptionWithoutParam(exception.TOO_MANY_REQUESTS))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
