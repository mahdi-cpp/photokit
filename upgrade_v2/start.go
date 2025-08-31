package upgrade

import (
	"fmt"
	"log"
	"path/filepath"
)

func StartRename(accountManager *account.ClientManager) {

	for _, user := range accountManager.Users {

		currentDir := filepath.Join(usersDir, user.PhoneNumber)

		// Check if the directory exists
		exists, err := IsDirectoryExist(currentDir)
		if err != nil {
			//fmt.Printf("Error checking directory '%s': %v\n", usersDir, err)
			continue
		} else if exists {
			//fmt.Printf("✅ The directory '%s' exists.\n", usersDir)
		} else {
			//fmt.Printf("❌ The directory '%s' does not exist.\n", usersDir)
			continue
		}

		err = RenameDirectory(currentDir, filepath.Join(usersDir, user.ID))
		if err != nil {
			fmt.Println("Error renaming directory:", err)
			return
		}
	}

	fmt.Printf("Renaming user directories operation are completed.\n\n")
}

func Start(accountManager *account.ClientManager) {

	dirToDelete := filepath.Join(metadataDir, newVersion)
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

	_, err = upgradePHAssets("018f3a8b-1b32-729a-f7e5-5467c1b2d3e4", albumArrayV1, tripArrayV1, personArrayV1)
	if err != nil {
		log.Fatalf("PHAsset upgrade_v2 failed: %v", err)
	}

	log.Println("Upgrade completed successfully!")
}
