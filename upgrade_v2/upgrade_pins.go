package upgrade

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func upgradePins() ([]*PinnedV1, error) {

	const directoryName = "pins"

	itemFile := filepath.Join(metadataDir, currentVersion, "pinned.json")

	// Read V1 data
	data, err := os.ReadFile(itemFile)
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}

	var itemArrayV1 []*PinnedV1
	if err = json.Unmarshal(data, &itemArrayV1); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	for _, item := range itemArrayV1 {

		u7, err2 := uuid.NewV7()
		if err2 != nil {
			return nil, fmt.Errorf("error generating UUIDv7: %w", err)
		}
		item.NewID = u7.String()
	}

	// Convert to V2 with pointers
	itemArrayV2 := make([]*PinnedV2, len(itemArrayV1))
	var index = 1
	for i := range itemArrayV1 {
		// Initialize destination pointer
		itemArrayV2[i] = &PinnedV2{}

		// Copy matching fields
		copyMatchingFields(&itemArrayV2[i], &itemArrayV1[i])

		// Set new UUID-based ID
		itemArrayV2[i].ID = itemArrayV1[i].NewID

		itemArrayV2[i].Type = itemArrayV1[i].Type

		itemArrayV2[i].CreatedAt = itemArrayV1[i].CreationDate
		itemArrayV2[i].UpdatedAt = time.Now()
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

	log.Printf("Upgraded %d pins", len(itemArrayV2))
	return itemArrayV1, nil
}
