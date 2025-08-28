package upgrade

import "time"

type AlbumV1 struct {
	ID               int `json:"id"`
	NewID            string
	Title            string    `json:"title"`
	Subtitle         string    `json:"subtitle"`
	AlbumType        string    `json:"albumType"`
	Count            int       `json:"count"`
	IsCollection     bool      `json:"isCollection"`
	IsHidden         bool      `json:"isHidden"`
	CreationDate     time.Time `json:"creationDate"`
	ModificationDate time.Time `json:"modificationDate"`
}

type PersonV1 struct {
	ID               int `json:"id"`
	NewID            string
	Title            string    `json:"title,omitempty"`
	Subtitle         string    `json:"subtitle,omitempty"`
	Count            int       `json:"count"`
	IsCollection     bool      `json:"isCollection"`
	CreationDate     time.Time `json:"creationDate"`
	ModificationDate time.Time `json:"modificationDate"`
}

type PinnedV1 struct {
	ID               int `json:"id"`
	NewID            string
	Title            string    `json:"title"`
	Subtitle         string    `json:"subtitle"`
	Type             string    `json:"type"`
	AlbumID          int       `json:"albumID"`
	Icon             string    `json:"icon"`
	Count            int       `json:"count"`
	Index            int       `json:"index"`
	CreationDate     time.Time `json:"creationDate"`
	ModificationDate time.Time `json:"modificationDate"`
}

type TripV1 struct {
	ID               int `json:"id"`
	NewID            string
	Title            string    `json:"title,omitempty"`
	Subtitle         string    `json:"subtitle,omitempty"`
	TripType         string    `json:"tripType,omitempty"`
	Count            int       `json:"count"`
	IsCollection     bool      `json:"isCollection"`
	CreationDate     time.Time `json:"creationDate"`
	ModificationDate time.Time `json:"modificationDate"`
}

//---------------------------------------------------------

type MediaType string

type PHAssetV1 struct {
	ID                  int       `json:"id"`
	UserID              int       `json:"userID"`
	Url                 string    `json:"url"`
	Filename            string    `json:"filename"`
	Filepath            string    `json:"filepath"`
	Format              string    `json:"format"`
	MediaType           MediaType `json:"mediaType"`
	Orientation         int       `json:"orientation"`
	PixelWidth          int       `json:"pixelWidth"`
	PixelHeight         int       `json:"pixelHeight"`
	Place               PlaceV1   `json:"place"`
	CameraMake          string    `json:"cameraMake"`
	CameraModel         string    `json:"cameraModel"`
	IsCamera            bool      `json:"isCamera"`
	IsFavorite          bool      `json:"isFavorite"`
	IsScreenshot        bool      `json:"isScreenshot"`
	IsHidden            bool      `json:"isHidden"`
	Albums              []int     `json:"albums"`
	Trips               []int     `json:"trips"`
	Persons             []int     `json:"persons"`
	Duration            float64   `json:"duration"`
	CanDelete           bool      `json:"canDelete"`
	CanEditContent      bool      `json:"canEditContent"`
	CanAddToSharedAlbum bool      `json:"canAddToSharedAlbum"`
	IsUserLibraryAsset  bool      `json:"IsUserLibraryAsset"`
	CapturedDate        time.Time `json:"capturedDate"`
	CreationDate        time.Time `json:"creationDate"`
	ModificationDate    time.Time `json:"modificationDate"`
}

const (
	UnknownType   MediaType = "unknown"
	ImageTypeJPEG MediaType = "image/jpeg"
	ImageTypePNG  MediaType = "image/png"
	ImageTypeGIF  MediaType = "image/gif"
	VideoTypeMP4  MediaType = "video/mp4"
	VideoTypeMOV  MediaType = "video/quicktime"
)

type PlaceV1 struct {
	Latitude  float64 `json:"location"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
	Province  string  `json:"province"`
	County    string  `json:"county"`
	City      string  `json:"city"`
	Village   string  `json:"village"`
}
