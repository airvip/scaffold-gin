package model

import (
	"scaffold-gin/common/global"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(255);not null;default:'';index;comment:'唯一标识'" json:"identity"` // 用户唯一标识
	Password string `gorm:"column:password;type:varchar(255);not null;default:'';index;comment:'密码'" json:"password"` // 用户密码
	Nickname string `gorm:"column:nickname;type:varchar(50);not null;default:'';comment:'用户昵称'" json:"nickname"`        // 用户昵称
	Email    string `gorm:"column:email;type:varchar(100);not null;default:'';comment:'用户邮箱'" json:"email"`             // 邮箱
	Mobile    string `gorm:"column:mobile;type:varchar(100);not null;default:'';comment:'用户手机'" json:"mobile"`             // 手机
}

func (*UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList(nickname string) *gorm.DB{
	return global.DB.Model(new(UserBasic)).Omit("password").Where("nickname like ? OR email like ?","%"+nickname+"%","%"+nickname+"%")
}

func GetUserModel() *gorm.DB {
    return global.DB.Model(new(UserBasic))
}