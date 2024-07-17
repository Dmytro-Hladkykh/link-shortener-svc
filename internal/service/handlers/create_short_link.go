package handlers

import (
	"net/http"

	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/data"
	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/data/pg"
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

    shortLinkQ := pg.NewShortLinkQ(DB(r))

    shortCode, err := data.GenerateShortCode()
    if err != nil {
        Log(r).WithError(err).Error("failed to generate short code")
        ape.RenderErr(w, problems.InternalError())
        return
    }

    shortLink, err := shortLinkQ.Insert(data.ShortLink{
        OriginalURL: request.OriginalURL,
        ShortCode:   shortCode,
    })

    if err != nil {
        Log(r).WithError(err).Error("failed to create short link")
        ape.RenderErr(w, problems.InternalError())
        return
    }

    ape.Render(w, map[string]interface{}{"short_code": shortLink.ShortCode})
}