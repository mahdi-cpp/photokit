package village

import "time"

func (a *Village) SetID(id string)          { a.ID = id }
func (a *Village) SetCreatedAt(t time.Time) { a.CreatedAt = t }
func (a *Village) SetUpdatedAt(t time.Time) { a.UpdatedAt = t }
func (a *Village) GetID() string            { return a.ID }
func (a *Village) GetCreatedAt() time.Time  { return a.CreatedAt }
func (a *Village) GetUpdatedAt() time.Time  { return a.UpdatedAt }

type Village struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Version   string    `json:"version"`
}

type Polygon struct {
	Name string        `json:"name"`
	Data [][][]float64 `json:"data"`
}
