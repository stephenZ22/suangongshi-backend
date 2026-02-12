package repository

import (
	"time"

	"github.com/stephenz22/suangongshi/internal/model"
	"gorm.io/gorm"
)

type WorklogsRepository struct {
	db *gorm.DB
}

func NewWorklogsRepository(db *gorm.DB) *WorklogsRepository {
	return &WorklogsRepository{db: db}
}

func (r *WorklogsRepository) CreateWorklog(worklog *model.WorkLog) error {
	return r.db.Create(worklog).Error
}

func (r *WorklogsRepository) GetAllWorklogs() (*[]model.WorkLog, error) {
	var worklogs []model.WorkLog
	err := r.db.Find(&worklogs).Error
	return &worklogs, err
}

func (r *WorklogsRepository) GetWorklogsByUserID(userID uint) (*[]model.WorkLog, error) {
	var worklogs []model.WorkLog
	err := r.db.Where("user_id = ?", userID).Find(&worklogs).Error
	return &worklogs, err
}

func (r *WorklogsRepository) GetMonthWorklogs(userID uint, month int, year int) (*[]model.WorkLog, error) {
	var worklogs []model.WorkLog
	var startTime, endTime time.Time

	// 按月查询：2026-02-01 到 2026-03-01 (不含)
	startTime = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endTime = startTime.AddDate(0, 1, 0)
	err := r.db.Where("user_id = ? AND work_date >= ? AND work_date < ?", userID, startTime, endTime).Find(&worklogs).Error
	return &worklogs, err
}

func (r *WorklogsRepository) GetYearWorklogs(userID uint, year int) (*[]model.WorkLog, error) {
	var worklogs []model.WorkLog
	var startTime, endTime time.Time
	startTime = time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	endTime = startTime.AddDate(1, 0, 0)
	err := r.db.Where("user_id = ? AND work_date >= ? AND work_date < ?", userID, startTime, endTime).Find(&worklogs).Error
	return &worklogs, err
}
