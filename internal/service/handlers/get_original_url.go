package handlers

import (
	"net/http"

	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/data/pg"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
    shortCode := chi.URLParam(r, "shortCode")
    
    shortLinkQ := pg.NewShortLinkQ(DB(r))
    
    shortLink, err := shortLinkQ.FilterByShortCode(shortCode).Get()
    if err != nil {
        Log(r).WithError(err).Error("failed to get original URL")
        ape.RenderErr(w, problems.NotFound())
        return
    }

    if shortLink == nil {
        ape.RenderErr(w, problems.NotFound())
        return
    }

    ape.Render(w, map[string]interface{}{"original_url": shortLink.OriginalURL})
}
