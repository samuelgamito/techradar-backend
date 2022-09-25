package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"techradar-backend/internal/handler/dto"
)

func registerUpdateTechnologyRoutes(handler *Handler) {

	handler.Gin.PATCH("/technologies/:id", func(context *gin.Context) {
		updateRequest := dto.UpdateTechnologyDTO{}

		if err := context.BindJSON(&updateRequest); err != nil {
			context.AbortWithStatus(400)
			return
		}

		var tech dto.TechnologyDTO
		tech.Active = true
		tech.ID = uuid.New().String()
		tech.Moved = 0
		tech.Quadrant = updateRequest.Quadrant
		tech.Title = updateRequest.Title
		tech.Description = updateRequest.Description

		context.JSON(200, tech)
	})
}

var ModuleUpdateTechnology = fx.Options(fx.Invoke(registerUpdateTechnologyRoutes))
