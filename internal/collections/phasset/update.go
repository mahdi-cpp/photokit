package phasset

import (
	"github.com/mahdi-cpp/api-go-pkg/update"
	"time"
)

// Initialize updater
var metadataUpdater = update.NewUpdater[PHAsset, UpdateOptions]()

func init() {

	// Configure scalar field updates
	metadataUpdater.AddScalarUpdater(func(a *PHAsset, u UpdateOptions) {
		if u.FileName != "" {
			a.FileName = u.FileName
		}
	})

	metadataUpdater.AddScalarUpdater(func(a *PHAsset, u UpdateOptions) {
		if u.FileSize != "" {
			a.FileInfo.FileSize = u.FileSize
		}
	})

	// Add other scalar fields similarly...

	// Configure collection operations
	metadataUpdater.AddCollectionUpdater(func(a *PHAsset, u UpdateOptions) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Albums,
			Add:         u.AddAlbums,
			Remove:      u.RemoveAlbums,
		}
		a.Albums = update.ApplyCollectionUpdate(a.Albums, op)
	})

	metadataUpdater.AddCollectionUpdater(func(a *PHAsset, u UpdateOptions) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Trips,
			Add:         u.AddTrips,
			Remove:      u.RemoveTrips,
		}
		a.Trips = update.ApplyCollectionUpdate(a.Trips, op)
	})

	metadataUpdater.AddCollectionUpdater(func(a *PHAsset, u UpdateOptions) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.Persons,
			Add:         u.AddPersons,
			Remove:      u.RemovePersons,
		}
		a.Persons = update.ApplyCollectionUpdate(a.Persons, op)
	})

	// Set modification timestamp
	metadataUpdater.AddPostUpdateHook(func(a *PHAsset) {
		a.UpdatedAt = time.Now()
	})
}

func Update(p *PHAsset, update UpdateOptions) *PHAsset {
	metadataUpdater.Apply(p, update)
	return p
}

//func (p *PHAsset) Save() error {
//	p.mutex.Lock()
//	defer p.mutex.Unlock()
//
//	return utils.WriteData(p, p.Filepath)
//}

// IsEmpty checks if the Place struct contains zero values for all its fields.
func (p Place) IsEmpty() bool {
	return p.Latitude == 0.0 &&
		p.Longitude == 0.0 &&
		p.City == "" &&
		p.Country == ""
}
