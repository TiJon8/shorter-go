package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// func Redirect(urlRepo UrlRepository) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		logger := r.Context().Value(middleware.LoggerContextKey).(*slog.Logger)
// 		short_alias := chi.URLParam(r, "alias")
// 		fmt.Println(logger)
// 		logger.Info("", slog.String("alias", short_alias))
// 		r.Context().Done()

// 		surl, err := urlRepo.GetURL(short_alias)
// 		if err != nil {
// 			render.JSON(w, r, ErrorResponse("failed to get source url by provided alias", err))
// 			return
// 		}

// 		http.Redirect(w, r, surl, 302)
// 	}
// }


func (h *HTTPHandlers) RedirectURL(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")
	url, err := h.service.HandleRedirectURL(alias)
	if err != nil {
		Response(w, http.StatusNotFound, ErrorResponse("failed to get source url by provided alias", err))
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}