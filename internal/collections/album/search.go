package album

import (
	"github.com/mahdi-cpp/photokit/internal/search"
	"strings"
)

var LessFuncs = map[string]search.LessFunction[*Album]{
	"id":        func(a, b *Album) bool { return a.ID < b.ID },
	"createdAt": func(a, b *Album) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"updatedAt": func(a, b *Album) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*Album] {

	fn, exists := LessFuncs[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "end" {
		return func(a, b *Album) bool { return !fn(a, b) }
	}
	return fn
}

func BuildAlbumSearch(with SearchOptions) search.SearchCriteria[*Album] {

	return func(a *Album) bool {

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

func Search(chats []*Album, with SearchOptions) []*Album {

	// Build criteria
	criteria := BuildAlbumSearch(with)

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
	final := make([]*Album, len(results))
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
