package requests

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type CreateShortLinkRequest struct {
    OriginalURL string `json:"original_url"`
}

func NewCreateShortLinkRequest(r *http.Request) (CreateShortLinkRequest, error) {
    var request CreateShortLinkRequest

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        return request, errors.Wrap(err, "failed to unmarshal")
    }

    return request, validateCreateShortLinkRequest(request)
}

func validateCreateShortLinkRequest(request CreateShortLinkRequest) error {
    if request.OriginalURL == "" {
        return errors.New("original URL is required")
    }

    _, err := url.ParseRequestURI(request.OriginalURL)
    if err != nil {
        return errors.Wrap(err, "invalid URL format")
    }

    return nil
}