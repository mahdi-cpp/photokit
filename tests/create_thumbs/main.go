package main

import (
	"fmt"
	"github.com/mahdi-cpp/photokit/tools/exiftool_v1"
)

func main() {
	//thumbnail_v2.Create("018fe65d-8e4a-74b0-8001-c8a7c29367e1")

	metadata, err := exiftool_v1.Start("/media/mahdi/happle/cloud/com.helium.photos/users/assets/0198c111-0fe1-7e2d-b38b-62b1a1d89907.jpg")
	if err != nil {
		return
	}

	fmt.Println("\nExtracted Metadata:")
	fmt.Printf("  File Size: %s\n", metadata.FileSize)
	fmt.Printf("  File Type: %s\n", metadata.FileType)
	fmt.Printf("  Make: %s\n", metadata.Make)
	fmt.Printf("  Model: %s\n", metadata.Model)
	fmt.Printf("  Orientation: %s\n", metadata.Orientation)
	fmt.Printf("  Create Date: %s\n", metadata.CreateDate)
}
