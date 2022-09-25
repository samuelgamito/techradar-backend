package modules

import (
	"backend/internal/handler"
	"backend/internal/usecase"
	"go.uber.org/fx"
)

var (
	handlerFactories = fx.Provide(
		handler.NewCreateResource,
		handler.NewFindResource,
		handler.NewInfosHandler,
		fx.Annotate(
			usecase.NewCreateTechnology,
			fx.As(new(handler.CreateTechnologyUseCase)),
		),
		fx.Annotate(
			usecase.NewFindTechnology,
			fx.As(new(handler.FindTechnologyUseCase)),
		),
		fx.Annotate(
			usecase.NewInfoUseCase,
			fx.As(new(handler.InfosUseCase)),
		),
	)
	HandlersModule = fx.Options(
		handlerFactories,
		handler.ModuleHandler,
		handler.ModuleFindTechnology,
		handler.ModuleMoveTechnology,
		handler.ModuleCreateTechnology,
		handler.ModuleUpdateTechnology,
		handler.ModuleActuator,
		handler.ModuleInfos,
	)
)
