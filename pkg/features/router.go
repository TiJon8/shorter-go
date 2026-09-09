package features

import (
	transport "github.com/TiJon8/shorter-go/pkg/features/url/http"
	"github.com/TiJon8/shorter-go/pkg/features/url/service"
	"github.com/TiJon8/shorter-go/pkg/logger"
	"github.com/TiJon8/shorter-go/pkg/storage"

	mid "github.com/TiJon8/shorter-go/pkg/server/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func InitRouter(storage *storage.Storage, l *logger.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(mid.Logger(l.Logger))
	router.Use(middleware.Recoverer)

	router.Get("/ping", Ping)
	
	
	Service := service.InitService(storage)
	HttpHandlers := transport.InitHTTPHandlers(Service)
	
	router.Get("/{alias}", HttpHandlers.RedirectURL)
	router.Route("/short", func(r chi.Router) {
		r.Post("/", HttpHandlers.SaveURL)
		r.Patch("/", HttpHandlers.PatchURL)
		r.Delete("/", HttpHandlers.DeleteURL)
	})

	return router
}