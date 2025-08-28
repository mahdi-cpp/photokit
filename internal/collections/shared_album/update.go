package shared_album

import (
	"github.com/mahdi-cpp/api-go-pkg/update"
	"time"
)

// Initialize updater
var metadataUpdater = update.NewUpdater[SharedAlbum, UpdateOptions]()

func init() {

	// Configure scalar field updates
	metadataUpdater.AddScalarUpdater(func(a *SharedAlbum, u UpdateOptions) {
		if u.Title != "" {
			a.Title = u.Title
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *SharedAlbum, u UpdateOptions) {
		if u.Subtitle != "" {
			a.Subtitle = u.Subtitle
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *SharedAlbum, u UpdateOptions) {
		if u.Type != "" {
			a.Type = u.Type
		}
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *SharedAlbum) {
		a.UpdatedAt = time.Now()
	})
}

func Update(item *SharedAlbum, update UpdateOptions) *SharedAlbum {
	metadataUpdater.Apply(item, update)
	return item
}
