package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"techradar-backend/internal/handler/dto"
)

type UpdateTechnology struct {
	useCase UpsertTechnologyUseCase
}

func NewUpdateTechnology(useCase UpsertTechnologyUseCase) UpdateTechnology {
	return UpdateTechnology{
		useCase: useCase,
	}
}

func registerUpdateTechnologyRoutes(u UpdateTechnology, handler *Handler) {
	handler.Gin.PATCH("/team/:team/technologies/:technology_friendly_title", u.updateTechnologyController)
}

func (u *UpdateTechnology) updateTechnologyController(context *gin.Context) {
	updateRequest := dto.UpdateTechnologyDTO{}

	team := context.Param("team")
	technologyFriendlyTitle := context.Param("technology_friendly_title")

	if err := context.BindJSON(&updateRequest); err != nil {
		context.AbortWithStatus(400)
		return
	}

	if err := updateRequest.IsValid(); err != nil {
		context.AbortWithStatusJSON(err.StatusCode, err.Body)
		return
	}

	res, err := u.useCase.UpdateTechnology(team, technologyFriendlyTitle, updateRequest.ToDomain())

	if err != nil {
		context.AbortWithStatusJSON(err.StatusCode, err.Body)
		return
	}

	context.JSON(200, res)
}

var ModuleUpdateTechnology = fx.Options(fx.Invoke(registerUpdateTechnologyRoutes))
