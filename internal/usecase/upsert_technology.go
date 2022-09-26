package usecase

import (
	"fmt"
	"log"
	"techradar-backend/internal/domain"
	"techradar-backend/internal/handler/dto"
	"techradar-backend/internal/misc"
)

type UpsertTechnologyUseCase struct {
	technologyRepository TechnologyRepository
	infoRepository       InfoRepository
}

func NewUpsertTechnology(technologyRepository TechnologyRepository, infoRepository InfoRepository) *UpsertTechnologyUseCase {
	return &UpsertTechnologyUseCase{
		technologyRepository: technologyRepository,
		infoRepository:       infoRepository,
	}
}

func (c *UpsertTechnologyUseCase) getTechnology(team string, friendlyTitle string) ([]domain.TechnologyDomain, *dto.ErrorResponse) {
	res, err := c.technologyRepository.GetTechnologyByTeamAndTitle(team, friendlyTitle)

	if err != nil {
		log.Println(err)
		return nil, dto.DefaultError()
	}

	return res, nil
}

func (c *UpsertTechnologyUseCase) CreateTechnology(request *domain.TechnologyDomain) *dto.ErrorResponse {

	quadrant, err := c.infoRepository.GetQuadrantInfoById(request.Quadrant)

	if err != nil {
		log.Println(err)
		return dto.DefaultError()
	}

	if quadrant == nil {
		errorResponse := dto.ErrorResponse{
			StatusCode: 400,
			Body: dto.ErrorBodyDTO{
				Messages: []string{fmt.Sprintf("Quadrant id %d not created", request.Quadrant)},
			},
		}
		return &errorResponse
	}

	res, errResp := c.getTechnology(request.Team, request.FriendlyTitle)

	if errResp != nil {

		log.Println(err)
		return errResp
	}

	if res != nil {
		errorResponse := dto.ErrorResponse{
			StatusCode: 409,
			Body: dto.ErrorBodyDTO{
				Messages: []string{fmt.Sprintf("Team %s already has the technology %s included", request.Team, request.FriendlyTitle)},
			},
		}
		return &errorResponse
	}

	if err := c.technologyRepository.CreateTechnology(request); err != nil {
		return dto.DefaultError()
	}

	return nil
}

func (c *UpsertTechnologyUseCase) UpdateTechnology(team string, friendlyTitle string, request *domain.UpdateTechnologyDomain) (*domain.TechnologyDomain, *dto.ErrorResponse) {
	techArr, errResp := c.getTechnology(team, friendlyTitle)

	if errResp != nil {

		log.Println(errResp)
		return nil, errResp
	}

	if techArr == nil {
		return nil, &dto.ErrorResponse{
			StatusCode: 404,
			Body: dto.ErrorBodyDTO{
				Messages: []string{fmt.Sprintf("Resources not found to team %s title %s", team, friendlyTitle)},
			},
		}
	}

	techObj := techArr[0]

	misc.BuildUpdateObject(&techObj, request)

	if err := c.technologyRepository.UpdateTechnology(&techObj); err != nil {

		log.Println(err)
		return nil, dto.DefaultError()
	}

	return &techObj, nil
}

func (c *UpsertTechnologyUseCase) MoveTechnology(team string, friendlyTitle string, request *domain.MoveTechnologyDomain) *dto.ErrorResponse {

	techArr, errResp := c.getTechnology(team, friendlyTitle)

	if errResp != nil {

		log.Println(errResp)
		return errResp
	}

	if techArr == nil {
		return &dto.ErrorResponse{
			StatusCode: 404,
			Body: dto.ErrorBodyDTO{
				Messages: []string{fmt.Sprintf("Resources not found to team %s title %s", team, friendlyTitle)},
			},
		}
	}

	techObj := techArr[0]

	misc.BuildHistoryObject(&techObj, request)

	if err := c.technologyRepository.UpdateTechnology(&techObj); err != nil {

		log.Println(err)
		return dto.DefaultError()
	}

	return nil
}
