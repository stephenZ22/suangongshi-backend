package repository

import (
	"github.com/stephenz22/suangongshi/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetAllUsers() (*[]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return &users, err
}

func (r *UserRepository) GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, userID).Error
	return &user, err
}
