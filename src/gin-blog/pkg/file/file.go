package file

import (
	"io"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(f multipart.File) (int64, error) {

	return f.Seek(0, io.SeekEnd)

}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

func CheckPermission(path string) bool {
	_, err := os.Stat(path)
	return os.IsPermission(err)

}
func MkDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil

}

func IsNotExistMkDir(src string) error {
	if exist := CheckExist(src); !exist {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
