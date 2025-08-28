package pinned

import "time"

func (a *Pinned) SetID(id string)          { a.ID = id }
func (a *Pinned) SetCreatedAt(t time.Time) { a.CreatedAt = t }
func (a *Pinned) SetUpdatedAt(t time.Time) { a.UpdatedAt = t }
func (a *Pinned) GetID() string            { return a.ID }
func (a *Pinned) GetCreatedAt() time.Time  { return a.CreatedAt }
func (a *Pinned) GetUpdatedAt() time.Time  { return a.UpdatedAt }

type Pinned struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Type      string    `json:"type"`
	AlbumID   string    `json:"albumID"`
	Icon      string    `json:"icon"`
	Count     int       `json:"count"`
	Index     int       `json:"index"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Version   string    `json:"version"`
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
