package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"techradar-backend/internal/handler/dto"
)

type InfosHandler struct {
	useCase InfosUseCase
}

func NewInfosHandler(useCase InfosUseCase) InfosHandler {
	return InfosHandler{
		useCase: useCase,
	}
}

func registerInfosRoutes(f InfosHandler, handler *Handler) {
	handler.Gin.GET("/infos/team", f.GetListOfTeams)
	handler.Gin.GET("/infos/quadrants", f.GetQuadrantsInfo)
	handler.Gin.POST("/infos/quadrants", f.UpsertQuadrantsInfo)
}

func (i *InfosHandler) GetListOfTeams(context *gin.Context) {

	res, _ := i.useCase.GetListOfTeams()

	context.JSON(200, res)
}

func (i *InfosHandler) GetQuadrantsInfo(context *gin.Context) {

	res, _ := i.useCase.GetQuadrantsInfo()

	context.JSON(200, res)
}

func (i *InfosHandler) UpsertQuadrantsInfo(context *gin.Context) {
	upsertRequest := dto.Quadrant{}

	if err := context.BindJSON(&upsertRequest); err != nil {
		context.AbortWithStatus(400)
		return
	}

	_ = i.useCase.UpsertQuadrantsInfo(upsertRequest.ToDomain())

	context.Status(204)
}

var ModuleInfos = fx.Options(fx.Invoke(registerInfosRoutes))
