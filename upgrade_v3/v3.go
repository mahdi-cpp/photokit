package upgrade_v3

import (
	"time"
)

type AlbumV3 struct {
	ID                string    `json:"id"`
	Title             string    `json:"title"`
	Subtitle          string    `json:"subtitle"`
	Type              string    `json:"type"`
	Number            int       `json:"number"`
	IsCollectionValid bool      `json:"isCollectionValid"`
	IsHidden          bool      `json:"isHidden"`
	LastSeen          time.Time `json:"lastSeen"`
	Assets            []string  `json:"assets"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         time.Time `json:"deletedAt"`
	Version           string    `json:"version"`
}

type PersonV3 struct {
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

type PinnedV3 struct {
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

type TripV3 struct {
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
//
//type PHAssetV3 struct {
//	ID     string `json:"id"`
//	UserID string `json:"userID"`
//
//	File     FileInfo   `json:"file,omitempty"`
//	Image    ImageInfo  `json:"image,omitempty"`
//	Video    VideoInfo  `json:"video,omitempty"`
//	Camera   CameraInfo `json:"camera"`
//	Location Location   `json:"location,omitempty"`
//
//	IsCamera     bool `json:"isCamera"`
//	IsFavorite   bool `json:"isFavorite"`
//	IsScreenshot bool `json:"isScreenshot"`
//	IsHidden     bool `json:"isHidden"`
//
//	Albums  []string `json:"albums"`
//	Trips   []string `json:"trips"`
//	Persons []string `json:"persons"`
//
//	CanDelete           bool `json:"canDelete"`
//	CanEditContent      bool `json:"canEditContent"`
//	CanAddToSharedAlbum bool `json:"canAddToSharedAlbum"`
//	IsUserLibraryAsset  bool `json:"IsUserLibraryAsset"`
//
//	CreatedAt time.Time `json:"createdAt"`
//	UpdatedAt time.Time `json:"updatedAt"`
//	DeletedAt time.Time `json:"deletedAt"`
//	Version   string    `json:"version"`
//}
//
//type FileInfo struct {
//	BaseURL  string `json:"baseURL"`
//	FileSize string `json:"fileSize"`
//	FileType string `json:"fileType"`
//	MimeType string `json:"mimeType"`
//}
//
//type ImageInfo struct {
//	Width           int    `json:"width,omitempty"`
//	Height          int    `json:"height,omitempty"`
//	Megapixels      int    `json:"megapixels,omitempty"`
//	Orientation     string `json:"orientation,omitempty"`
//	ColorSpace      string `json:"colorSpace,omitempty"`
//	EncodingProcess string `json:"encodingProcess,omitempty"`
//}
//type CameraInfo struct {
//	Make             string    `json:"make,omitempty"`
//	Model            string    `json:"model,omitempty"`
//	Software         string    `json:"software,omitempty"`
//	DateTimeOriginal time.Time `json:"dateTimeOriginal,omitempty"`
//	ExposureTime     string    `json:"exposureTime,omitempty"`
//	FNumber          float64   `json:"fNumber,omitempty"` // دیافراگم معمولاً float است
//	ISO              int       `json:"iso,omitempty"`     // ISO معمولاً عدد صحیح است
//	FocalLength      string    `json:"focalLength,omitempty"`
//	FocalLength35mm  string    `json:"focalLength35mm,omitempty"`
//	Flash            string    `json:"flash,omitempty"`
//	LightSource      string    `json:"lightSource,omitempty"`
//	ExposureMode     string    `json:"exposureMode,omitempty"`
//	WhiteBalance     string    `json:"whiteBalance,omitempty"`
//}
//
//type VideoInfo struct {
//	MediaDuration      string  `json:"mediaDuration,omitempty"`  // Video duration
//	Width              int     `json:"width,omitempty"`          // Video frame width in pixels
//	Height             int     `json:"height,omitempty"`         // Video frame height in pixels
//	VideoFrameRate     float64 `json:"videoFrameRate,omitempty"` // Video frame rate
//	AvgBitrate         string  `json:"avgBitrate,omitempty"`     // Average bitrate (quality and data volume)
//	Encoder            string  `json:"encoder,omitempty"`        // Video encoding software
//	Rotation           int     `json:"rotation,omitempty"`
//	AudioFormat        string  `json:"audioFormat,omitempty"`
//	AudioChannels      int     `json:"audioChannels,omitempty"`   // Number of audio channels (e.g., 2 for stereo)
//	AudioSampleRate    int     `json:"audioSampleRate,omitempty"` // Audio sample rate
//	AudioBitsPerSample int     `json:"audioBitsPerSample,omitempty"`
//}
//
//type Location struct {
//	Latitude   float64 `json:"location,omitempty"`
//	Longitude  float64 `json:"longitude,omitempty"`
//	Country    string  `json:"country,omitempty"`
//	Province   string  `json:"province,omitempty"`
//	County     string  `json:"county,omitempty"`
//	City       string  `json:"city,omitempty"`
//	Village    string  `json:"village,omitempty"`
//	Malard     string  `json:"malard,omitempty"`
//	Electronic int     `json:"electronic,omitempty"`
//}

//-----------------------------------------------------

//func exampleCamera() {
//
//	info := CameraInfo{
//		Make:     "samsung",
//		Model:    "GT-I9515",
//		Software: "I9515XXU1BPE1",
//		//DateTimeOriginal:  time.Date("2018:03:28 11:14:47"),
//		ExposureTime:    "1/720",
//		FNumber:         2.2,
//		ISO:             50,
//		FocalLength:     "4.2 mm",
//		FocalLength35mm: "31 mm",
//		Flash:           "No Flash",
//		LightSource:     "Unknown",
//		ExposureMode:    "Auto",
//		WhiteBalance:    "Auto",
//	}
//
//	fmt.Println(info.Model)
//}
