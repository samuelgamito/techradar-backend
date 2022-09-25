package modules

import (
	"go.uber.org/fx"
	"techradar-backend/internal/config"
)

var (
	configFactories = fx.Provide(
		config.CreateConfig,
	)

	ConfigModule = fx.Options(
		configFactories,
	)
)
