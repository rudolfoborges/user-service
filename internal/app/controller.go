package app

import (
	"github.com/amandakeren/user-service/internal/entrypoint/controller"
	"go.uber.org/fx"
)

var controllerModule = fx.Module(
	"controller",
	fx.Provide(
		controller.NewUserController,
	),
)
