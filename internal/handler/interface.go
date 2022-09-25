package handler

import (
	"techradar-backend/internal/domain"
	"techradar-backend/internal/handler/dto"
)

type CreateTechnologyUseCase interface {
	CreateTechnology(request *domain.TechnologyDomain) (string, *dto.ErrorResponse)
}

type FindTechnologyUseCase interface {
	GetTechnologyByTeamAndTitle(name string, title string) ([]domain.TechnologyDomain, *dto.ErrorResponse)
	GetTechnologyByTeamAndQuadrants(team string, quadrant int) ([]domain.TechnologyDomain, *dto.ErrorResponse)
	GetTechnologyByTeam(team string) ([]domain.TechnologyDomain, *dto.ErrorResponse)
}

type InfosUseCase interface {
	GetListOfTeams() ([]string, *dto.ErrorResponse)
	GetQuadrantsInfo() ([]domain.Quadrant, *dto.ErrorResponse)
	UpsertQuadrantsInfo(quadrant domain.Quadrant) *dto.ErrorResponse
}
