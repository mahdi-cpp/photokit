package config

import (
	"fmt"
	"path/filepath"
)

const rootDir = "/media/mahdi/happle/cloud"
const applicationDir = "com.helium.photos"
const usersDir = "users"

const versionCode = 2

func GetPath(file string) string {
	return filepath.Join(rootDir, applicationDir, file)
}

func GetUserPath(phone string, file string) string {
	pp := filepath.Join(rootDir, applicationDir, usersDir, phone, file)
	fmt.Println(pp)
	return pp
}
