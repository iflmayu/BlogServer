package domain

import "BlogServer/internal/common/domain"

type User struct {
	domain.BaseModel
	Username string   `gorm:"size:32;uniqueIndex;not null" json:"username"`
	Password string   `gorm:"size:128;not null" json:"-"` // 密码不序列化到 JSON
	Nickname string   `gorm:"size:32" json:"nickname"`
	Avatar   string   `gorm:"size:256" json:"avatar"`
	Email    string   `gorm:"size:256;uniqueIndex" json:"email"`
	Role     RoleType `gorm:"default:2;not null" json:"role"` //1管理员 2普通用户 3游客
}
