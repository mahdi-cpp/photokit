package upgrade_v3

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mahdi-cpp/photokit/internal/collections/phasset"
	"github.com/mahdi-cpp/photokit/tools/exiftool"
	"github.com/mahdi-cpp/photokit/tools/exiftool_v1"
)

func upgradePHAssets(userID string) error {

	const directoryName = "assets"

	// Read all directory entries
	files, err := os.ReadDir(assetsDir) // Use os.ReadDir(dirPath) for Go 1.16+
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	var images []string

	for _, file := range files {
		if file.IsDir() {
			continue // Skip subdirectories
		}

		if strings.HasSuffix(file.Name(), ".jpg") {
			images = append(images, file.Name())
		}
	}

	m := filepath.Join(metadatasDir, newVersion, directoryName)
	err = CreateDirectory(m)
	if err != nil {
		return err
	}

	fmt.Println("images count : ", len(images))

	for _, image := range images {

		time.Sleep(1 * time.Second)

		metadata, err2 := exiftool_v1.Start(filepath.Join(assetsDir, image))
		if err2 != nil {
			// Log the error but continue to the next image
			fmt.Printf("Exiftool error for image '%s': %v\n", image, err)
			continue // Move to the next image in the list
		}

		// Initialize destination pointer
		asset := phasset.PHAsset{}
		asset.ID = strings.Replace(image, ".jpg", "", 1)
		asset.UserID = userID
		asset.FileInfo.BaseURL = filepath.Join("com.iris.photos/users", asset.UserID, "assets")

		asset.FileInfo.FileSize = metadata.FileSize
		asset.FileInfo.FileType = metadata.FileType
		asset.Image.Orientation = metadata.Orientation

		asset.Camera.Make = metadata.Make
		asset.Camera.Model = metadata.Model

		// Test the first format
		//parsedTime1, err1 := ParseFlexibleTime(metadata.CreateDate)
		//if err1 != nil {
		//	//fmt.Printf("Error parsing '%s': %v\n", metadata.CreateDate, err1)
		//	asset.CapturedDate = time.Now()
		//} else {
		//	//fmt.Printf("Parsed '%s': %v\n", metadata.CreateDate, parsedTime1)
		//	asset.CapturedDate = parsedTime1
		//}

		//asset.CapturedDate = parsedTime1
		asset.CreatedAt = time.Now()
		asset.UpdatedAt = time.Now()
		asset.Version = newVersion

		err = WriteData(filepath.Join(metadatasDir, newVersion, directoryName, asset.ID+".json"), &asset)
		if err != nil {
			// Log the error but continue to the next image
			fmt.Printf("WriteData error for image '%s': %v\n", image, err.Error())
			continue // Move to the next image
		}
	}

	return nil
}

func upgradePHAssetsDeepseek(userID string) error {

	const directoryName = "assets"
	const maxWorkers = 1 // Limit concurrent ExifTool processes

	files, err := os.ReadDir(assetsDir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	var images []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToLower(file.Name()), ".jpg") {
			images = append(images, file.Name())
		}
	}

	m := filepath.Join(metadatasDir, newVersion, directoryName)
	err = CreateDirectory(m)
	if err != nil {
		return err
	}

	fmt.Println("images count:", len(images))

	// Create channels for work items and results
	jobs := make(chan string, len(images))
	results := make(chan processingResult, len(images))

	// Start worker goroutines
	for w := 1; w <= maxWorkers; w++ {
		go processImageWorker(w, jobs, results, userID, m)
	}

	// Send all images to the jobs channel
	for _, image := range images {
		time.Sleep(10 * time.Millisecond)
		jobs <- image
	}
	close(jobs) // lsof +D /media/mahdi/happle | wc -l

	// Collect results
	successCount := 0
	failCount := 0
	for i := 0; i < len(images); i++ {
		result := <-results
		if result.err != nil {
			fmt.Printf("Error processing %s: %v\n", result.image, result.err)
			failCount++
		} else {
			successCount++
		}
	}

	fmt.Printf("Processing complete. Successful: %d, Failed: %d\n", successCount, failCount)
	return nil
}

func upgradePHAssetsV3(userID string) error {

	const directoryName = "assets"

	files, err := os.ReadDir(assetsDir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
		return err
	}

	// Recommended: Log the error if the file can't be closed properly.
	defer func() {
		if err := et.Close(); err != nil {
			// Here, you should log the error.
			// For example, using the standard "log" package:
			log.Printf("error closing file in upgradePHAssetsV3: %v", err)
		}
	}()

	m := filepath.Join(metadatasDir, newVersion, directoryName) // create v3 directory
	err = CreateDirectory(m)
	if err != nil {
		return err
	}

	var images []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToLower(file.Name()), ".jpg") {
			images = append(images, file.Name())
		}
	}

	fmt.Println("images count:", len(images))

	fmt.Println("Start ----------------------------------------------------- ")

	for _, image := range images {
		fileInfos := et.ExtractMetadata(filepath.Join(assetsDir, image))

		fmt.Println("----------------------------------------------------- ")

		if fileInfos[0].Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfos[0].File, fileInfos[0].Err)
			continue
		}

		asset := phasset.PHAsset{
			Camera: phasset.CameraInfo{},
		}
		asset.UserID = userID

		fmt.Println(fileInfos[0].Fields["FileSize"], "          ", fileInfos[0].Fields["FileName"])

		asset.FileInfo.BaseURL = filepath.Join("com.iris.photos/users", asset.UserID, "assets")
		if FileSize, ok := fileInfos[0].Fields["FileSize"].(string); ok {
			asset.FileInfo.FileSize = FileSize
		}
		if FileType, ok := fileInfos[0].Fields["FileType"].(string); ok {
			asset.FileInfo.FileType = FileType
		}
		if MimeType, ok := fileInfos[0].Fields["MIMEType"].(string); ok {
			asset.FileInfo.MimeType = MimeType
		}

		if FileType, ok := fileInfos[0].Fields["FileType"].(string); ok {
			switch FileType {
			case "JPEG":
				asset.ID = strings.Replace(image, ".jpg", "", 1)

				if Orientation, ok := fileInfos[0].Fields["Orientation"].(string); ok {
					asset.Image.Orientation = Orientation
				}
				if ImageWidth, ok := fileInfos[0].Fields["ImageWidth"].(float64); ok {
					asset.Image.Width = int(ImageWidth)
				}
				if ImageHeight, ok := fileInfos[0].Fields["ImageHeight"].(float64); ok {
					asset.Image.Height = int(ImageHeight)
				}
				if CameraMake, ok := fileInfos[0].Fields["Make"].(string); ok {
					asset.Camera.Make = CameraMake
				}
				if Model, ok := fileInfos[0].Fields["Model"].(string); ok {
					asset.Camera.Model = Model
				}
			case "MP4":

				asset.ID = strings.Replace(image, ".mp4", "", 1)
				if MediaDuration, ok := fileInfos[0].Fields["MediaDuration"].(string); ok {
					asset.Video.MediaDuration = MediaDuration
				}
				if ImageWidth, ok := fileInfos[0].Fields["ImageWidth"].(float64); ok {
					asset.Video.Width = int(ImageWidth)
				}
				if ImageHeight, ok := fileInfos[0].Fields["ImageHeight"].(float64); ok {
					asset.Video.Height = int(ImageHeight)
				}
				if VideoFrameRate, ok := fileInfos[0].Fields["VideoFrameRate"].(float64); ok {
					asset.Video.VideoFrameRate = VideoFrameRate
				}
				if AvgBitrate, ok := fileInfos[0].Fields["AvgBitrate"].(string); ok {
					asset.Video.AvgBitrate = AvgBitrate
				}
				if Encoder, ok := fileInfos[0].Fields["Encoder"].(string); ok {
					asset.Video.Encoder = Encoder
				}
				if Rotation, ok := fileInfos[0].Fields["Rotation"].(string); ok {
					asset.Video.Encoder = Rotation
				}
				if AudioFormat, ok := fileInfos[0].Fields["AudioFormat"].(string); ok {
					asset.Video.AudioFormat = AudioFormat
				}
				if AudioChannels, ok := fileInfos[0].Fields["AudioChannels"].(int); ok {
					asset.Video.AudioChannels = AudioChannels
				}
				if AudioSampleRate, ok := fileInfos[0].Fields["AudioSampleRate"].(int); ok {
					asset.Video.AudioSampleRate = AudioSampleRate
				}
				if AudioBitsPerSample, ok := fileInfos[0].Fields["AudioBitsPerSample"].(int); ok {
					asset.Video.AudioBitsPerSample = AudioBitsPerSample
				}
			}
		}

		// Test the first format
		//parsedTime1, err1 := ParseFlexibleTime(metadata.CreateDate)
		//if err1 != nil {
		//	//fmt.Printf("Error parsing '%s': %v\n", metadata.CreateDate, err1)
		//	asset.CapturedDate = time.Now()
		//} else {
		//	//fmt.Printf("Parsed '%s': %v\n", metadata.CreateDate, parsedTime1)
		//	asset.CapturedDate = parsedTime1
		//}

		asset.CreatedAt = time.Now()
		asset.UpdatedAt = time.Now()
		asset.Version = newVersion

		err = WriteData(filepath.Join(metadatasDir, newVersion, directoryName, asset.ID+".json"), &asset)
		if err != nil {
			// Log the error but continue to the next image
			fmt.Printf("WriteData error for image '%s': %v\n", image, err.Error())
			continue // Move to the next image
		}
	}

	fmt.Println("End ------------------------------------------------------- ")

	return nil
}

type processingResult struct {
	image string
	err   error
}

func processImageWorker(id int, jobs <-chan string, results chan<- processingResult, userID string, metadataDir string) {
	for image := range jobs {
		err := processSingleImage(image, userID, metadataDir)
		results <- processingResult{image: image, err: err}
	}
}

func processSingleImage(image string, userID string, metadataDir string) error {

	metadata, err := exiftool_v1.Start(filepath.Join(assetsDir, image))
	if err != nil {
		return fmt.Errorf("exiftool_v1 error: %w", err)
	}

	// Defensive check, usually redundant if API is well-behaved
	if metadata == nil {
		return fmt.Errorf("exiftool_v1 returned nil metadata but no error")
	}

	var DateTimeOriginal time.Time
	if parsedTime, err := ParseFlexibleTime(metadata.CreateDate); err == nil {
		DateTimeOriginal = parsedTime
	} else {
		DateTimeOriginal = time.Now()
	}

	asset := phasset.PHAsset{
		ID:     strings.TrimSuffix(image, ".jpg"),
		UserID: userID,
		FileInfo: phasset.FileInfo{
			BaseURL:  filepath.Join("com.iris.photos/users", userID, "assets"),
			FileSize: metadata.FileSize,
			FileType: metadata.FileType,
		},
		Image: phasset.ImageInfo{
			Orientation: metadata.Orientation,
		},
		Camera: phasset.CameraInfo{
			Make:             metadata.Make,
			Model:            metadata.Model,
			DateTimeOriginal: DateTimeOriginal,
		},
		Video: phasset.VideoInfo{
			MediaDuration: "150",
		},

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   newVersion,
	}

	outputPath := filepath.Join(metadataDir, asset.ID+".json")
	return WriteData(outputPath, &asset)
}
