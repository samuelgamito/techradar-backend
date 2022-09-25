package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"techradar-backend/internal/handler/dto"
)

func registerMoveTechnologyRoutes(handler *Handler) {
	handler.Gin.POST("/technologies/:id", func(context *gin.Context) {
		moveRequest := dto.MoveTechnologyDTO{}

		if err := context.BindJSON(&moveRequest); err != nil {
			context.AbortWithStatus(400)
			return
		}

		var tech dto.TechnologyDTO
		tech.Active = true
		tech.ID = uuid.New().String()
		tech.Moved = 0
		tech.Score = moveRequest.Score

		context.JSON(200, tech)
	})
}

var ModuleMoveTechnology = fx.Options(fx.Invoke(registerMoveTechnologyRoutes))
