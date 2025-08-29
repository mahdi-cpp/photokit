package trip

import (
	"strings"

	"github.com/mahdi-cpp/api-go-pkg/search"
)

const MaxLimit = 1000

var LessFunks = map[string]search.LessFunction[*Trip]{
	"id":        func(a, b *Trip) bool { return a.ID < b.ID },
	"createdAt": func(a, b *Trip) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"updatedAt": func(a, b *Trip) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*Trip] {

	fn, exists := LessFunks[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "end" {
		return func(a, b *Trip) bool { return !fn(a, b) }
	}
	return fn
}

func BuildTripSearch(with SearchOptions) search.SearchCriteria[*Trip] {

	return func(a *Trip) bool {

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

func Search(chats []*Trip, with SearchOptions) []*Trip {

	// Build criteria
	criteria := BuildTripSearch(with)

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
	final := make([]*Trip, len(results))
	for i, item := range results {
		final[i] = item.Value
	}

	if with.Limit == 0 { // if not set default is MAX_LIMIT
		with.Limit = MaxLimit
	}

	// Apply pagination
	start := with.Offset
	end := start + with.Limit
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}
