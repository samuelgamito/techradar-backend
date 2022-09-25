package modules

import (
	"go.uber.org/fx"
	"techradar-backend/internal/repository"
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
