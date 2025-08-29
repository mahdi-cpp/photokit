package person

import (
	"time"

	"github.com/mahdi-cpp/api-go-pkg/update"
)

// Initialize updater
var metadataUpdater = update.NewUpdater[Person, UpdateOptions]()

func init() {

	// Configure scalar field updates
	metadataUpdater.AddScalarUpdater(func(a *Person, u UpdateOptions) {
		if u.Title != "" {
			a.Title = u.Title
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *Person, u UpdateOptions) {
		if u.Subtitle != "" {
			a.Subtitle = u.Subtitle
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *Person, u UpdateOptions) {
		if u.Type != "" {
			a.Type = u.Type
		}
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *Person) {
		a.UpdatedAt = time.Now()
	})
}

func Update(item *Person, update UpdateOptions) *Person {
	metadataUpdater.Apply(item, update)
	return item
}
