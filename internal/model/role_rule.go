package model

import "gorm.io/gorm"


type RoleRule struct {
	gorm.Model
	RoleId string `gorm:"column:role_id;type:int(11);not null;default:0;comment:'角色id'" json:"role_id"` // 角色id
	RuleId string `gorm:"column:rule_id;type:int(11);not null;default:0;comment:'规则id'" json:"rule_id"` // 规则id
	RuleBasic *RuleBasic `gorm:"foreignKey:ID;references:RuleId"`        // 关联规则基础信息
}

func (*RoleRule) TableName() string {
	return "role_rule"
}
