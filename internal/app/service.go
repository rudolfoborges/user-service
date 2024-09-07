package app

import (
	"github.com/amandakeren/user-service/internal/service"
	"go.uber.org/fx"
)

var serviceModule = fx.Module(
	"service",
	fx.Provide(
		service.NewUserService,
		service.NewSessionService,
	),
)
