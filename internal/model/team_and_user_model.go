package model

type TeamAndUser struct {
	TeamID uint `gorm:"primaryKey"`
	UserID uint `gorm:"primaryKey"`
}
