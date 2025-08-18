package build_asset

import (
	"github.com/mahdi-cpp/api-go-pkg/asset"
	"github.com/mahdi-cpp/photokit/internal/search"
	"strings"
)

func SearchAssets(assets []*asset.PHAsset, opts asset.Options) []*asset.PHAsset {

	// Build criteria
	criteria := BuildPhotoSearchCriteria(opts)

	// Execute search
	results := search.Search(assets, criteria)

	// Sort results if needed
	if opts.SortBy != "" {
		lessFn := GetLessFunc(opts.SortBy, opts.SortOrder)
		if lessFn != nil {
			search.SortIndexedItems(results, lessFn)
		}
	}

	// Extract final assets
	final := make([]*asset.PHAsset, len(results))
	for i, item := range results {
		final[i] = item.Value
	}

	// Apply pagination
	start := opts.FetchOffset
	end := start + opts.FetchLimit
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}

func BuildPhotoSearchCriteria(with asset.Options) search.SearchCriteria[*asset.PHAsset] {

	return func(a *asset.PHAsset) bool {

		// Query filter
		if with.Query != "" {
			query := strings.ToLower(with.Query)
			if !strings.Contains(strings.ToLower(a.Filename), query) && !strings.Contains(strings.ToLower(a.Url), query) {
				return false
			}
		}

		// Media type filter
		if with.MediaType != "" && a.MediaType != with.MediaType {
			return false
		}

		// Date range filter
		if !search.TimeInRange(a.CreationDate, *with.StartDate, *with.EndDate) {
			return false
		}

		// Boolean flags filter
		if with.IsFavorite != nil && *with.IsFavorite != a.IsFavorite {
			return false
		}
		if with.IsCamera != nil && *with.IsCamera != a.IsCamera {
			return false
		}
		if with.IsFavorite != nil && *with.IsFavorite != a.IsFavorite {
			return false
		}
		if with.IsScreenshot != nil && *with.IsScreenshot != a.IsScreenshot {
			return false
		}
		if with.IsHidden != nil && *with.IsHidden != a.IsHidden {
			return false
		}
		if with.NotInOneAlbum != nil {
		}
		if with.HideScreenshot != nil && *with.HideScreenshot == false && a.IsScreenshot == true {
			return false
		}

		// Filter by  int
		if with.PixelWidth != 0 && a.PixelWidth != with.PixelWidth {
			return false
		}
		if with.PixelHeight != 0 && a.PixelHeight != with.PixelHeight {
			return false
		}

		// Filter by landscape orientation
		if with.IsLandscape != nil {
			isLandscape := a.PixelWidth > a.PixelHeight
			if isLandscape != *with.IsLandscape {
				return false
			}
		}

		// Collection membership filters
		if len(with.Albums) > 0 {
			found := false
			for _, albumID := range with.Albums {
				if search.StringInSlice(albumID, a.Albums) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}

		// Collection membership filters
		if len(with.Trips) > 0 {
			found := false
			for _, tripID := range with.Trips {
				if search.StringInSlice(tripID, a.Trips) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}

		// Collection membership filters
		if len(with.Persons) > 0 {
			found := false
			for _, personID := range with.Persons {
				if search.StringInSlice(personID, a.Persons) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}

		return true
	}
}

var PHAssetLessFuncs = map[string]search.LessFunction[*asset.PHAsset]{
	"id":               func(a, b *asset.PHAsset) bool { return a.ID < b.ID },
	"capturedDate":     func(a, b *asset.PHAsset) bool { return a.CapturedDate.Before(b.CapturedDate) },
	"creationDate":     func(a, b *asset.PHAsset) bool { return a.CreationDate.Before(b.CreationDate) },
	"modificationDate": func(a, b *asset.PHAsset) bool { return a.ModificationDate.Before(b.ModificationDate) },
	"filename":         func(a, b *asset.PHAsset) bool { return a.Filename < b.Filename },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*asset.PHAsset] {

	fn, exists := PHAssetLessFuncs[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "desc" {
		return func(a, b *asset.PHAsset) bool { return !fn(a, b) }
	}
	return fn
}
