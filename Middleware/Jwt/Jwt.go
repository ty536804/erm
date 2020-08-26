package Jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var code int
		//var data interface{}
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
		//
		//code = E.SUCCESS
		//token := c.Query("token")
		//if token == "" {
		//	code = E.INVALID_PARAMS
		//} else {
		//	claims, err := Util.ParseToken(token)
		//	if err != nil {
		//		code = E.ERROR_AUTH_CHECK_TOKEN_FAIL
		//	} else if time.Now().Unix() > claims.ExpiresAt {
		//		code = E.ERROR_AUTH_CHECK_TOKEN_FAIL
		//	}
		//}
		//
		//if code != E.SUCCESS {
		//	c.JSON(http.StatusUnauthorized, gin.H{
		//		"code": code,
		//		"msg":  E.GetMsg(code),
		//		"data": data,
		//	})
		//
		//	c.Abort()
		//	return
		//}
		//
		//c.Next()
	}
}
