package upgrade

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func upgradePHAssets(userID string, albumArrayV1 []*AlbumV1, tripArrayV1 []*TripV1, personArrayV1 []*PersonV1) ([]*PHAssetV2, error) {

	const directoryName = "assets"

	dirPath := filepath.Join(metadataDir, currentVersion, directoryName)

	itemArrayV1, err := readJsonFilesFromDirectory(dirPath)
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}

	// Convert to V2 with pointers
	itemArrayV2 := make([]*PHAssetV2, len(itemArrayV1))
	var index = 1

	for i := range itemArrayV1 {
		// Initialize destination pointer
		itemArrayV2[i] = &PHAssetV2{}

		// Copy matching fields
		copyMatchingFields(&itemArrayV2[i], &itemArrayV1[i])

		u7, err2 := uuid.NewV7()
		if err2 != nil {
			return nil, fmt.Errorf("error generating UUIDv7: %w", err)
		}

		itemArrayV2[i].ID = u7.String()
		itemArrayV2[i].UserID = userID

		itemArrayV2[i].FileName = itemArrayV1[i].Filename
		itemArrayV2[i].FilePath = itemArrayV1[i].Filepath
		itemArrayV2[i].BaseURL = filepath.Join("com.iris.photos/users", itemArrayV2[i].UserID, "assets")
		//itemArrayV2[i].Url = itemArrayV2[i].ID + ".jpg"

		if itemArrayV1[i].Albums != nil {
			for _, item := range albumArrayV1 {
				if contains(itemArrayV1[i].Albums, item.ID) {
					itemArrayV2[i].Albums = append(itemArrayV2[i].Albums, item.NewID)
				}
			}
		}
		if itemArrayV1[i].Trips != nil {
			for _, item := range tripArrayV1 {
				if contains(itemArrayV1[i].Trips, item.ID) {
					itemArrayV2[i].Trips = append(itemArrayV2[i].Trips, item.NewID)
				}
			}
		}
		if itemArrayV1[i].Persons != nil {
			for _, item := range personArrayV1 {
				if contains(itemArrayV1[i].Persons, item.ID) {
					itemArrayV2[i].Persons = append(itemArrayV2[i].Persons, item.NewID)
				}
			}
		}

		itemArrayV2[i].CreatedAt = itemArrayV1[i].CreationDate
		itemArrayV2[i].UpdatedAt = time.Now()

		/*
			renameErr := RenameFile(filepath.Join(assetsDir, strconv.Itoa(itemArrayV1[i].ID)+".jpg"), filepath.Join(assetsDir, itemArrayV2[i].ID+".jpg"))
			if renameErr != nil {
				return nil, renameErr
			}

			renameErr = RenameFile(filepath.Join(thumbnailsDir, strconv.Itoa(itemArrayV1[i].ID)+"_270.jpg"), filepath.Join(thumbnailsDir, itemArrayV2[i].ID+"_270.jpg"))
			if renameErr != nil {
				return nil, renameErr
			}

		*/

		index++
	}

	m := filepath.Join(metadataDir, newVersion, directoryName)
	err = CreateDirectory(m)
	if err != nil {
		return nil, err
	}

	for _, item := range itemArrayV2 {
		err = WriteData(filepath.Join(m, item.ID+".json"), &item)
		if err != nil {
			return nil, err
		}
	}

	log.Printf("Upgraded %d PHAssets", len(itemArrayV2))
	return itemArrayV2, nil
}
