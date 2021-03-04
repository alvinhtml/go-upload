package router

import (
	"alvinhtml.com/go-upload/file"
	"alvinhtml.com/go-upload/middleware"
	"github.com/gin-gonic/gin"
)

// Init 初始化总路由
func Init() *gin.Engine {
	Router := gin.Default()

	// 跨域
	Router.Use(middleware.Cors())

	ApiGroup := Router.Group("/api")

	ApiGroup.PUT("/cors", file.TestCors)
	ApiGroup.PUT("/upload/:md5/:index", file.Upload)
	ApiGroup.POST("/makefile/:md5/:chunkTotal", file.ExecMerge)

	return Router
}
