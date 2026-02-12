package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stephenz22/suangongshi/internal/model"
	"github.com/stephenz22/suangongshi/internal/repository"
)

type UsersHandler struct {
	user_repo     *repository.UserRepository
	worklogs_repo *repository.WorklogsRepository
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func NewUsersHandler(user_repo *repository.UserRepository, worklogs_repo *repository.WorklogsRepository) *UsersHandler {
	return &UsersHandler{user_repo: user_repo, worklogs_repo: worklogs_repo}
}
func (h *UsersHandler) CreateUser(c *gin.Context) {
	// User creation logic goes here

	req := CreateUserRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &model.User{
		Username: req.Username,
		Phone:    req.Phone,
		Password: req.Password, // In production, hash the password before storing
		Role:     "worker",     // Default role
	}

	if err := h.user_repo.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    500,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"code":    200,
		"message": "User created successfully",
	})

}

type GetMonthWorkCountRequest struct {
	Month int `json:"month" binding:"required"`
	Year  int `json:"year" binding:"required"`
}

func (h *UsersHandler) GetMonthWorkCountByUserID(c *gin.Context) {
	// User work count logic goes here
	// /month_work_count/:user_id"
	user_id := c.Param("user_id")
	var uid uint
	_, err := fmt.Sscan(user_id, &uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	req := GetMonthWorkCountRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs, err := h.worklogs_repo.GetMonthWorklogs(uid, req.Month, req.Year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve work logs"})
		return
	}
	count := len(*logs)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"work_count": count,
			"logs":       logs,
		},
		"code":    200,
		"message": "Work count retrieved successfully",
	})

}
func (h *UsersHandler) GetUserProfile(c *gin.Context) {
	// User creation logic goes here

	c.JSON(http.StatusOK, gin.H{"message": "User profile retrieved successfully"})
}

func (h *UsersHandler) UpdateUserRate(c *gin.Context) {
	// User retrieval logic goes here

	c.JSON(http.StatusOK, gin.H{"message": "User rate updated successfully"})
}

func (h *UsersHandler) UpdateUser() {
	// User update logic goes here
}

func (h *UsersHandler) DeleteUser() {
	// User deletion logic goes here
}
