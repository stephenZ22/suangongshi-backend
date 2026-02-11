package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SuanGongShiServer struct {
	HttpServer *gin.Engine
	DataBase   *gorm.DB
}
