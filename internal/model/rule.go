package model

import (
	"gorm.io/gorm"
)

type RuleBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(255);not null;default:'';index;comment:'唯一标识'" json:"identity"` // 规则唯一标识
	ParentId int    `gorm:"column:parent_id;type:int(11);not null;default:0;index;comment:'父级id'" json:"parent_id"`     // 父级id
	RuleName  string `gorm:"column:rule_name;type:varchar(50);not null;default:'';comment:'规则名称'" json:"rule_name"`   // 规则名称
	RuleTitle string `gorm:"column:rule_title;type:varchar(50);not null;default:'';comment:'规则标题'" json:"rule_title"` // 规则标题
	Sort      int    `gorm:"column:sort;type:int(11);not null;default:0;comment:' 排序'" json:"sort"`                // 排序
}

func (*RuleBasic) TableName() string {
	return "rule_basic"
}
