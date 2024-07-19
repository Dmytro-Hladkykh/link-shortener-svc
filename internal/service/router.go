package service

import (
	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/config"
	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config, createShortLinkHandler *handlers.CreateShortLinkHandler, getOriginalURLHandler *handlers.GetOriginalURLHandler) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxDB(cfg.DB()),
		),
	)

	r.Route("/link-shortener", func(r chi.Router) {
		r.Post("/", createShortLinkHandler.CreateShortLink)
		r.Get("/{shortCode}", getOriginalURLHandler.GetOriginalURL)
	})

	return r
}
