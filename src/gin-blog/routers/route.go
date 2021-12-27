package routers

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "xing.learning.gin/src/gin-blog/docs"
	"xing.learning.gin/src/gin-blog/middleware/jwt"
	"xing.learning.gin/src/gin-blog/pkg/setting"
	"xing.learning.gin/src/gin-blog/routers/api"
	v1 "xing.learning.gin/src/gin-blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.Static("/html", "./public")
	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		/*---------------------------------*/
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:article_id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:article_id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
