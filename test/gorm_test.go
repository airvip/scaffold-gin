package test

import (
	"fmt"
	"scaffold-gin/internal/model"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)



func TestGorm(t *testing.T) {
	// dsn := "root:123456@tcp(127.0.0.1:3307)/v_record?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(192.168.226.136:3307)/v_record?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: false,
		// Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		t.Fatal(err)
	}
	// 表不存在新建
	db.AutoMigrate(&model.RoleBasic{},&model.RoleRule{})

	list := make([]*model.UserBasic, 0)
	err = db.Find(&list).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range list {
		fmt.Println(v)
	}
}