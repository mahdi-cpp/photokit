package storage

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-pkg/asset_metadata_manager"
	"github.com/mahdi-cpp/api-go-pkg/collection"
	"github.com/mahdi-cpp/api-go-pkg/image_loader"
	"github.com/mahdi-cpp/api-go-pkg/network"
	"github.com/mahdi-cpp/api-go-pkg/shared_model"
	"github.com/mahdi-cpp/api-go-pkg/thumbnail"
	"github.com/mahdi-cpp/photocloud_v2/config"
	"github.com/mahdi-cpp/photocloud_v2/internal/domain/model"
	"log"
	"sync"
	"time"
)

type MainStorageManager struct {
	mu           sync.RWMutex
	users        map[int]*shared_model.User
	userStorages map[int]*UserStorage // Maps user IDs to their UserStorage
	iconLoader   *image_loader.ImageLoader
	ctx          context.Context
}

func NewMainStorageManager() (*MainStorageManager, error) {

	// Handler the manager
	manager := &MainStorageManager{
		userStorages: make(map[int]*UserStorage),
		users:        make(map[int]*shared_model.User),
		ctx:          context.Background(),
	}

	// Alternative using PHCollection directly
	type UserPHCollection struct {
		Collections []*shared_model.PHCollection[shared_model.User] `json:"collections"`
	}

	userControl := network.NewNetworkControl[UserPHCollection]("http://localhost:8080/api/v1/user/")

	// Make request (nil body if not needed)
	response, err := userControl.Read("list", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Process the response
	for _, userCollection := range response.Collections {
		user := userCollection.Item
		fmt.Printf("User ID: %d\n", user.ID)
		fmt.Printf("Username: %s\n", user.Username)
		fmt.Printf("Name: %s %s\n", user.FirstName, user.LastName)
		fmt.Println("-----")
		manager.users[user.ID] = &user
	}

	manager.iconLoader = image_loader.NewImageLoader(1000, config.GetPath("/data/icons"), 0)
	manager.loadAllIcons()

	return manager, nil
}

func (us *MainStorageManager) GetUserStorage(c *gin.Context, userID int) (*UserStorage, error) {

	us.mu.Lock()
	defer us.mu.Unlock()

	var err error

	if userID <= 0 {
		return nil, fmt.Errorf("user id is Invalid")
	}

	var user = us.users[userID]

	// Check if userStorage already exists for this user
	if storage, exists := us.userStorages[userID]; exists {
		return storage, nil
	}

	fmt.Println("GetUserStorage.... 1")
	// Handler context for background workers
	ctx, cancel := context.WithCancel(context.Background())

	// Ensure user directories exist
	//userDirs := []string{userAssetDir, userMetadataDir, userThumbnailsDir}
	//for _, dir := range userDirs {
	//	if err := os.MkdirAll(dir, 0755); err != nil {
	//		return nil, fmt.Errorf("failed to create user directory %s: %w", dir, err)
	//	}
	//}

	// Handler new userStorage for this user
	userStorage := &UserStorage{
		user:              *user,
		metadata:          asset_metadata_manager.NewMetadataManager(config.GetUserPath(user.PhoneNumber, "metadata")),
		thumbnail:         thumbnail.NewThumbnailManager(config.GetUserPath(user.PhoneNumber, "thumbnails")),
		maintenanceCtx:    ctx,
		cancelMaintenance: cancel,
	}

	userStorage.originalImageLoader = image_loader.NewImageLoader(50, config.GetUserPath(user.PhoneNumber, "assets"), 5*time.Minute)
	userStorage.tinyImageLoader = image_loader.NewImageLoader(30000, config.GetUserPath(user.PhoneNumber, "thumbnails"), 60*time.Minute)

	userStorage.assets, err = userStorage.metadata.LoadUserAllMetadata()
	if err != nil {
		return nil, fmt.Errorf("failed to load metadata for user %s: %w", userID, err)
	}

	userStorage.AlbumManager, err = collection.NewCollectionManager[*model.Album](config.GetUserPath(user.PhoneNumber, "data/albums.json"), false)
	if err != nil {
		panic(err)
	}

	userStorage.SharedAlbumManager, err = collection.NewCollectionManager[*model.SharedAlbum](config.GetUserPath(user.PhoneNumber, "data/shared_albums.json"), false)
	if err != nil {
		panic(err)
	}

	userStorage.TripManager, err = collection.NewCollectionManager[*model.Trip](config.GetUserPath(user.PhoneNumber, "data/trips.json"), false)
	if err != nil {
		panic(err)
	}

	userStorage.PersonManager, err = collection.NewCollectionManager[*model.Person](config.GetUserPath(user.PhoneNumber, "data/persons.json"), false)
	if err != nil {
		panic(err)
	}

	userStorage.PinnedManager, err = collection.NewCollectionManager[*model.Pinned](config.GetUserPath(user.PhoneNumber, "data/pinned.json"), false)
	if err != nil {
		panic(err)
	}

	userStorage.VillageManager, err = collection.NewCollectionManager[*model.Village](config.GetPath("/data/villages.json"), false)
	if err != nil {
		panic(err)
	}

	userStorage.prepareAlbums()
	userStorage.prepareTrips()
	userStorage.preparePersons()
	userStorage.prepareCameras()
	userStorage.preparePinned()

	// Store the new userStorage
	us.userStorages[userID] = userStorage

	return userStorage, nil
}

func (us *MainStorageManager) loadAllIcons() {
	us.iconLoader.GetLocalBasePath()

	// Scan metadata directory
	//files, err := os.ReadDir(us.iconLoader.GetLocalBasePath())
	//if err != nil {
	//	fmt.Println("failed to read metadata directory: %w", err)
	//}

	//var images []string
	//for _, file := range files {
	//	if strings.HasSuffix(file.Name(), ".png") {
	//		images = append(images, "/media/mahdi/Cloud/apps/Photos/parsa_nasiri/assets/"+file.Name())
	//	}
	//}
}

func (us *MainStorageManager) GetAssetManager(c *gin.Context, userID int) (*collection.Manager[*model.Person], error) {
	userStorage, err := us.GetUserStorage(c, userID)
	if err != nil {
		return nil, err
	}

	return userStorage.PersonManager, nil
}

func (us *MainStorageManager) periodicMaintenance() {

	saveTicker := time.NewTicker(10 * time.Second)
	statsTicker := time.NewTicker(30 * time.Minute)
	rebuildTicker := time.NewTicker(24 * time.Hour)
	cleanupTicker := time.NewTicker(1 * time.Hour)

	for {
		select {
		case <-saveTicker.C:
			fmt.Println("saveTicker")
		case <-rebuildTicker.C:
			fmt.Println("rebuildTicker")
		case <-statsTicker.C:
			fmt.Println("statsTicker")
		case <-cleanupTicker.C:
			fmt.Println("cleanupTicker")
		}
	}
}

func (us *MainStorageManager) RepositoryGetOriginalImage(userID int, filename string) ([]byte, error) {
	return us.userStorages[userID].originalImageLoader.LoadImage(us.ctx, filename)
}

func (us *MainStorageManager) RepositoryGetTinyImage(userID int, filename string) ([]byte, error) {
	return us.userStorages[userID].tinyImageLoader.LoadImage(us.ctx, filename)
}

func (us *MainStorageManager) RepositoryGetIcon(filename string) ([]byte, error) {
	return us.iconLoader.LoadImage(us.ctx, filename)
}

func (us *MainStorageManager) RemoveStorageForUser(userID int) {
	us.mu.Lock()
	defer us.mu.Unlock()

	if storage, exists := us.userStorages[userID]; exists {
		// Cancel any background operations
		storage.cancelMaintenance()
		// Remove from map
		delete(us.userStorages, userID)
	}
}
