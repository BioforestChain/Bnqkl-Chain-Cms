package middleware

import (
	"bnqkl/chain-cms/logger"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func NewApiTimmerMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 计算请求时间
		start := time.Now()
		uri := ctx.Request.URL.Path
		// 调用后续处理的函数
		ctx.Next()
		end := time.Since(start)
		if end > time.Second/2 {
			log.Warn(fmt.Sprintf("api: %s => %d", uri, end))
		}
	}
}
