package transport

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)


type PatchURLRequest struct {
	NewUrl string `json:"url" validate:"required,url"`
	Alias string `json:"alias" validate:"required"`
}

type PatchURLResponse struct {
	ActualUrl string `json:"actual_url"`
	Alias string `json:"alias"`
	RedirectUrl string `json:"redirect_url"`
}

func (h *HTTPHandlers) PatchURL(w http.ResponseWriter, r *http.Request) {
	var req PatchURLRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		Response(w, http.StatusInternalServerError, ErrorResponse("failed to decode request body", err))
		return
	}

	if err := validator.New().Struct(req); err != nil {
		Response(w, http.StatusBadRequest, ErrorResponse("failed to validate request body", err))
		return
	}

	newUrl, redirectUrl, err := h.service.HandlePatchURL(req.NewUrl, req.Alias)
	if err != nil {
		Response(w, http.StatusNotFound, ErrorResponse("failed to patch url", err))
		return
	}
	Response(w, http.StatusOK, PatchURLResponse{ActualUrl: newUrl, Alias: req.Alias, RedirectUrl: redirectUrl})
}