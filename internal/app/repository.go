package app

import (
	"github.com/amandakeren/user-service/internal/gateway/repository"
	"go.uber.org/fx"
)

var repositoryModule = fx.Module(
	"repository",
	fx.Provide(
		repository.NewUserRepository,
	),
)
