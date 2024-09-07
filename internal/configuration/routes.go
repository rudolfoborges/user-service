package configuration

import (
	"time"

	"github.com/amandakeren/user-service/internal/entrypoint/controller"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

func Routes(
	mux *chi.Mux,
	userController *controller.UserController,
	sessionController *controller.SessionController,
) *chi.Mux {
	mux.Use(middleware.Logger)
	mux.Use(middleware.Heartbeat("/heartbeat"))
	mux.Use(httprate.LimitByIP(500, time.Minute))

	userController.Routes(mux)
	sessionController.Routes(mux)

	return mux
}
