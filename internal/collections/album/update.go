package album

import (
	"time"

	"github.com/mahdi-cpp/api-go-pkg/update"
)

// Initialize updater
var metadataUpdater = update.NewUpdater[Album, UpdateOptions]()

func init() {

	// Configure scalar field updates
	metadataUpdater.AddScalarUpdater(func(a *Album, u UpdateOptions) {
		if u.Title != "" {
			a.Title = u.Title
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *Album, u UpdateOptions) {
		if u.Subtitle != "" {
			a.Subtitle = u.Subtitle
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *Album, u UpdateOptions) {
		if u.Type != "" {
			a.Type = u.Type
		}
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Album) {
		a.UpdatedAt = time.Now()
	})
}

func Update(item *Album, update UpdateOptions) *Album {
	metadataUpdater.Apply(item, update)
	return item
}
