package handlers

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
    logCtxKey ctxKey = iota
    dbCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
    return func(ctx context.Context) context.Context {
        return context.WithValue(ctx, logCtxKey, entry)
    }
}

func Log(r *http.Request) *logan.Entry {
    return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxDB(db *sqlx.DB) func(context.Context) context.Context {
    return func(ctx context.Context) context.Context {
        return context.WithValue(ctx, dbCtxKey, db)
    }
}

func DB(r *http.Request) *sqlx.DB {
    return r.Context().Value(dbCtxKey).(*sqlx.DB)
}
