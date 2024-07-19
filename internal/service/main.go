package service

import (
	"net"
	"net/http"

	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/config"
	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/data/pg"
	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
}

func (s *service) run(cfg config.Config) error {
	s.log.Info("Service started")

	// create db connection
	db := cfg.DB()

	// create repo
	shortLinkRepo := pg.NewShortLinkQ(db)

	// create handlers
	createShortLinkHandler := handlers.NewCreateShortLinkHandler(shortLinkRepo)
	getOriginalURLHandler := handlers.NewGetOriginalURLHandler(shortLinkRepo)

	r := s.router(cfg, createShortLinkHandler, getOriginalURLHandler)

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(cfg); err != nil {
		panic(err)
	}
}
