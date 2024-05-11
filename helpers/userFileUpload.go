package helpers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strings"
	"userapp/initializers"
)

func GetUserFilePath(userName string) string {
	path := fmt.Sprintf("%s/%s", initializers.BasePath, userName)

	return path
}

func IsFileImage(file *multipart.FileHeader) error {
	// Check Content-Type Header
	contentType := strings.Split(file.Header.Get("Content-Type"), "/")
	if strings.Compare(contentType[0], "image") != 0 {
		return errors.New("file must be an image type (jpg,jpeg,png)")
	}
	return nil
}

func GetFileExtension(fileName string) string {
	ext := strings.Split(fileName, ".")
	return ext[len(ext)-1]
}

func SetFileName(fileName string) string {
	newFileName := GenerateRandomString() + "." + GetFileExtension(fileName)
	return newFileName
}
