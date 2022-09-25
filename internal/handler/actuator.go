package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func registerActuatorRoutes(handler *Handler) {

	handler.Gin.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Health OK"})
	})
}

var ModuleActuator = fx.Options(fx.Invoke(registerActuatorRoutes))
