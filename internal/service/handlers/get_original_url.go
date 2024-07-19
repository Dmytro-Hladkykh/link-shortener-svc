package handlers

import (
	"net/http"

	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/data"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

type GetOriginalURLHandler struct {
	repo data.ShortLinkQ
}

func NewGetOriginalURLHandler(repo data.ShortLinkQ) *GetOriginalURLHandler {
	return &GetOriginalURLHandler{repo: repo}
}

func (h *GetOriginalURLHandler) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")

	shortLink, err := h.repo.FilterByShortCode(shortCode).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get original URL")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if shortLink == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, map[string]interface{}{"original_url": shortLink.OriginalURL})
}
