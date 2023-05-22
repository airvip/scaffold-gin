package req

import "time"

type RoleBasic struct {
	ID        uint        `json:"id"`        // id
	CreatedAt time.Time   `json:"create_at"` // 创建时间
	UpdatedAt time.Time   `json:"update_at"` // 更新时间
	Identity  string      `json:"identity"`  // 角色唯一标识
	RoleName  string      `json:"role_name"` // 角色名
	Remark    string      `json:"remark"`    // 角色说明
	RoleRule  []*RoleRule `json:"role_rule"` // 关联角色规则表
}

type RoleRule struct {
	ID        uint       `json:"id"`         // id
	CreatedAt time.Time  `json:"create_at"`  // 创建时间
	UpdatedAt time.Time  `json:"update_at"`  // 更新时间
	RoleId    string     `json:"role_id"`    // 角色id
	RuleId    string     `json:"rule_id"`    // 规则id
	RuleBasic *RuleBasic `json:"rule_basic"` // 关联规则基础信息
}

type RuleBasic struct {
	ID        uint      `json:"id"`         // id
	CreatedAt time.Time `json:"create_at"`  // 创建时间
	UpdatedAt time.Time `json:"update_at"`  // 更新时间
	Identity  string    `json:"identity"`   // 规则唯一标识
	ParentId  int       `json:"parent_id"`  // 父级id
	RuleName  string    `json:"rule_name"`  // 规则名称
	RuleTitle string    `json:"rule_title"` // 规则标题
	Sort      int       `json:"sort"`       // 排序
}