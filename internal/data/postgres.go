package data

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func CreateShortLink(ctx context.Context, db *sqlx.DB, originalURL string) (string, error) {
    shortCode, err := GenerateShortCode()
    if err != nil {
        return "", errors.Wrap(err, "failed to generate short code")
    }

    _, err = db.ExecContext(ctx, 
        "INSERT INTO links (original_url, short_code) VALUES ($1, $2)",
        originalURL, shortCode)
    if err != nil {
        return "", errors.Wrap(err, "failed to insert link")
    }
    return shortCode, nil
}

func GetOriginalURL(ctx context.Context, db *sqlx.DB, shortCode string) (string, error) {
    var originalURL string
    err := db.GetContext(ctx, &originalURL,
        "UPDATE links SET clicks = clicks + 1 WHERE short_code = $1 RETURNING original_url",
        shortCode)
    if err == sql.ErrNoRows {
        return "", errors.New("link not found")
    }
    if err != nil {
        return "", errors.Wrap(err, "failed to get original URL")
    }
    return originalURL, nil
}