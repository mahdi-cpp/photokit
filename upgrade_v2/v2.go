package upgrade

import "time"

type AlbumV2 struct {
	ID                string    `json:"id"`
	Title             string    `json:"title"`
	Subtitle          string    `json:"subtitle"`
	Type              string    `json:"type"`
	Number            int       `json:"number"`
	IsCollectionValid bool      `json:"isCollectionValid"`
	IsHidden          bool      `json:"isHidden"`
	LastSeen          time.Time `json:"lastSeen"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         time.Time `json:"deletedAt"`
	Version           string    `json:"version"`
}

type PersonV2 struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Subtitle     string    `json:"subtitle"`
	Type         string    `json:"type"`
	Count        int       `json:"count"`
	IsCollection bool      `json:"isCollection"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
	Version      string    `json:"version"`
}

type PinnedV2 struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Type      string    `json:"type"`
	AlbumID   string    `json:"albumID"`
	Icon      string    `json:"icon"`
	Count     int       `json:"count"`
	Index     int       `json:"index"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Version   string    `json:"version"`
}

type TripV2 struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Subtitle     string    `json:"subtitle"`
	Type         string    `json:"type"`
	Count        int       `json:"count"`
	IsCollection bool      `json:"isCollection"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
	Version      string    `json:"version"`
}

//-----------------------------------------------------

type PHAssetV2 struct {
	ID      string `json:"id"`
	UserID  string `json:"userID"`
	BaseURL string `json:"baseURL"`
	//Url                 string    `json:"url"`
	FileName            string    `json:"fileName"`
	FilePath            string    `json:"filePath"`
	Format              string    `json:"format"`
	MediaType           MediaType `json:"mediaType"`
	Orientation         int       `json:"orientation"`
	PixelWidth          int       `json:"pixelWidth"`
	PixelHeight         int       `json:"pixelHeight"`
	Place               PlaceV2   `json:"place"`
	CameraMake          string    `json:"cameraMake"`
	CameraModel         string    `json:"cameraModel"`
	IsCamera            bool      `json:"isCamera"`
	IsFavorite          bool      `json:"isFavorite"`
	IsScreenshot        bool      `json:"isScreenshot"`
	IsHidden            bool      `json:"isHidden"`
	Albums              []string  `json:"albums"`
	Trips               []string  `json:"trips"`
	Persons             []string  `json:"persons"`
	Duration            float64   `json:"duration"`
	CanDelete           bool      `json:"canDelete"`
	CanEditContent      bool      `json:"canEditContent"`
	CanAddToSharedAlbum bool      `json:"canAddToSharedAlbum"`
	IsUserLibraryAsset  bool      `json:"IsUserLibraryAsset"`
	CapturedDate        time.Time `json:"capturedDate"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	DeletedAt           time.Time `json:"deletedAt"`
	Version             string    `json:"version"`
}

type PlaceV2 struct {
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
