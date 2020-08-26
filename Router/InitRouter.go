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

	v1 := r.Group("/api/v1")
	v1.Use(Jwt.JWT())
	{
		//管理员操作
		v1.POST("/admin", V1.GetAdmin)
		v1.POST("/adminList", V1.GetAdmins)
		v1.POST("/adminAdd", V1.AddAdmin)
		v1.POST("/login", V1.Login)
	}
	return r
}
