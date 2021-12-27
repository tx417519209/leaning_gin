package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xing.learning.gin/src/gin-blog/pkg/e"
	"xing.learning.gin/src/gin-blog/pkg/logging"
	"xing.learning.gin/src/gin-blog/pkg/upload"
)

func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}
	if header == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(header.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()
		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(header, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
