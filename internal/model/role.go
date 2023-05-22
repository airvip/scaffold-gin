package model

import (
	"scaffold-gin/common/global"

	"gorm.io/gorm"
)

type RoleBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(255);not null;default:'';comment:'角色唯一标识';index" json:"identity"` // 角色唯一标识
	RoleName string `gorm:"column:role_name;type:varchar(50);not null;default:'';comment:'角色名'" json:"role_name"`        // 角色名
	Remark string `gorm:"column:remark;type:varchar(255);not null;default:'';comment:'角色说明'" json:"remark"`        // 角色说明
    RoleRule []*RoleRule `gorm:"foreignKey:role_id;references:id"` // 关联角色规则表
}

func (*RoleBasic) TableName() string {
	return "role_basic"
}


func GetRoleList(keyword,rule_identity string) *gorm.DB{
	tx := global.DB.Model(new(RoleBasic)).Preload("RoleRule").Preload("RoleRule.RuleBasic").
	Where("role_name like ? OR remark like ?","%"+keyword+"%","%"+keyword+"%")

	if rule_identity != "" {
		tx.Joins("LEFT JOIN role_rule rr ON rr.role_id=role_basic.id").
		Where("rr.rule_id=(SELECT rub.id FROM rule_basic rub WHERE rub.identity=?)",rule_identity)
	}

	return tx
}