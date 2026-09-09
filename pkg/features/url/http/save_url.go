package transport

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)


type SaveURLRequest struct {
	Url string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type SaveURLResponse struct {
	Url string `json:"url"`
	Alias string `json:"alias"`
	RedirectUrl string `json:"redirect_url"`
}

func (h *HTTPHandlers) SaveURL(w http.ResponseWriter, r *http.Request) {
	var req SaveURLRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		Response(w, http.StatusInternalServerError, ErrorResponse("failed to decode request body", err))
		return
	}

	if err := validator.New().Struct(req); err != nil {
		Response(w, http.StatusBadRequest, ErrorResponse("failed to validate request body", err))
		return
	}
	saved, alias, redirectedUrl, err := h.service.HandleSaveURL(req.Url, req.Alias)
	if err != nil {
		Response(w, http.StatusBadRequest, ErrorResponse("failed to short url", err))
		return
	}
	Response(w, http.StatusCreated, SaveURLResponse{
		Url: saved,
		Alias: alias,
		RedirectUrl: redirectedUrl,
	})
}