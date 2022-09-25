package modules

import (
	"backend/internal/config"
	"go.uber.org/fx"
)

var (
	configFactories = fx.Provide(
		config.CreateConfig,
	)

	ConfigModule = fx.Options(
		configFactories,
	)
)
