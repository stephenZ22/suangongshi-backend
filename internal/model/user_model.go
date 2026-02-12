package model

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	// 头像
	Avatar string `json:"avatar"`
	// 真实姓名
	FullName string `json:"full_name"`

	Phone     string    `gorm:"unique;not null" json:"phone"`
	Role      string    `gorm:"not null" json:"role"`             // 角色，例如 "worker", "admin"
	Rate      int       `gorm:"default:0" json:"rate"`            // 用户评分
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"` // 创建时自动写入
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
