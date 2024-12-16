package models

import "time"

// Feed represents an RSS feed.
type Feed struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	PublishedAt time.Time `json:"published_at"`
	Read        bool      `json:"read"`
}
