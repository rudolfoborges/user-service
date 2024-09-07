package app

import (
	"github.com/amandakeren/user-service/internal/infrastructure"
	"go.uber.org/fx"
)

var infrastructureModule = fx.Module(
	"infrastructure",
	fx.Provide(
		infrastructure.DatabaseProvider,
		infrastructure.HttpServeMuxProvider,
		infrastructure.HttpServerProvider,
	),
)
