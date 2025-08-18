package shared_album

import "time"

func (a *SharedAlbum) SetID(id string)          { a.ID = id }
func (a *SharedAlbum) SetCreatedAt(t time.Time) { a.CreatedAt = t }
func (a *SharedAlbum) SetUpdatedAt(t time.Time) { a.UpdatedAt = t }
func (a *SharedAlbum) GetID() string            { return a.ID }
func (a *SharedAlbum) GetCreatedAt() time.Time  { return a.CreatedAt }
func (a *SharedAlbum) GetUpdatedAt() time.Time  { return a.UpdatedAt }

type SharedAlbum struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Subtitle     string    `json:"subtitle"`
	Type         string    `json:"type"`
	Count        int       `json:"count"`
	IsCollection bool      `json:"isCollection"`
	IsHidden     bool      `json:"isHidden"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
	Version      string    `json:"version"`
}

type UpdateOptions struct {
	ID       string `json:"id"`
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Type     string `json:"type,omitempty"`
}

type SearchOptions struct {
	ID       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Type     string `json:"type,omitempty"`

	TextQuery string `json:"textQuery,omitempty"`

	// Date filters
	CreatedAfter  *time.Time `json:"createdAfter,omitempty"`
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`
	ActiveAfter   *time.Time `json:"activeAfter,omitempty"`

	// Pagination
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`

	// Sorting
	SortBy    string `json:"sortBy,omitempty"`    // "title", "created", "members", "lastActivity"
	SortOrder string `json:"sortOrder,omitempty"` // "asc" or "desc"
}
