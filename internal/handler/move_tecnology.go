package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"techradar-backend/internal/handler/dto"
)

type MoveTechnology struct {
	useCase UpsertTechnologyUseCase
}

func NewMoveTechnology(service UpsertTechnologyUseCase) MoveTechnology {
	return MoveTechnology{
		useCase: service,
	}
}

func registerMoveTechnologyRoutes(mv MoveTechnology, handler *Handler) {
	handler.Gin.POST("/team/:team/technologies/:technology_friendly_title", mv.moveTechnologyController)
}

func (mv *MoveTechnology) moveTechnologyController(context *gin.Context) {

	team := context.Param("team")
	technologyFriendlyTitle := context.Param("technology_friendly_title")

	moveRequest := dto.MoveTechnologyDTO{}

	if err := context.BindJSON(&moveRequest); err != nil {
		context.AbortWithStatus(400)
		return
	}

	if err := moveRequest.IsValid(); err != nil {
		context.AbortWithStatusJSON(err.StatusCode, err.Body)
		return
	}

	if err := mv.useCase.MoveTechnology(team, technologyFriendlyTitle, moveRequest.ToDomain()); err != nil {
		context.AbortWithStatusJSON(err.StatusCode, err.Body)
		return
	}

	context.Status(204)
}

var ModuleMoveTechnology = fx.Options(fx.Invoke(registerMoveTechnologyRoutes))
