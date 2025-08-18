package phasset

import (
	"time"
)

func (a *PHAsset) SetID(id string)          { a.ID = id }
func (a *PHAsset) SetCreatedAt(t time.Time) { a.CreatedAt = t }
func (a *PHAsset) SetUpdatedAt(t time.Time) { a.UpdatedAt = t }
func (a *PHAsset) GetID() string            { return a.ID }
func (a *PHAsset) GetCreatedAt() time.Time  { return a.CreatedAt }
func (a *PHAsset) GetUpdatedAt() time.Time  { return a.UpdatedAt }

type MediaType string

type PHAsset struct {
	ID                    string    `json:"id"`
	UserID                string    `json:"userID"`
	Url                   string    `json:"url"`
	FileName              string    `json:"fileName"`
	FilePath              string    `json:"filePath"`
	Format                string    `json:"format"`
	MediaType             MediaType `json:"mediaType"`
	Orientation           int       `json:"orientation"`
	PixelWidth            int       `json:"pixelWidth"`
	PixelHeight           int       `json:"pixelHeight"`
	Place                 Place     `json:"place"`
	CameraMake            string    `json:"cameraMake"`
	CameraModel           string    `json:"cameraModel"`
	IsCamera              bool      `json:"isCamera"`
	IsFavorite            bool      `json:"isFavorite"`
	IsScreenshot          bool      `json:"isScreenshot"`
	IsHidden              bool      `json:"isHidden"`
	Albums                []string  `json:"albums"`
	Trips                 []string  `json:"trips"`
	Persons               []string  `json:"persons"`
	Duration              float64   `json:"duration"`
	CanDelete             bool      `json:"canDelete"`
	CanEditContent        bool      `json:"canEditContent"`
	CanAddToSharedPHAsset bool      `json:"canAddToSharedPHAsset"`
	IsUserLibraryAsset    bool      `json:"IsUserLibraryAsset"`
	CapturedDate          time.Time `json:"capturedDate"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	DeletedAt             time.Time `json:"deletedAt"`
	Version               string    `json:"version"`
}

type Place struct {
	Latitude   float64 `json:"location"`
	Longitude  float64 `json:"longitude"`
	Country    string  `json:"country"`
	Province   string  `json:"province"`
	County     string  `json:"county"`
	City       string  `json:"city"`
	Village    string  `json:"village"`
	Malard     string  `json:"malard"`
	Electronic int     `json:"electronic"`
}

type UpdateOptions struct {
	AssetIds []string `json:"assetIds,omitempty"` // Asset Ids

	FileName  string    `json:"fileName,omitempty"`
	MediaType MediaType `json:"mediaType,omitempty"`

	CameraMake  *string `json:"cameraMake,omitempty"`
	CameraModel *string `json:"cameraModel,omitempty"`

	IsCamera        *bool
	IsFavorite      *bool
	IsScreenshot    *bool
	IsHidden        *bool
	NotInOnePHAsset *bool

	Albums       *[]string `json:"albums,omitempty"`       // Full album replacement
	AddAlbums    []string  `json:"addAlbums,omitempty"`    // PHAssets to add
	RemoveAlbums []string  `json:"removeAlbums,omitempty"` // PHAssets to remove

	Trips       *[]string `json:"trips,omitempty"`       // Full trip replacement
	AddTrips    []string  `json:"addTrips,omitempty"`    // Trips to add
	RemoveTrips []string  `json:"removeTrips,omitempty"` // Trips to remove

	Persons       *[]string `json:"persons,omitempty"`       // Full Person replacement
	AddPersons    []string  `json:"addPersons,omitempty"`    // Persons to add
	RemovePersons []string  `json:"removePersons,omitempty"` // Persons to remove
}

type SearchOptions struct {
	ID     string
	UserID string

	TextQuery string

	FileName string `json:"fileName"`
	FilePath string `json:"filePath"`
	Format   string `json:"format"`

	MediaType   MediaType
	PixelWidth  int
	PixelHeight int

	CameraMake  string
	CameraModel string

	IsCamera        *bool
	IsFavorite      *bool
	IsScreenshot    *bool
	IsHidden        *bool
	IsLandscape     *bool
	NotInOnePHAsset *bool

	HideScreenshot *bool `json:"hideScreenshot"`

	Albums  []string
	Trips   []string
	Persons []string

	NearPoint    []float64 `json:"nearPoint"`    // [latitude, longitude]
	WithinRadius float64   `json:"withinRadius"` // in kilometers
	BoundingBox  []float64 `json:"boundingBox"`  // [minLat, minLon, maxLat, maxLon]

	// Date filters
	CreatedAfter  *time.Time `json:"createdAfter,omitempty"`
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`
	ActiveAfter   *time.Time `json:"activeAfter,omitempty"`

	// Pagination
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`

	// Sorting
	SortBy    string `json:"sortBy,omitempty"`    // "title", "created", "members", "lastActivity"
	SortOrder string `json:"sortOrder,omitempty"` // "asc" or "desc"
}

//type Delete struct {
//	AssetID string `json:"assetID"`
//}

// https://chat.deepseek.com/a/chat/s/9b010f32-b23d-4f9b-ae0c-31a9b2c9408c

//type PHFetchResult[T any] struct {
//	Items  []T `json:"items"`
//	Total  int `json:"total"`
//	Limit  int `json:"limit"`
//	Offset int `json:"offset"`
//}
