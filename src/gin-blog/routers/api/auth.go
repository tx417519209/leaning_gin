package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xing.learning.gin/src/gin-blog/models"
	"xing.learning.gin/src/gin-blog/param"
	"xing.learning.gin/src/gin-blog/pkg/e"
	"xing.learning.gin/src/gin-blog/pkg/util"
)

// @Summary 获取token
// @Accept application/json
// @Produce application/json
// @Param  account body param.AuthParam true "Account"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"ok"}"
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	var param param.AuthParam
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if err := c.Bind(&param); err == nil {
		isExist := models.CheckAuth(param.UserName, param.PassWord)
		if isExist {
			token, err := util.GenerateToken(param.UserName, param.PassWord)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_NOT_EXIST_USER
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
