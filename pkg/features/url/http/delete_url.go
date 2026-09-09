package transport

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type DeleteURLRequest struct {
	Alias string `json:"alias" validate:"required"`
}

type DeleteURLResponse struct {
	OK bool `json:"ok"`
}

func (h *HTTPHandlers) DeleteURL(w http.ResponseWriter, r *http.Request) {
	var req DeleteURLRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		Response(w, http.StatusInternalServerError, ErrorResponse("failed to decode request body", err))
		return
	}

	if err := validator.New().Struct(req); err != nil {
		Response(w, http.StatusBadRequest, ErrorResponse("failed to validate request body", err))
		return
	}

	if err := h.service.HandleDeleteURL(req.Alias); err != nil {
		Response(w, http.StatusNotFound, ErrorResponse("failed to delete url", err))
		return
	}
	Response(w, http.StatusNoContent, DeleteURLResponse{OK: true})
}