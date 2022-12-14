package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"techradar-backend/internal/handler/dto"
)

type CreateTechnology struct {
	useCase UpsertTechnologyUseCase
}

func NewCreateResource(service UpsertTechnologyUseCase) CreateTechnology {
	return CreateTechnology{
		useCase: service,
	}
}
func registerCreateTechnologyRoutes(c CreateTechnology, handler *Handler) {
	handler.Gin.POST("/technologies", c.createTechnologyController)
}

func (c *CreateTechnology) createTechnologyController(context *gin.Context) {
	createRequest := dto.CreateTechnologyDTO{}

	if err := context.BindJSON(&createRequest); err != nil {
		context.AbortWithStatus(400)
		return
	}

	if err := createRequest.IsValid(); err != nil {
		context.AbortWithStatusJSON(err.StatusCode, err.Body)
		return
	}

	resource := createRequest.ToDomain()
	err := c.useCase.CreateTechnology(resource)
	if err != nil {
		context.AbortWithStatusJSON(err.StatusCode, err.Body)
		return
	}

	context.JSON(200, resource)
}

var ModuleCreateTechnology = fx.Options(fx.Invoke(registerCreateTechnologyRoutes))
