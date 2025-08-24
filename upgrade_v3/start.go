package upgrade_v3

import (
	"fmt"
	"github.com/mahdi-cpp/go-account-service/account"
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

	dirToDelete := filepath.Join(metadatasDir, newVersion)
	err := DeleteNestedDirectory(dirToDelete)
	if err != nil {
		log.Fatalf("Error deleting directory '%s': %v", dirToDelete, err)
		return
	}

	err = upgradePHAssetsV3("018fe65d-8e4a-74b0-8001-c8a7c29367e1")
	if err != nil {
		log.Fatalf("PHAsset to version %s failed: %v", newVersion, err)
	}

	log.Println("Upgrade completed successfully!")
}
