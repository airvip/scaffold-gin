package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 authorization header
		method := c.Request.Method
		origin := c.GetHeader("Origin")
		// c.GetHeader("Origin")

		if origin != "" {
			//接收客户端发送的origin 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*") 
			// 设置预验请求有效期为 86400 秒
			c.Header("Access-Control-Max-Age", "86400")
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			//允许客户端传递校验信息比如 cookie 
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
