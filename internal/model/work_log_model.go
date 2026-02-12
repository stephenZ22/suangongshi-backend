package model

import "time"

type WorkLog struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	UserID   uint      `gorm:"not null;index" json:"user_id"` // 关联的用户ID
	TeamID   *uint     `gorm:"index" json:"team_id"`          // 关联的团队ID
	WorkDate time.Time `gorm:"type:date;not null;index" json:"work_date"`
	Hours    int       `gorm:"not null;default:8" json:"hours"` // 工作小时数
	Notes    string    `json:"notes"`                           // 备注
}
