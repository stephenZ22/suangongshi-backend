package database

import (
	"log"
	"time"

	"github.com/stephenz22/suangongshi/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 初始化 PostgreSQL 连接
// dsn 示例: "host=localhost user=sg_admin password=sg_pass_123 dbname=suangongshi port=5432 sslmode=disable TimeZone=Asia/Shanghai"
func InitDB(dsn string) *gorm.DB {
	// 1. 配置 GORM 的日志级别（开发阶段建议开启 Info，查看生成生成的 SQL）
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// 2. 打开连接
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("无法连接到 PostgreSQL 数据库: %v", err)
	}

	// 3. 配置底层连接池 (sql.DB)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取底层 sql.DB 失败: %v", err)
	}

	// 设置最大闲置连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置连接最大复用时间（防止连接被数据库主动关闭导致报错）
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ PostgreSQL 数据库连接初始化成功 [用户: sg_admin]")

	return db
}

func MigrateDB(db *gorm.DB) error {
	// 这里可以调用 db.AutoMigrate(&User{}, &Project{}, &WorkSession{}) 来自动迁移数据库结构
	// 例如:
	// err := db.AutoMigrate(&User{}, &Project{}, &WorkSession{})
	return db.AutoMigrate(&model.User{})

}
