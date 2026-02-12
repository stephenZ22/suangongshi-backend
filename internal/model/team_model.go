package model

type Team struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"unique;not null" json:"name"`
	LeaderID uint   `gorm:"not null" json:"leader_id"` // 负责人用户ID
}
