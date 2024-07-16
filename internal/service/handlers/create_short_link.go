package handlers

import (
	"net/http"

	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/data"
	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateShortLink(w http.ResponseWriter, r *http.Request) {
    request, err := requests.NewCreateShortLinkRequest(r)
    if err != nil {
        ape.RenderErr(w, problems.BadRequest(err)...)
        return
    }

    shortCode, err := data.CreateShortLink(r.Context(), DB(r), request.OriginalURL)
    if err != nil {
        Log(r).WithError(err).Error("failed to create short link")
        ape.RenderErr(w, problems.InternalError())
        return
    }

    ape.Render(w, map[string]interface{}{"short_code": shortCode})
}
