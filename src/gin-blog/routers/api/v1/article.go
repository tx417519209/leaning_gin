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

// @Summary 获取指定文章
// @Produce json
// @Param token query  string  true "Token"
// @Param id query int true "Article ID"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"ok"}"
// @Router /api/v1/articles/{id} [get]
func GetArticle(c *gin.Context) {

	var param param.GetArticleParam
	var data interface{}
	code := e.SUCCESS
	if err := c.BindUri(&param); err != nil {
		code = e.INVALID_PARAMS
	} else {
		if models.ExistArticleByID(param.ID) {
			data = models.GetArticle(param.ID)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary 获取全部文章
// @Produce json
// @Param token query  string  true "Token"
// @Param id query int false "Tag ID"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"ok"}"
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	var param param.GetArticlesParam
	code := e.SUCCESS

	if err := c.BindQuery(&param); err != nil {
		code = e.INVALID_PARAMS
	} else {
		if param.State != nil {
			maps["state"] = *(param.State)
		}
		if param.TagID != nil {
			if !models.ExistTagByID(*param.TagID) {
				code = e.ERROR_NOT_EXIST_TAG
			} else {
				maps["tag_id"] = *(param.TagID)
			}
		}
		data["lists"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// @Summary 添加文章
// @Produce json
// @Param token query  string  true "Token"
// @Param info body param.ADDArticleParam true "Article Info"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context) {
	code := e.SUCCESS
	var param param.ADDArticleParam
	if err := c.Bind(&param); err != nil {
		code = e.INVALID_PARAMS
	} else {
		if models.ExistTagByID(param.TagID) {
			data := gin.H{
				"tag_id":     param.TagID,
				"title":      param.Title,
				"desc":       param.Desc,
				"content":    param.Content,
				"created_by": param.CreatedBy,
				"state":      param.State,
			}
			models.AddArticle(data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// @Summary 修改文章
// @Produce json
// @Param id path int true "ArticleID"
// @Param token query  string  true "Token"
// @Param info body param.EditArticleParam true "Article Info"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"ok"}"
// @Router /api/v1/articles/{id} [put]
func EditArticle(c *gin.Context) {
	code := e.SUCCESS
	id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		if id < 0 {
			code = e.ERROR_NOT_EXIST_ARTICLE
		} else {
			var param param.EditArticleParam
			if err := c.Bind(&param); err != nil {
				code = e.INVALID_PARAMS
			} else {
				if models.ExistArticleByID(id) && models.ExistTagByID(param.TagID) {
					datas := map[string]interface{}{
						"tag_id":      param.TagID,
						"title":       param.Title,
						"desc":        param.Desc,
						"content":     param.Content,
						"modified_by": param.ModifiedBy,
						"state":       param.State,
					}
					models.EditArticle(id, datas)
				} else {
					code = e.ERROR_NOT_EXIST_ARTICLE
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// @Summary 删除文章
// @Produce json
// @Param id path int true "Article ID"
// @Param token query  string  true "Token"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"ok"}"
// @Failure 200 {string} json "{"code":200,"data":{},"msg":"error"}""
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
	code := e.SUCCESS
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		if id < 0 {
			code = e.ERROR_NOT_EXIST_ARTICLE
		} else {

			models.DeleteArticle(id)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
