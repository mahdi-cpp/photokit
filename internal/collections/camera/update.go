package camera

import (
	"github.com/mahdi-cpp/photokit/internal/update"
	"time"
)

// Initialize updater
var metadataUpdater = update.NewUpdater[Camera, UpdateOptions]()

func init() {

	// Configure scalar field updates
	metadataUpdater.AddScalarUpdater(func(a *Camera, u UpdateOptions) {
		if u.Title != "" {
			a.Title = u.Title
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *Camera, u UpdateOptions) {
		if u.Subtitle != "" {
			a.Subtitle = u.Subtitle
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *Camera, u UpdateOptions) {
		if u.Type != "" {
			a.Type = u.Type
		}
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Camera) {
		a.UpdatedAt = time.Now()
	})
}

func Update(item *Camera, update UpdateOptions) *Camera {
	metadataUpdater.Apply(item, update)
	return item
}
