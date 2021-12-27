package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"xing.learning.gin/src/gin-blog/pkg/file"
	"xing.learning.gin/src/gin-blog/pkg/logging"
	"xing.learning.gin/src/gin-blog/pkg/setting"
	"xing.learning.gin/src/gin-blog/pkg/util"
)

func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + setting.AppSetting.ImageSavePath + name
}

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImageName(name string) string {
	ext := file.GetExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {

		return fmt.Errorf("file.CheckPermission Permission denied src: %s",
			src)
	}

	return nil
}
