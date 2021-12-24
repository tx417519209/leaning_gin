package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"xing.learning.gin/src/gin-blog/models"
	"xing.learning.gin/src/gin-blog/param"
	"xing.learning.gin/src/gin-blog/pkg/e"
	"xing.learning.gin/src/gin-blog/pkg/setting"
	"xing.learning.gin/src/gin-blog/pkg/util"
)

// @Summary 获取全部tag
// @Produce json
// @Param name query string false "TagName"
// @Param state query int false "TagState"
// @Param token query string  true "Token"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"ok"}"
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	var param param.GetTagsParam
	c.Bind(&param)
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	maps["state"] = 1
	if param.Name != nil {
		maps["name"] = *(param.Name)
	}
	if param.State != nil {
		maps["state"] = *(param.State)
	}
	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data})
}

// @Summary 添加tag
// @Produce json
// @Param token query string  true "Token"
// @Param "info" body param.AddTagParam true "Tag Info"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"ok"}"
// @Failure 200 {string} json "{"code":200,"data":{},"msg":"error"}""
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	var param param.AddTagParam
	code := e.SUCCESS
	if err := c.Bind(&param); err != nil {
		code = e.INVALID_PARAMS
	} else {
		if !models.ExistTagByName(param.Name) {
			models.AddTag(param.Name, param.State, param.CreatedBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary 修改tag
// @Produce json
// @Param id path int true "TagID"
// @Param token query string  true "Token"
// @Param "info" body param.EditTagParam true "Tag Info"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	code := e.SUCCESS
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		if id < 0 {
			code = e.INVALID_PARAMS
		} else {
			var param param.EditTagParam
			if err := c.Bind(&param); err != nil {
				code = e.INVALID_PARAMS
			} else {
				if models.ExistTagByID(id) {
					datas := make(map[string]interface{})
					datas["modified_by"] = param.ModifiedBy
					if param.Name != nil {
						datas["name"] = *(param.Name)
					}
					if param.State != nil {
						datas["state"] = *(param.State)
					}
					models.EditTag(id, datas)
				} else {
					code = e.ERROR_NOT_EXIST_TAG
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary 删除tag
// @Produce json
// @Param id path int true "Tag ID"
// @Param token query string  true "Token"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"ok"}"
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	var param param.DeleteTagParam
	code := e.SUCCESS
	if err := c.BindUri(&param); err != nil {
		code = e.INVALID_PARAMS
	} else {
		if models.ExistTagByID(param.ID) {
			models.DeleteTag(param.ID)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
