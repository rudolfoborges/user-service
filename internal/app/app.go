package app

import (
	"github.com/amandakeren/user-service/internal/configuration"
	"github.com/amandakeren/user-service/internal/infrastructure"
	"go.uber.org/fx"
)

func Run() {
	app := fx.New(
		infrastructureModule,
		repositoryModule,
		serviceModule,
		controllerModule,

		fx.Decorate(configuration.Routes),

		fx.Invoke(infrastructure.HttpServerInvoke),
	)
	app.Run()
}
