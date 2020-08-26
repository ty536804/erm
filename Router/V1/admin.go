package V1

import (
	"erm/Pkg/E"
	"erm/Services"
	"github.com/gin-gonic/gin"
)

// @Summer 添加/编辑用户
func AddAdmin(c *gin.Context) {
	code, msg := Services.AddAdmin(c)
	c.SecureJSON(E.SUCCESS, gin.H{
		"code": code,
		"msg":  msg,
		"data": "",
	})
}

// @Summer 获取权限列表
func GetAdmins(c *gin.Context) {
	data := Services.GetAdmins(c)
	c.SecureJSON(E.SUCCESS, gin.H{
		"code": E.SUCCESS,
		"msg":  "操作成功",
		"data": data,
	})
}

// @Summer 通过制定ID获取管理员信息
func GetAdmin(c *gin.Context) {
	data := Services.GetAdmin(c)
	c.SecureJSON(E.SUCCESS, gin.H{
		"code": E.SUCCESS,
		"msg":  "操作成功",
		"data": data,
	})
}

// @Summer 登录
func Login(c *gin.Context) {
	code, msg := Services.Login(c)
	c.SecureJSON(E.SUCCESS, gin.H{
		"code": code,
		"msg":  msg,
		"data": "",
	})
}
