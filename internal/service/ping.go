package service

import "github.com/gin-gonic/gin"

// PingServer
// @Summary 服务ping
// @Schemes
// @Description 服务ping通测试
// @Tags 公用方法
// @Accept json
// @Produce json
// @Router /ping [get]
// @Security ApiKeyAuth
func PingServer(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "pong",
	})
}
