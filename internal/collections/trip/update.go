package trip

import (
	"github.com/mahdi-cpp/photokit/internal/update"
	"time"
)

// Initialize updater
var metadataUpdater = update.NewUpdater[Trip, UpdateOptions]()

func init() {

	// Configure scalar field updates
	metadataUpdater.AddScalarUpdater(func(a *Trip, u UpdateOptions) {
		if u.Title != "" {
			a.Title = u.Title
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *Trip, u UpdateOptions) {
		if u.Subtitle != "" {
			a.Subtitle = u.Subtitle
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *Trip, u UpdateOptions) {
		if u.Type != "" {
			a.Type = u.Type
		}
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Trip) {
		a.UpdatedAt = time.Now()
	})
}

func Update(item *Trip, update UpdateOptions) *Trip {
	metadataUpdater.Apply(item, update)
	return item
}
