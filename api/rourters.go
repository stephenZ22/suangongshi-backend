package api

import (
	"github.com/stephenz22/suangongshi/api/handler"
	"github.com/stephenz22/suangongshi/internal/repository"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// todo: fix imports errors
func RegisterRouters(db *gorm.DB) *gin.Engine {
	// 基础健康检查，确认后端活着
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// API 版本控制
	v1 := r.Group("/api/v1")
	{

		// --- 工时/记录模块 ---
		// recordRoutes := v1.Group("/records")
		// {
		// 	// 这里的接口名要直观
		// 	recordRoutes.POST("/upsert", handler.UpsertWorkRecord) // 记一笔/改一笔
		// 	recordRoutes.GET("/month", handler.GetMonthSummary)    // 按月统计
		// 	recordRoutes.GET("/list", handler.GetRecordList)       // 历史列表
		// }

		// --- 用户模块 ---
		userRoutes := v1.Group("/users")
		{
			userRepo := repository.NewUserRepository(db)
			user_handler := handler.NewUsersHandler(userRepo)
			userRoutes.GET("/profile", user_handler.GetUserProfile)
			userRoutes.POST("/update", user_handler.UpdateUserRate) // 更新默认工价
		}
	}

	return r
}
