package data

import (
	"time"
)

type Link struct {
    ID          int64     `db:"id"`
    OriginalURL string    `db:"original_url"`
    ShortCode   string    `db:"short_code"`
    CreatedAt   time.Time `db:"created_at"`
    Clicks      int       `db:"clicks"`
}