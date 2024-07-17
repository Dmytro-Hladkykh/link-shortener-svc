package pg

import (
	"database/sql"

	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/data"
	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const shortLinkTableName = "shortened_urls"

func NewShortLinkQ(db *pgdb.DB) data.ShortLinkQ {
    return &shortLinkQ{
        db:  db,
        sql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
    }
}

type shortLinkQ struct {
    db  *pgdb.DB
    sql sq.StatementBuilderType
}

func (q *shortLinkQ) New() data.ShortLinkQ {
    return NewShortLinkQ(q.db)
}

func (q *shortLinkQ) Get() (*data.ShortLink, error) {
    var result data.ShortLink
    stmt := q.sql.Select("*").From(shortLinkTableName)
    err := q.db.Get(&result, stmt)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, errors.Wrap(err, "failed to get short link from db")
    }
    return &result, nil
}

func (q *shortLinkQ) Select() ([]data.ShortLink, error) {
    var result []data.ShortLink
    stmt := q.sql.Select("*").From(shortLinkTableName)
    err := q.db.Select(&result, stmt)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, errors.Wrap(err, "failed to select short links from db")
    }
    return result, nil
}

func (q *shortLinkQ) Insert(link data.ShortLink) (*data.ShortLink, error) {
    clauses := map[string]interface{}{
        "original_url": link.OriginalURL,
        "short_code":   link.ShortCode,
    }
    var result data.ShortLink
    stmt := sq.Insert(shortLinkTableName).SetMap(clauses).Suffix("RETURNING *")
    err := q.db.Get(&result, stmt)
    if err != nil {
        return nil, errors.Wrap(err, "failed to insert short link to db")
    }
    return &result, nil
}

func (q *shortLinkQ) Update(link data.ShortLink) (*data.ShortLink, error) {
    clauses := map[string]interface{}{
        "original_url": link.OriginalURL,
        "short_code":   link.ShortCode,
    }
    var result data.ShortLink
    stmt := q.sql.Update(shortLinkTableName).SetMap(clauses).Where(sq.Eq{"id": link.ID}).Suffix("RETURNING *")
    err := q.db.Get(&result, stmt)
    if err != nil {
        return nil, errors.Wrap(err, "failed to update short link in db")
    }
    return &result, nil
}

func (q *shortLinkQ) Delete() error {
    stmt := q.sql.Delete(shortLinkTableName)
    err := q.db.Exec(stmt)
    if err != nil {
        return errors.Wrap(err, "failed to delete short links from db")
    }
    return nil
}

func (q *shortLinkQ) FilterByShortCode(shortCode string) data.ShortLinkQ {
    q.sql = q.sql.Where(sq.Eq{"short_code": shortCode})
    return q
}

func (q *shortLinkQ) FilterByOriginalURL(originalURL string) data.ShortLinkQ {
    q.sql = q.sql.Where(sq.Eq{"original_url": originalURL})
    return q
}
