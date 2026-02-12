package model

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	// 头像
	Avatar string
	// 真实姓名
	FullName string

	Email     string    `gorm:"unique;not null"`
	Phone     string    `gorm:"unique;not null"`
	Role      string    `gorm:"not null"`       // 角色，例如 "worker", "admin"
	Rate      int       `gorm:"default:0"`      // 用户评分
	CreatedAt time.Time `gorm:"autoCreateTime"` // 创建时自动写入
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
