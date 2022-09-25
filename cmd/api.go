package main

import (
	modules2 "backend/cmd/modules"
	"backend/internal/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			modules2.UseCase,
			modules2.HandlersModule,
			modules2.ConfigModule,
			modules2.RepositoryModule,
			fx.Invoke(
				bootstrap.RegisterHooks,
			),
		),
	).Run()
}
