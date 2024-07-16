package main

import (
	"context"
	"net/http"

	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/config"
	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/service"
	"github.com/jmoiron/sqlx"
	"gitlab.com/distributed_lab/kit/kv"
)

func main() {
    getter := kv.MustFromEnv()
    cfg := config.New(getter)

    log := cfg.Logger()

    db, err := sqlx.Connect("postgres", cfg.DSN())
    if err != nil {
        log.WithError(err).Fatal("failed to connect to database")
    }

    ctx := context.Background()
    ctx = service.CtxDB(db)(ctx)
    ctx = service.CtxLog(log)(ctx)

    r := service.Router(cfg)
    log.Info("starting server")

    if err := http.ListenAndServe(cfg.Listen(), r); err != nil {
        log.WithError(err).Fatal("server stopped")
    }
}
