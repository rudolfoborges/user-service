package infrastructure

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func HttpServeMuxProvider() *chi.Mux {
	return chi.NewRouter()
}

func HttpServerProvider(lc fx.Lifecycle, mux *chi.Mux) *http.Server {
	port := os.Getenv("PORT")
	srv := &http.Server{Addr: port, Handler: mux}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			slog.Info(fmt.Sprintf("Listening on %s", port))
			go func() {
				if err := srv.ListenAndServe(); err != nil {
					slog.Error("Error starting server", "error", err)
					os.Exit(1)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			slog.Info("Shutting down the server")
			return nil
		},
	})

	return srv
}

func HttpServerInvoke(*http.Server) {}
