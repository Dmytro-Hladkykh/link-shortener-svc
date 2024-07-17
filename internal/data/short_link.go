package data

import (
	"time"
)

type ShortLink struct {
	ID          int64     `structs:"id" db:"id"`
	OriginalURL string    `structs:"original_url" db:"original_url"`
	ShortCode   string    `structs:"short_code" db:"short_code"`
	CreatedAt   time.Time `structs:"created_at" db:"created_at"`
	Clicks      int       `structs:"clicks" db:"clicks"`
}

type ShortLinkQ interface {
	New() ShortLinkQ

	Get() (*ShortLink, error)
	Select() ([]ShortLink, error)
	Insert(ShortLink) (*ShortLink, error)
	Update(ShortLink) (*ShortLink, error)
	Delete() error

	FilterByShortCode(shortCode string) ShortLinkQ
	FilterByOriginalURL(originalURL string) ShortLinkQ
}
