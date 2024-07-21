// internal/service/handlers/get_original_url.go

package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	db := DB(r)

	shortCode := chi.URLParam(r, "shortCode")

	shortLink, err := db.ShortLink().FilterByShortCode(shortCode).Get()
	if err != nil {
		log.WithError(err).Error("failed to get original URL")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if shortLink == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, map[string]interface{}{"original_url": shortLink.OriginalURL})
}
