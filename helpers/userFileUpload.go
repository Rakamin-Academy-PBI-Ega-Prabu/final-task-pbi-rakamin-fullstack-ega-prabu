package helpers

import (
	"fmt"
	"strings"
	"userapp/initializers"
)

func GetUserFilePath(userName string) string {
	path := fmt.Sprintf("%s/%s", initializers.BasePath, userName)

	return path
}

func IsFileImage(FileHeader map[string]string) (bool, string) {
	// Check Content-Type Header
	contentType := strings.Split(FileHeader["Content-Type"], "/")
	if strings.Compare(contentType[0], "image") != 0 {
		return false, ""
	}
	return true, contentType[1]
}

func GetFileExtension(fileName string) string {
	ext := strings.Split(fileName, ".")
	return ext[1]
}

func SetFileName(fileName string) string {
	newFileName := GenerateRandomString() + "." + GetFileExtension(fileName)
	return newFileName
}
