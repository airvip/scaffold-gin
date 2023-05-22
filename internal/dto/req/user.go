package req

import "time"

type UserBase struct {
	ID        uint      `json:"id"`        // id
	CreatedAt time.Time `json:"create_at"` // 创建时间
	UpdatedAt time.Time `json:"update_at"` // 更新时间
	Identity  string    `json:"identity"`  // 用户唯一标识
	Nickname  string    `json:"nickname" ` // 昵称
	Email     string    `json:"email" `    // 邮箱
	Mobile    string    `json:"mobile" `   // 手机
}

type AddUserDto struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,checkNickname"` // 昵称
	Password string `form:"password" json:"password" binding:"required"`               // 密码
	Email    string `form:"email" json:"email" binding:"required,email"`               // 邮箱
	Mobile   string `form:"mobile" json:"mobile" binding:"required,len=11"`            // 手机
}

type LoginUserDto struct {
	Password string `form:"password" json:"password" binding:"required"`    // 密码
	Mobile   string `form:"mobile" json:"mobile" binding:"required,len=11"` // 手机
}

type UpdateUserDto struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,checkNickname"` // 昵称
	Password string `form:"password" json:"password" `                                 // 密码
	Email    string `form:"email" json:"email" binding:"required,email"`               // 邮箱
	Mobile   string `form:"mobile" json:"mobile" binding:"required,len=11"`            // 手机
}