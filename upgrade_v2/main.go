package main

import (
	"log"
	"path/filepath"
)

const (
	currentVersion = "v1"
	newVersion     = "v2"
)

func main() {

	dirToDelete := filepath.Join(metadatasDir, newVersion)
	err := DeleteNestedDirectory(dirToDelete)
	if err != nil {
		log.Fatalf("Error deleting directory '%s': %v", dirToDelete, err)
		return
	}

	albumArrayV1, err := upgradeAlbums()
	if err != nil {
		log.Fatalf("Album upgrade_v2 failed: %v", err)
	}

	tripArrayV1, err := upgradeTrips()
	if err != nil {
		log.Fatalf("Trip upgrade_v2 failed: %v", err)
	}

	personArrayV1, err := upgradePersons()
	if err != nil {
		log.Fatalf("Persons upgrade_v2 failed: %v", err)
	}

	_, err = upgradePins()
	if err != nil {
		log.Fatalf("pins upgrade_v2 failed: %v", err)
	}

	_, err = upgradePHAssets("user_id_45", albumArrayV1, tripArrayV1, personArrayV1)
	if err != nil {
		log.Fatalf("PHAsset upgrade_v2 failed: %v", err)
	}

	log.Println("Upgrade completed successfully!")
}
