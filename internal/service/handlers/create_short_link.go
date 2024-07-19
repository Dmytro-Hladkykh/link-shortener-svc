package handlers

import (
	"net/http"

	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/data"
	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

type CreateShortLinkHandler struct {
	repo data.ShortLinkQ
}

func NewCreateShortLinkHandler(repo data.ShortLinkQ) *CreateShortLinkHandler {
	return &CreateShortLinkHandler{repo: repo}
}

func (h *CreateShortLinkHandler) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateShortLinkRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to create request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// check if link already exists
	existingLink, err := h.repo.FilterByOriginalURL(request.OriginalURL).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to check existing link")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if existingLink != nil {
		// if exist then return short code
		ape.Render(w, map[string]interface{}{"short_code": existingLink.ShortCode})
		return
	}

	// if new link then generate short code
	shortCode, err := data.GenerateShortCode()
	if err != nil {
		Log(r).WithError(err).Error("failed to generate short code")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// create new link in db
	newLink, err := h.repo.Insert(data.ShortLink{
		OriginalURL: request.OriginalURL,
		ShortCode:   shortCode,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to insert new short link")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, map[string]interface{}{"short_code": newLink.ShortCode})
}
