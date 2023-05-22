package global

import (
	"fmt"
	"log"
	"scaffold-gin/common/config"
	"scaffold-gin/util"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB = InitDB()

func InitDB() *gorm.DB {
	// 初始化配置文件
	user := config.Conf.DB.User
	pass := config.Conf.DB.Pass
	host := config.Conf.DB.Host
	port := config.Conf.DB.Port
	dbname := config.Conf.DB.DbName
	charset := config.Conf.DB.Charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, pass, host, port, dbname, charset)
	// log.Println(dsn)

	// logger.Warn 只打印慢查询 默认的SlowThreshold为200ms
	// logger.Info 打印所有sql

	newLogger := logger.New(
		// log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		log.New(util.Logger, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Warn,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,                  // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Println("mysql Init Error", err)
		panic(err)
	}
	return db
}
