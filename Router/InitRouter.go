package Router

import (
	"erm/Middleware/Jwt"
	"erm/Router/V1"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func InitRouter() *gin.Engine {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()
	// 记录到文件。
	f, _ := os.Create(".air.conf")
	gin.DefaultErrorWriter = io.MultiWriter(f, os.Stdin)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	apiV1 := r.Group("/api/v1")
	apiV1.Use(Jwt.JWT())
	{
		//管理员操作
		r.GET("/admin", V1.GetAdmin)
		apiV1.GET("/adminList", V1.GetAdmins)
		apiV1.GET("/adminAdd", V1.AddAdmin)
	}
	return r
}
