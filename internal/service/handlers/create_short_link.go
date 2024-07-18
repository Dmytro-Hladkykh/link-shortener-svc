package handlers

import (
	"database/sql"
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
        Log(r).WithError(err).Error("failed to create request")
        ape.RenderErr(w, problems.BadRequest(err)...)
        return
    }

    db := DB(r)
    if db == nil {
        Log(r).Error("database connection is nil")
        ape.RenderErr(w, problems.InternalError())
        return
    }

    shortLinkQ := pg.NewShortLinkQ(db)
    if shortLinkQ == nil {
        Log(r).Error("shortLinkQ is nil")
        ape.RenderErr(w, problems.InternalError())
        return
    }

    // check for existing link
    existingLink, err := shortLinkQ.FilterByOriginalURL(request.OriginalURL).Get()
    if err != nil && err != sql.ErrNoRows {
        Log(r).WithError(err).Error("failed to check existing link")
        ape.RenderErr(w, problems.InternalError())
        return
    }

    if existingLink != nil {
        // if exist then return short code
        ape.Render(w, map[string]interface{}{"short_code": existingLink.ShortCode})
        return
    }

    // if new link then generate new short code
    shortCode, err := data.GenerateShortCode()
    if err != nil {
        Log(r).WithError(err).Error("failed to generate short code")
        ape.RenderErr(w, problems.InternalError())
        return
    }

    // create new short link in db
    newLink, err := shortLinkQ.Insert(data.ShortLink{
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