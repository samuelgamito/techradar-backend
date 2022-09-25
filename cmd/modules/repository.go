package modules

import (
	"backend/internal/repository"
	"go.uber.org/fx"
)

var (
	repositoryFactories = fx.Provide(
		repository.NewTechnologyRepository,
		repository.NewInfoRepository,
	)

	RepositoryModule = fx.Options(
		repositoryFactories,
	)
)
