package usecase

import (
	"backend/internal/domain"
	"backend/internal/handler/dto"
	"fmt"
)

type CreateTechnologyUseCase struct {
	technologyRepository TechnologyRepository
	infoRepository       InfoRepository
}

func NewCreateTechnology(technologyRepository TechnologyRepository, infoRepository InfoRepository) *CreateTechnologyUseCase {
	return &CreateTechnologyUseCase{
		technologyRepository: technologyRepository,
		infoRepository:       infoRepository,
	}
}

func (c *CreateTechnologyUseCase) CreateTechnology(request *domain.TechnologyDomain) (string, *dto.ErrorResponse) {

	quadrant, err := c.infoRepository.GetQuadrantInfoById(request.Quadrant)

	if quadrant == nil {
		errorResponse := dto.ErrorResponse{
			StatusCode: 400,
			Body: dto.ErrorBodyDTO{
				Messages: []string{fmt.Sprintf("Quadrant id %d not created", request.Quadrant)},
			},
		}
		return "", &errorResponse
	}

	res, err := c.technologyRepository.GetTechnologyByTeamAndTitle(request.Team, request.FriendlyTitle)

	if err != nil {
		return "", dto.DefaultError()
	}

	if res != nil {
		errorResponse := dto.ErrorResponse{
			StatusCode: 409,
			Body: dto.ErrorBodyDTO{
				Messages: []string{fmt.Sprintf("Team %s already has the technology %s included", request.Team, request.Title)},
			},
		}
		return "", &errorResponse
	}

	if err != nil {
		return "", dto.DefaultError()
	}

	if err := c.technologyRepository.CreateTechnology(request); err != nil {
		return "", dto.DefaultError()
	}

	return request.FriendlyTitle, nil
}
