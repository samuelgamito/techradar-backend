package bootstrap

import (
	"backend/internal/handler"
	"context"
	"fmt"
	"go.uber.org/fx"
)

func RegisterHooks(lifecycle fx.Lifecycle, h *handler.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting application in :8080")
				go h.Gin.Run(":8080")
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
