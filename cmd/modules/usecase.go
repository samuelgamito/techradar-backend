package modules

import (
	"go.uber.org/fx"
	"techradar-backend/internal/repository"
	"techradar-backend/internal/usecase"
)

var (
	useCaseFactories = fx.Provide(
		fx.Annotate(
			repository.NewTechnologyRepository,
			fx.As(new(usecase.TechnologyRepository)),
		),
		fx.Annotate(
			repository.NewInfoRepository,
			fx.As(new(usecase.InfoRepository)),
		),
	)

	UseCase = fx.Options(
		useCaseFactories,
	)
)
