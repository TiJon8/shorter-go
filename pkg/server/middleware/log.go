package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/TiJon8/shorter-go/pkg/logger"
	"github.com/go-chi/chi/v5/middleware"
)

type loggerContextKeyType struct {}

var (
	LoggerContextKey loggerContextKeyType
)

type Middleware func(next http.Handler) http.Handler


func Logger(l *logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var log *slog.Logger
			if l == nil {
				log = logger.DefaultLogger
			} else {
				log = l.Logger
			}
			log = log.With(slog.String("component", "middleware/logger"))
			log.Info("Middleware logger has enabled")

			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_address", r.RemoteAddr),
				slog.String("request_id", middleware.GetReqID(r.Context())),
			)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			ctx := context.WithValue(r.Context(), LoggerContextKey, entry)
			now := time.Now()
			defer func ()  {
				entry.Info("request complete", 
					slog.Int("status_code", ww.Status()),
					slog.Int("bytes_written", ww.BytesWritten()),
					slog.Duration("duration", time.Since(now)),
				)	
			}()
			next.ServeHTTP(ww, r.WithContext(ctx))
		})
	}
}