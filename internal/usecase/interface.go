package usecase

import (
	"techradar-backend/internal/domain"
)

type (
	TechnologyRepository interface {
		CreateTechnology(tech *domain.TechnologyDomain) error
		GetTechnologyByTeamAndTitle(name string, teamName string) ([]domain.TechnologyDomain, error)
		GetTechnologyByTeamAndQuadrants(team string, quadrant int) ([]domain.TechnologyDomain, error)
		GetTechnologyByTeam(team string) ([]domain.TechnologyDomain, error)
		UpdateTechnology(tech *domain.TechnologyDomain) error
	}

	InfoRepository interface {
		GetDistinctTeam() ([]string, error)
		GetQuadrantsInfo() ([]domain.Quadrant, error)
		UpsertQuadrantsInfo(quadrant domain.Quadrant) error
		GetQuadrantInfoById(id int) (*domain.Quadrant, error)
	}
)
