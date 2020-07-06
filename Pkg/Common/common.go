package Common

import (
	"crypto/md5"
	"encoding/hex"
	"erm/Pkg/E"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"os"
	"regexp"
	"strings"
)

// @Summary md5加密 16进制32位
// @param data string 密码
func Md5Pwd(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	return hex.EncodeToString(md5.Sum(nil))
}

// @Summary 获取绝对路径
func GetDir() string {
	dir, _ := os.Getwd()
	return dir
}

// @Summary 去除两侧空白
func Trim(con string) string {
	return strings.TrimSpace(con)
}

// @Summary 返回错误内容
func Error(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(E.ERROR, gin.H{
		"code": E.ERROR,
		"msg":  msg,
		"data": data,
	})
}

// @Summary 返回正确内容
func Success(ctx *gin.Context, msg string, data interface{}) {
	ctx.SecureJSON(E.SUCCESS, gin.H{
		"code": E.SUCCESS,
		"msg":  msg,
		"data": data,
	})
}

// @Summer 解析错误原因
func ViewErr(c *gin.Context, valid validation.Validation) {
	for _, err := range valid.Errors {
		Error(c, err.Message, "")
	}
	Success(c, "操作成功", "")
}

// @Summer 手机号校验
func CheckMobile(mobile string) (err error, msg string) {
	res, _ := regexp.MatchString(`^1(\d){11}`, mobile)
	if len(mobile) < 11 && res {
		return nil, "手机号码合法"
	}
	return err, "手机号码不正确"
}
