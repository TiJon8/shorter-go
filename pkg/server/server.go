package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/TiJon8/shorter-go/pkg/config"
	"github.com/TiJon8/shorter-go/pkg/logger"
	"github.com/TiJon8/shorter-go/pkg/storage"

	"github.com/go-chi/chi/v5"
)


type HTTPServer struct {
	config *config.ServerConfig
	mux *chi.Mux
	logger *slog.Logger
	storage *storage.Storage
}

func NewHTTPServer(cfg *config.ServerConfig, log *logger.Logger, router *chi.Mux, st *storage.Storage) *HTTPServer {
	var l *slog.Logger
	if log == nil {
		l = logger.DefaultLogger
	} else {
		l = log.Logger
	}
	return  &HTTPServer{
		config: cfg,
		mux: router,
		logger: l,
		storage: st,
	}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	server := http.Server{
		Addr: s.config.Addr,
		Handler: s.mux,
	}
	defer s.storage.Close()

	ch := make(chan error, 1)

	go func() {
		s.logger.Warn("Server has started on", slog.String("port", s.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("Error on bootstrap server: %w", err)
		}
	case <-ctx.Done():
		now := time.Now()
		s.logger.Warn("Server is stopping", slog.Time("now", now))
		ctx, cancel := context.WithTimeout(context.Background(), s.config.GracefullShutdownDuration)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			server.Close()
			return fmt.Errorf("Server hasn't stopped in 9 seconds so it was closed immediately: %w", err)
		}
		s.logger.Warn("Server succesfully was stopped", slog.Duration("for", time.Since(now)))
	}
	return nil
}