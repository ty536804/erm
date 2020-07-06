package Router

import (
	"erm/Middleware/Jwt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func InitRouter() *gin.Engine {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()
	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultErrorWriter = io.MultiWriter(f, os.Stdin)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	apiV1 := r.Group("/api/v1")
	apiV1.Use(Jwt.JWT())
	{
		apiV1.GET("/powerShow")
	}
	return r
}
