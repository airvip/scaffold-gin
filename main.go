package main

import (
	"io"
	"os"
	"scaffold-gin/common/config"

	// "scaffold-gin/common/global"
	// "scaffold-gin/internal/model"
	"scaffold-gin/internal/router"

	"github.com/gin-gonic/gin"
)

var c string

// @title scaffold-gin API
// @version v1
// @description This is a gin scaffold
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {


	gin.SetMode(config.Conf.Server.RunMode)
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
    gin.DisableConsoleColor()

	// 如果开启数据库迁移记得开启上面的 import
	// db := global.DB
	// db.AutoMigrate(&model.UserBasic{}, &model.RoleBasic{}, &model.RuleBasic{}, &model.RoleRule{})


	// 记录到文件。
    f, _ := os.Create("log/gin.log")
    // gin.DefaultWriter = io.MultiWriter(f)
    // 如果需要同时将日志写入文件和控制台，请使用以下代码。
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := router.Router()
	r.Run(config.Conf.Server.Port)

}
