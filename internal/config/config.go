package config

import (
	"fmt"
	"path/filepath"
)

const (
	root        = "/app/iris/"
	application = "com.iris.photos"
	users       = "users"
	Metadata    = "metadata"
	Version     = "v3"
)

func GetPath(file string) string {
	return filepath.Join(root, application, file)
}

func GetUserPath(userID string) string {
	pp := filepath.Join(root, application, users, userID)
	fmt.Println(pp)
	return pp
}

func GetUserMetadataPath(id string, directory string) string {
	pp := filepath.Join(root, application, users, id, Metadata, Version, directory)
	fmt.Println(pp)
	return pp
}
