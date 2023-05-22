package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GinLogMiddleware() gin.HandlerFunc {

	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		// 自定义自己的日志格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			// params.TimeStamp.Format(time.RFC3339),
			params.TimeStamp.Format("2006-01-02 15:04:05"),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	})
}
