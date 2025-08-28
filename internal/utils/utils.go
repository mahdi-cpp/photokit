package utils

import (
	"os"
)

func GetFileSize(filepath string) (int64, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func GetBoolPointer(b bool) *bool {
	return &b
}

func CheckVersionIsUpToDate(versionCode int) bool {
	return true
}
