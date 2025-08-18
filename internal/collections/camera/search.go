package camera

import (
	"github.com/mahdi-cpp/photokit/internal/search"
	"strings"
)

var LessFuncs = map[string]search.LessFunction[*Camera]{
	"id":        func(a, b *Camera) bool { return a.ID < b.ID },
	"createdAt": func(a, b *Camera) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"updatedAt": func(a, b *Camera) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*Camera] {

	fn, exists := LessFuncs[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "end" {
		return func(a, b *Camera) bool { return !fn(a, b) }
	}
	return fn
}

func BuildCameraSearch(with SearchOptions) search.SearchCriteria[*Camera] {

	return func(a *Camera) bool {

		// ID filter
		if with.ID != "" && a.ID != with.ID {
			return false
		}

		// Title search_manager (case-insensitive)
		if with.TextQuery != "" {
			query := strings.ToLower(with.Title)
			title := strings.ToLower(a.Subtitle)
			if !strings.Contains(title, query) {
				return false
			}
		}

		// Date filters
		if with.CreatedAfter != nil && a.CreatedAt.Before(*with.CreatedAfter) {
			return false
		}
		if with.CreatedBefore != nil && a.CreatedAt.After(*with.CreatedBefore) {
			return false
		}

		return true
	}
}

func Search(chats []*Camera, with SearchOptions) []*Camera {

	// Build criteria
	criteria := BuildCameraSearch(with)

	// Execute search_manager
	results := search.Search(chats, criteria)

	// Sort results if needed
	if with.SortBy != "" {
		lessFn := GetLessFunc(with.SortBy, with.SortOrder)
		if lessFn != nil {
			search.SortIndexedItems(results, lessFn)
		}
	}

	// Extract final assets
	final := make([]*Camera, len(results))
	for i, item := range results {
		final[i] = item.Value
	}

	// Apply pagination
	start := with.Offset
	end := start + with.Limit
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}
