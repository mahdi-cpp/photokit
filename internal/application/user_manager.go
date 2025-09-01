package application

import (
	"context"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"sync"

	"github.com/mahdi-cpp/api-go-pkg/collection_manager_uuid7"
	"github.com/mahdi-cpp/photokit/internal/config"

	"github.com/mahdi-cpp/photokit/internal/collection_manager_v3"
	asset "github.com/mahdi-cpp/photokit/internal/collections"
	"github.com/mahdi-cpp/photokit/internal/collections/album"
	"github.com/mahdi-cpp/photokit/internal/collections/camera"
	"github.com/mahdi-cpp/photokit/internal/collections/person"
	"github.com/mahdi-cpp/photokit/internal/collections/phasset"
	"github.com/mahdi-cpp/photokit/internal/collections/pinned"
	"github.com/mahdi-cpp/photokit/internal/collections/shared_album"
	"github.com/mahdi-cpp/photokit/internal/collections/trip"
	"github.com/mahdi-cpp/photokit/internal/collections/village"
	"github.com/mahdi-cpp/photokit/internal/utils"
)

type PhotoAssetCollection[T collection_manager_uuid7.CollectionItem] struct {
	Collection     *collection_manager_uuid7.Manager[T]
	PhotoAssetList map[string][]*phasset.PHAsset
}

func NewPhotoAssetCollection[T collection_manager_uuid7.CollectionItem](path string) *PhotoAssetCollection[T] {

	collection, err := collection_manager_uuid7.NewCollectionManager[T](path, false)
	if err != nil {
		panic(err)
	}

	photoAssetCollection := &PhotoAssetCollection[T]{
		Collection:     collection,
		PhotoAssetList: make(map[string][]*phasset.PHAsset),
	}

	return photoAssetCollection
}

type Collection struct {
	Assets       *collection_manager_v3.Manager[*phasset.PHAsset]
	Album        *PhotoAssetCollection[*album.Album]
	Trips        *PhotoAssetCollection[*trip.Trip]
	Persons      *PhotoAssetCollection[*person.Person]
	Pinned       *PhotoAssetCollection[*pinned.Pinned]
	SharedAlbums *PhotoAssetCollection[*shared_album.SharedAlbum]
	Villages     *PhotoAssetCollection[*village.Village]
}

type UserManager struct {
	mu                sync.RWMutex
	userID            string
	collection        *Collection
	cameras           map[string]*asset.PHCollection[camera.Camera]
	maintenanceCtx    context.Context
	cancelMaintenance context.CancelFunc
	statsMu           sync.Mutex
}

func NewUserManager(userID string) (*UserManager, error) {

	// Handler context for background workers
	ctx, cancel := context.WithCancel(context.Background())

	userManager := &UserManager{
		userID:            userID,
		collection:        &Collection{},
		maintenanceCtx:    ctx,
		cancelMaintenance: cancel,
	}

	// Ensure user directories exist
	userDirectory := filepath.Join(config.GetUserPath(userID))
	userMetadata := filepath.Join(userDirectory, "metadata")
	userDirs := []string{userDirectory, userMetadata}
	for _, dir := range userDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create user directory %s: %w", dir, err)
		}
	}

	var err error
	userManager.collection.Assets, err = collection_manager_v3.NewCollectionManager[*phasset.PHAsset](config.GetUserMetadataPath(userID, "assets"), true)
	if err != nil {
		panic(err)
	}

	all, err := userManager.collection.Assets.GetAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(len(all))

	userManager.collection.Album = NewPhotoAssetCollection[*album.Album](config.GetUserMetadataPath(userID, "albums"))
	userManager.collection.SharedAlbums = NewPhotoAssetCollection[*shared_album.SharedAlbum](config.GetUserMetadataPath(userID, "shared_albums"))
	userManager.collection.Trips = NewPhotoAssetCollection[*trip.Trip](config.GetUserMetadataPath(userID, "trips"))
	userManager.collection.Persons = NewPhotoAssetCollection[*person.Person](config.GetUserMetadataPath(userID, "persons"))
	userManager.collection.Pinned = NewPhotoAssetCollection[*pinned.Pinned](config.GetUserMetadataPath(userID, "pins"))
	userManager.collection.Villages = NewPhotoAssetCollection[*village.Village](config.GetUserMetadataPath(userID, "villages"))

	userManager.prepareAlbums()
	userManager.prepareTrips()
	userManager.preparePersons()
	userManager.prepareCameras()
	userManager.preparePinned()

	return userManager, nil
}

func (m *UserManager) UpdateCollections() {
	m.prepareAlbums()
	m.prepareCameras()
	m.prepareTrips()
	m.preparePersons()
	m.preparePinned()
}

func (m *UserManager) UpdateAssets(updateOptions phasset.UpdateOptions) (string, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

	for _, id := range updateOptions.AssetIds {

		item, err := m.collection.Assets.Get(id)
		if err != nil {
			continue
		}

		phasset.Update(item, updateOptions)

		_, err = m.collection.Assets.Update(item)
		if err != nil {
			return "", err
		}
	}

	// Merging strings with the integer ID
	merged := fmt.Sprintf(" %s, %d:", "updateOptions person_test count: ", len(updateOptions.AssetIds))

	return merged, nil
}

func (m *UserManager) prepareAlbums() {

	items, err := m.collection.Album.Collection.GetAll()
	if err != nil {
	}

	for _, item := range items {

		with := &phasset.SearchOptions{
			UserID:    m.userID,
			Albums:    []string{item.ID},
			SortBy:    "createdAt",
			SortOrder: "start",
			Limit:     6,
		}

		filterAssets, err := m.FetchAssets(with)
		if err != nil {
			fmt.Printf("Error getting all person_test: %v\n", err)
			return
		}
		item.Count = len(filterAssets)
		m.collection.Album.PhotoAssetList[item.ID] = filterAssets
	}
}

func (m *UserManager) prepareTrips() {

	items, err := m.collection.Trips.Collection.GetAll()
	if err != nil {
	}

	for _, item := range items {
		with := &phasset.SearchOptions{
			UserID:    "",
			Trips:     []string{item.ID},
			SortBy:    "createdAt",
			SortOrder: "start",
			Limit:     2,
		}

		filterAssets, err := m.FetchAssets(with)
		if err != nil {
			fmt.Printf("Error getting all person_test: %v\n", err)
			return
		}
		item.Count = len(filterAssets)
		m.collection.Trips.PhotoAssetList[item.ID] = filterAssets
	}
}

func (m *UserManager) preparePersons() {

	items, err := m.collection.Persons.Collection.GetAll()
	if err != nil {
	}

	for _, item := range items {
		with := &phasset.SearchOptions{
			UserID:    m.userID,
			Persons:   []string{item.ID},
			SortBy:    "createdAt",
			SortOrder: "start",
			Limit:     1,
		}
		filterAssets, err := m.FetchAssets(with)
		if err != nil {
			fmt.Printf("Error getting all person_test: %v\n", err)
			return
		}
		item.Count = len(filterAssets)
		m.collection.Persons.PhotoAssetList[item.ID] = filterAssets
	}
}

func (m *UserManager) prepareCameras() {

	//items, err := m.CameraManager.GetAll()
	//if err != nil {
	//}

	if m.cameras == nil {
		m.cameras = map[string]*asset.PHCollection[camera.Camera]{}
	}

	assets, err := m.collection.Assets.GetAll()
	if err != nil {
		fmt.Printf("Error getting all person_test: %v\n", err)
		return
	}

	for _, phAsset := range assets {
		if phAsset.Camera.Model == "" {
			continue
		}

		cameraManager, exists := m.cameras[phAsset.Camera.Model]
		if exists {
			cameraManager.Item.Count = cameraManager.Item.Count + 1
			m.cameras[phAsset.Camera.Model] = cameraManager
		} else {
			collection := &asset.PHCollection[camera.Camera]{
				Item: camera.Camera{
					ID:          "1",
					CameraMake:  phAsset.Camera.Make,
					CameraModel: phAsset.Camera.Model,
					Count:       1},
			}
			//fmt.Println(collection)
			m.cameras[phAsset.Camera.Model] = collection
		}
	}

	for _, collection := range m.cameras {

		with := &phasset.SearchOptions{
			UserID:      m.userID,
			CameraMake:  collection.Item.CameraMake,
			CameraModel: collection.Item.CameraModel,
			SortBy:      "createdAt",
			SortOrder:   "start",
			Limit:       6,
		}

		filterAssets, err := m.FetchAssets(with)
		if err != nil {
			fmt.Printf("Error getting all person_test: %v\n", err)
			return
		}
		collection.Assets = filterAssets
	}
}

func (m *UserManager) preparePinned() {

	items, err := m.collection.Pinned.Collection.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, item := range items {

		var with *phasset.SearchOptions

		switch item.Type {
		case "camera":
			with = &phasset.SearchOptions{
				IsCamera:  utils.GetBoolPointer(true),
				SortBy:    "createdAt",
				SortOrder: "start",
				Limit:     1,
			}
			break
		case "screenshot":
			with = &phasset.SearchOptions{
				IsScreenshot: utils.GetBoolPointer(true),
				SortBy:       "createdAt",
				SortOrder:    "start",
				Limit:        1,
			}
			break
		case "favorite":
			with = &phasset.SearchOptions{
				IsFavorite: utils.GetBoolPointer(true),
				SortBy:     "createdAt",
				SortOrder:  "start",
				Limit:      1,
			}
			break
		case "video":
			with = &phasset.SearchOptions{
				MimeType:  "video/mp4",
				SortBy:    "createdAt",
				SortOrder: "start",
				Limit:     1,
			}
			break
		case "map":
			var assets []*phasset.PHAsset
			asset1 := phasset.PHAsset{
				ID: "12",
				FileInfo: phasset.FileInfo{
					BaseURL:  "map",
					FileType: "map",
				},
			}
			assets = append(assets, &asset1)
			m.collection.Pinned.PhotoAssetList[item.ID] = assets
			break
		case "album":
			selectedAlbum, err := m.collection.Album.Collection.Get(item.AlbumID)
			if err != nil {
				continue
			}
			item.Title = selectedAlbum.Title
			with = &phasset.SearchOptions{
				Albums:    []string{selectedAlbum.ID},
				SortBy:    "createdAt",
				SortOrder: "start",
				Limit:     1,
			}
			break
		}

		if with == nil || item.Type == "map" {
			continue
		}

		filterAssets, err := m.FetchAssets(with)
		if err != nil {
			fmt.Printf("Error getting all person_test: %v\n", err)
			return
		}
		item.Count = len(filterAssets)
		m.collection.Pinned.PhotoAssetList[item.ID] = filterAssets
	}
}

func (m *UserManager) GetAllCameras() map[string]*asset.PHCollection[camera.Camera] {
	return m.cameras
}

func (m *UserManager) GetCollections() *Collection {
	return m.collection
}

func (m *UserManager) GetAllAssets() ([]*phasset.PHAsset, error) {
	return m.collection.Assets.GetAll()
}

func (m *UserManager) FetchAssets(with *phasset.SearchOptions) ([]*phasset.PHAsset, error) {
	allAssets, err := m.collection.Assets.GetAll()
	if err != nil {
		fmt.Printf("Error getting all person_test: %v\n", err)
		return nil, err
	}

	return phasset.Search(allAssets, with), nil
}

func (m *UserManager) DeleteAsset(id string) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	// Get person_test
	//person_test, err := m.GetAsset(id)
	//if err != nil {
	//	return err
	//}

	// Delete person_test file
	//assetPath := filepath.Join(m.config.AssetsDir, person_test.Filename)
	//if err := os.Remove(assetPath); err != nil {
	//	return fmt.Errorf("failed to delete person_test file: %w", err)
	//}

	// Delete metadata
	err := m.collection.Assets.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete metadata: %w", err)
	}
	//if err := m.metadata.DeleteMetadata(id); err != nil {
	//	return fmt.Errorf("failed to delete metadata: %w", err)
	//}

	// Delete thumbnail (if exists)
	//m.thumbnail.DeleteThumbnails(id)

	// Remove from indexes
	//m.removeFromIndexes(id)

	// Remove from memory
	//m.memory.Remove(id)

	// UpdateOptions stats
	//m.statsMu.Lock()
	//m.stats.TotalAssets--
	//m.statsMu.Unlock()

	return nil
}
