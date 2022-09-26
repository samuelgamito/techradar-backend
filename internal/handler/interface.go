package handler

import (
	"techradar-backend/internal/domain"
	"techradar-backend/internal/handler/dto"
)

type UpsertTechnologyUseCase interface {
	CreateTechnology(request *domain.TechnologyDomain) *dto.ErrorResponse
	UpdateTechnology(team string, friendlyTitle string, request *domain.UpdateTechnologyDomain) (*domain.TechnologyDomain, *dto.ErrorResponse)
	MoveTechnology(team string, friendlyTitle string, request *domain.MoveTechnologyDomain) *dto.ErrorResponse
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
