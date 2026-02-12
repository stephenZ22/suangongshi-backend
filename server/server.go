package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SuanGongShiServer struct {
	HttpServer *gin.Engine
	DataBase   *gorm.DB
}

func New(db *gorm.DB, r *gin.Engine) *SuanGongShiServer {

	return &SuanGongShiServer{
		HttpServer: r,
		DataBase:   db,
	}
}

func (s *SuanGongShiServer) Run(addr string) error {
	return s.HttpServer.Run(addr)
}
