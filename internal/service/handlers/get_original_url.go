package handlers

import (
	"net/http"

	"github.com/Dmytro-Hladkykh/link-shortener-svc/internal/data"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
    shortCode := chi.URLParam(r, "shortCode")
    
    originalURL, err := data.GetOriginalURL(r.Context(), DB(r), shortCode)
    if err != nil {
        Log(r).WithError(err).Error("failed to get original URL")
        ape.RenderErr(w, problems.NotFound())
        return
    }

    ape.Render(w, map[string]interface{}{"original_url": originalURL})
}
