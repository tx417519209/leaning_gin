package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"xing.learning.gin/src/gin-blog/pkg/e"
	"xing.learning.gin/src/gin-blog/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		var data interface{}
		token, ok := c.GetQuery("token")
		if !ok {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
