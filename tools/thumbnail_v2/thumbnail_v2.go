package thumbnail_v2

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

const (
	thumbnailSize = 270
	thumbnailsDir = "thumbnails"
)

// ImageMetadata struct to hold the extracted EXIF data.
// ساختار ImageMetadata برای نگهداری داده‌های EXIF استخراج شده.
type ImageMetadata struct {
	FileSize    string
	FileType    string
	Make        string
	Model       string
	Orientation string
	CreateDate  string
}

// Create generates thumbnails for all JPEG images in a user's asset directory
func Create(userID string) error {

	basePath := filepath.Join("/app/iris/com.iris.photos/users", userID, "assets")
	thumbPath := filepath.Join(basePath, thumbnailsDir)

	// Create thumbnails directory if it doesn't exist
	if err := os.MkdirAll(thumbPath, 0755); err != nil {
		return fmt.Errorf("failed to create thumbnails directory: %w", err)
	}

	entries, err := os.ReadDir(basePath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		if !isJPEGFile(entry) {
			continue
		}

		if err := processImage(basePath, thumbPath, entry.Name()); err != nil {
			fmt.Printf("Error processing %s: %v\n", entry.Name(), err)
		}
	}

	return nil
}

func isJPEGFile(entry os.DirEntry) bool {
	return !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".jpg")
}

func processImage(basePath, thumbPath, filename string) error {

	srcPath := filepath.Join(basePath, filename)
	thumbName := generateThumbnailName(filename)
	dstPath := filepath.Join(thumbPath, thumbName)

	// Skip if thumbnail already exists
	//if _, err := os.Stat(dstPath); err == nil {
	//	return nil
	//}

	srcImage, err := imaging.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}

	var portrait = false
	//if thumbnail.PhotoHasExifData(srcPath) {
	//
	//	has, orientation := thumbnail.ReadExifData(srcPath)
	//
	//	if has {
	//		fmt.Println("Orientation: ", orientation)
	//		if strings.Compare(orientation, "6") == 0 {
	//			portrait = true
	//		}
	//
	//		i, err := strconv.Atoi(orientation)
	//		if err != nil {
	//			fmt.Println("Orientation: ", err)
	//		} else {
	//			//Orientation = i
	//			fmt.Println(i)
	//		}
	//	}
	//}

	// Get orientation from EXIF data
	//orientation := getOrientation(srcPath)
	//srcImage = applyOrientation(srcImage, orientation)

	// Calculate the new dimensions while maintaining aspect ratio
	//bounds := srcImage.Bounds()
	//width, height := bounds.Dx(), bounds.Dy()

	var dstImage *image.NRGBA

	if portrait {
		// Resize the cropped image to width = 200px preserving the aspect ratio.
		dstImage = imaging.Resize(srcImage, 0, thumbnailSize, imaging.Lanczos)
		dstImage = imaging.Rotate270(dstImage)

	} else {
		// Resize the cropped image to width = 200px preserving the aspect ratio.
		dstImage = imaging.Resize(srcImage, thumbnailSize, 0, imaging.Lanczos)
	}

	err2 := imaging.Save(dstImage, dstPath)
	if err2 != nil {
		panic(err2)
	}

	//var thumbnail image.Image
	//
	//// Determine if image is portrait or landscape
	//if width > height {
	//	// Landscape: resize based on width
	//	thumbnail = imaging.Resize(srcImage, thumbnailSize, 0, imaging.Lanczos)
	//} else {
	//	// Portrait: resize based on height
	//	thumbnail = imaging.Resize(srcImage, 0, thumbnailSize, imaging.Lanczos)
	//}
	//
	//// Save thumbnail
	//if err := imaging.Save(thumbnail, dstPath); err != nil {
	//	return fmt.Errorf("failed to save thumbnail: %w", err)
	//}

	return nil
}

func generateThumbnailName(filename string) string {
	return strings.TrimSuffix(filename, ".jpg") + "_270.jpg"
}

//func getOrientation(filePath string) int {
//
//	f, err := os.Open(filePath)
//	if err != nil {
//		return 1
//	}
//	defer f.Close()
//
//	// Get the EXIF data
//	data, err := os.ReadFile(filePath)
//	if err != nil {
//		return 1
//	}
//
//	rawExif, err := exif.SearchAndExtractExif(data)
//	if err != nil {
//		return 1
//	}
//
//	// Parse the EXIF data
//	im := exifcommon.NewIfdMapping()
//	ti := exif.NewTagIndex()
//
//	// Collect all EXIF data
//	_, index, err := exif.Collect(im, ti, rawExif)
//	if err != nil {
//		return 1
//	}
//
//	// Look for orientation tag in the main IFD
//	rootIfd := index.RootIfd
//	if rootIfd == nil {
//		return 1
//	}
//
//	// Try to find the orientation tag
//	results, err := rootIfd.FindTagWithName("Orientation")
//	if err != nil || len(results) == 0 {
//		// Also check the EXIF IFD
//		if exifIfd := index.Lookup["IFD/Exif"]; exifIfd != nil {
//			results, err = exifIfd.FindTagWithName("Orientation")
//			if err != nil || len(results) == 0 {
//				return 1
//			}
//		} else {
//			return 1
//		}
//	}
//
//	//// Get the orientation value
//	//if orientation, err := strconv.Atoi(results[0].String()); err == nil {
//	//	return orientation
//	//}
//
//	// Get the orientation value
//	tagValue := fmt.Sprintf("%v", results[0].Value())
//	if orientation, err := strconv.Atoi(tagValue); err == nil {
//		return orientation
//	}
//
//	return 1
//}
//
//func applyOrientation(img image.Image, orientation int) image.Image {
//	switch orientation {
//	case 2:
//		return imaging.FlipH(img)
//	case 3:
//		return imaging.Rotate180(img)
//	case 4:
//		return imaging.FlipV(img)
//	case 5:
//		return imaging.Transpose(img)
//	case 6:
//		return imaging.Rotate270(img)
//	case 7:
//		return imaging.Transverse(img)
//	case 8:
//		return imaging.Rotate90(img)
//	default:
//		return img
//	}
//}
