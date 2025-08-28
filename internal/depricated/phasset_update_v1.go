package depricated

//import (
//	//"github.com/mahdi-cpp/api-go-pkg/asset"
//	"time"
//)
//
//func updateProcess(asset *asset.PHAsset, update asset.Update) *asset.PHAsset {
//
//	// Apply updates
//	if update.Filename != nil {
//		asset.Filename = *update.Filename
//	}
//	if update.MediaType != "" {
//		asset.MediaType = update.MediaType
//	}
//	if update.CameraMake != nil {
//		asset.CameraMake = *update.CameraMake
//	}
//	if update.CameraModel != nil {
//		asset.CameraModel = *update.CameraModel
//	}
//	if update.IsCamera != nil {
//		asset.IsCamera = *update.IsCamera
//	}
//	if update.IsFavorite != nil {
//		asset.IsFavorite = *update.IsFavorite
//	}
//	if update.IsScreenshot != nil {
//		asset.IsScreenshot = *update.IsScreenshot
//	}
//	if update.IsHidden != nil {
//		asset.IsHidden = *update.IsHidden
//	}
//
//	// Handle album operations
//	switch {
//	case update.Albums != nil:
//		// Full replacement
//		asset.Albums = *update.Albums
//	case len(update.AddAlbums) > 0 || len(update.RemoveAlbums) > 0:
//
//		// Handler a set for efficient lookups
//		albumSet := make(map[string]bool)
//		for _, id := range asset.Albums {
//			albumSet[id] = true
//		}
//
//		// Add new items (avoid duplicates)
//		for _, id := range update.AddAlbums {
//			if !albumSet[id] {
//				asset.Albums = append(asset.Albums, id)
//				albumSet[id] = true
//			}
//		}
//
//		// Remove specified items
//		if len(update.RemoveAlbums) > 0 {
//			removeSet := make(map[string]bool)
//			for _, id := range update.RemoveAlbums {
//				removeSet[id] = true
//			}
//
//			newAlbums := make([]string, 0, len(asset.Albums))
//			for _, id := range asset.Albums {
//				if !removeSet[id] {
//					newAlbums = append(newAlbums, id)
//				}
//			}
//			asset.Albums = newAlbums
//		}
//	}
//
//	// Handle trip operations
//	switch {
//	case update.Trips != nil:
//		// Full replacement
//		asset.Trips = *update.Trips
//	case len(update.AddTrips) > 0 || len(update.RemoveTrips) > 0:
//
//		// Handler a set for efficient lookups
//		tripSet := make(map[string]bool)
//		for _, id := range asset.Trips {
//			tripSet[id] = true
//		}
//
//		// Add new Persons (avoid duplicates)
//		for _, id := range update.AddTrips {
//			if !tripSet[id] {
//				asset.Trips = append(asset.Trips, id)
//				tripSet[id] = true
//			}
//		}
//
//		// Remove specified trips
//		if len(update.RemoveTrips) > 0 {
//			removeSet := make(map[string]bool)
//			for _, id := range update.RemoveTrips {
//				removeSet[id] = true
//			}
//
//			newTrips := make([]string, 0, len(asset.Trips))
//			for _, id := range asset.Trips {
//				if !removeSet[id] {
//					newTrips = append(newTrips, id)
//				}
//			}
//			asset.Trips = newTrips
//		}
//	}
//
//	// Handle person operations
//	switch {
//	case update.Persons != nil:
//		// Full replacement
//		asset.Persons = *update.Persons
//	case len(update.AddPersons) > 0 || len(update.RemovePersons) > 0:
//
//		// Handler a set for efficient lookups
//		personSet := make(map[string]bool)
//		for _, id := range asset.Persons {
//			personSet[id] = true
//		}
//
//		// Add new Persons (avoid duplicates)
//		for _, id := range update.AddPersons {
//			if !personSet[id] {
//				asset.Persons = append(asset.Persons, id)
//				personSet[id] = true
//			}
//		}
//
//		// Remove specified Persons
//		if len(update.RemovePersons) > 0 {
//			removeSet := make(map[string]bool)
//			for _, id := range update.RemovePersons {
//				removeSet[id] = true
//			}
//
//			newPersons := make([]string, 0, len(asset.Persons))
//			for _, id := range asset.Persons {
//				if !removeSet[id] {
//					newPersons = append(newPersons, id)
//				}
//			}
//			asset.Persons = newPersons
//		}
//	}
//
//	asset.ModificationDate = time.Now()
//
//	return asset
//
//}
