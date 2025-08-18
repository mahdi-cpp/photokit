package build_asset

import (
	"github.com/mahdi-cpp/api-go-pkg/asset"
	"github.com/mahdi-cpp/photokit/internal/update"
	"time"
)

// Initialize updater
var assetUpdater = update.NewUpdater[asset.PHAsset, asset.Update]()

func init() {

	// Configure scalar field updates
	assetUpdater.AddScalarUpdater(func(a *asset.PHAsset, u asset.Update) {
		if u.Filename != nil {
			a.Filename = *u.Filename
		}
	})

	assetUpdater.AddScalarUpdater(func(a *asset.PHAsset, u asset.Update) {
		if u.MediaType != "" {
			a.MediaType = u.MediaType
		}
	})

	// Add other scalar fields similarly...

	// Configure collection operations
	assetUpdater.AddCollectionUpdater(func(a *asset.PHAsset, u asset.Update) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Albums,
			Add:         u.AddAlbums,
			Remove:      u.RemoveAlbums,
		}
		a.Albums = update.ApplyCollectionUpdate(a.Albums, op)
	})

	assetUpdater.AddCollectionUpdater(func(a *asset.PHAsset, u asset.Update) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Trips,
			Add:         u.AddTrips,
			Remove:      u.RemoveTrips,
		}
		a.Trips = update.ApplyCollectionUpdate(a.Trips, op)
	})

	assetUpdater.AddCollectionUpdater(func(a *asset.PHAsset, u asset.Update) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Persons,
			Add:         u.AddPersons,
			Remove:      u.RemovePersons,
		}
		a.Persons = update.ApplyCollectionUpdate(a.Persons, op)
	})

	// Set modification timestamp
	assetUpdater.AddPostUpdateHook(func(a *asset.PHAsset) {
		a.ModificationDate = time.Now()
	})
}

// UpdateProcess Generic update processor
func UpdateProcess(asset *asset.PHAsset, update asset.Update) *asset.PHAsset {
	assetUpdater.Apply(asset, update)
	return asset
}
