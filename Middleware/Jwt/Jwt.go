package Jwt

import (
	"erm/Pkg/E"
	"erm/Pkg/Util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = E.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = E.INVALID_PARAMS
		} else {
			claims, err := Util.ParseToken(token)
			if err != nil {
				code = E.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = E.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}

		if code != E.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  E.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
