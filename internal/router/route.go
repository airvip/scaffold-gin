package router

import (
	"scaffold-gin/common/middleware"
	"scaffold-gin/internal/service"
	"scaffold-gin/internal/service/pub"

	"github.com/gin-gonic/gin"

	_ "scaffold-gin/docs" // swag init 产生的 docs
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)



func Router() *gin.Engine {
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CorsMiddleware(), middleware.GinLogMiddleware())
	// 配置路由

	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/ping", middleware.AuthMiddleware(), service.PingServer)
	
	r.GET("/sms-code-tx", pub.SmsTxCode)
	r.GET("/sms-code-ali", pub.SmsAliCode)
	r.POST("/upload", pub.UploadObj)

	r.POST("/user-register", service.AddUser)
	r.POST("/user-login", service.LoginUser)
	

	api := r.Group("v1")
	{
		api.GET("/user-detail", service.GetUserDetail)
		api.GET("/user-list", service.GetUserList)
		api.PUT("/user-update", service.UpdateUser)
		api.GET("/role-list", service.GetRoleList)
		api.GET("/role-detail", service.GetRoleDetail)
	}
	return r
}
