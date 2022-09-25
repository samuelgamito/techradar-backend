package modules

import (
	"backend/internal/repository"
	"backend/internal/usecase"
	"go.uber.org/fx"
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
