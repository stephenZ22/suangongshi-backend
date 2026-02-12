package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stephenz22/suangongshi/internal/model"
	"github.com/stephenz22/suangongshi/internal/repository"
)

type WorklogsHandler struct {
	repo *repository.WorklogsRepository
}

func NewWorklogsHandler(repo *repository.WorklogsRepository) *WorklogsHandler {
	return &WorklogsHandler{repo: repo}
}

const DATELAYOUT = "2006-01-02"

type CreateWorklogRequest struct {
	UserID   uint   `json:"user_id" binding:"required"`
	TeamID   *uint  `json:"team_id"`
	WorkDate string `json:"work_date" binding:"required"` // Expected in "YYYY-MM-DD" format
	Hours    int    `json:"hours"`
	Notes    string `json:"notes"`
}

func (h *WorklogsHandler) CreateWorklog(c *gin.Context) {
	var req CreateWorklogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create the worklog model
	date, _ := time.Parse(DATELAYOUT, req.WorkDate)
	worklog := &model.WorkLog{
		UserID:   req.UserID,
		TeamID:   req.TeamID,
		WorkDate: date,
		Hours:    req.Hours,
		Notes:    req.Notes,
	}

	if err := h.repo.CreateWorklog(worklog); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create worklog"})
		return
	}

	c.JSON(201, gin.H{"message": "Worklog created successfully"})
}

func (h *WorklogsHandler) GetWorklogsByUserID(c *gin.Context) {
	userIDParam := c.Param("user_id")
	var userID uint
	_, err := fmt.Sscan(userIDParam, &userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	worklogs, err := h.repo.GetWorklogsByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve worklogs"})
		return
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"worklogs": worklogs,
			"count":    len(*worklogs),
		},
		"code":    200,
		"message": "Worklogs retrieved successfully",
	})
}
