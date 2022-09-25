package main

import (
	"go.uber.org/fx"
	"techradar-backend/cmd/modules"
	"techradar-backend/internal/bootstrap"
)

func main() {
	fx.New(
		fx.Options(
			modules.UseCase,
			modules.HandlersModule,
			modules.ConfigModule,
			modules.RepositoryModule,
			fx.Invoke(
				bootstrap.RegisterHooks,
			),
		),
	).Run()
}
