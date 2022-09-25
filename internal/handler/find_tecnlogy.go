package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"strconv"
)

type FindTechnology struct {
	useCase FindTechnologyUseCase
}

func NewFindResource(service FindTechnologyUseCase) FindTechnology {
	return FindTechnology{
		useCase: service,
	}
}
func registerFindTechnologyRoutes(f FindTechnology, handler *Handler) {
	handler.Gin.GET("/team/:team/technologies", f.findTechnologyByTeamController)
	handler.Gin.GET("/team/:team/quadrants/:quadrant/technologies", f.findTechnologyByTeamAndQuadrantController)
	handler.Gin.GET("/team/:team/technologies/:technology_friendly_title", f.findTechnologyByTeamAndTitleController)
}

func (f *FindTechnology) findTechnologyByTeamController(context *gin.Context) {
	team := context.Param("team")

	b, err := f.useCase.GetTechnologyByTeam(team)

	if err != nil {
		context.AbortWithStatusJSON(err.StatusCode, err.Body)
		return
	}
	context.JSON(200, b)
}

func (f *FindTechnology) findTechnologyByTeamAndQuadrantController(context *gin.Context) {

	team := context.Param("team")
	quadrant, e := strconv.Atoi(context.Param("quadrant"))

	if e != nil {
		context.AbortWithStatus(400)
	}

	b, err := f.useCase.GetTechnologyByTeamAndQuadrants(team, quadrant)

	if err != nil {
		context.AbortWithStatusJSON(err.StatusCode, err.Body)
		return
	}
	context.JSON(200, b)
}

func (f *FindTechnology) findTechnologyByTeamAndTitleController(context *gin.Context) {

	team := context.Param("team")
	technologyFriendlyTitle := context.Param("technology_friendly_title")

	b, err := f.useCase.GetTechnologyByTeamAndTitle(team, technologyFriendlyTitle)

	if err != nil {
		context.AbortWithStatusJSON(err.StatusCode, err.Body)
		return
	}
	context.JSON(200, b)
}

var ModuleFindTechnology = fx.Options(fx.Invoke(registerFindTechnologyRoutes))
