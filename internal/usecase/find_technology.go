package usecase

import (
	"fmt"
	"log"
	"strconv"
	"techradar-backend/internal/domain"
	"techradar-backend/internal/handler/dto"
)

type FindTechnologyUseCase struct {
	repository TechnologyRepository
}

func NewFindTechnology(repository TechnologyRepository) *FindTechnologyUseCase {
	return &FindTechnologyUseCase{
		repository: repository,
	}
}

func validateArrayContent(resp []domain.TechnologyDomain, filter ...string) *dto.ErrorResponse {

	if resp == nil {
		return &dto.ErrorResponse{
			StatusCode: 404,
			Body: dto.ErrorBodyDTO{
				Messages: []string{fmt.Sprint("Resources not found to:", filter)},
			},
		}
	}

	return nil
}

func (f *FindTechnologyUseCase) GetTechnologyByTeam(team string) ([]domain.TechnologyDomain, *dto.ErrorResponse) {
	res, err := f.repository.GetTechnologyByTeam(team)

	if err != nil {
		log.Println(err)
		return nil, dto.DefaultError()
	}

	return res, validateArrayContent(res, team)
}

func (f *FindTechnologyUseCase) GetTechnologyByTeamAndQuadrants(team string, quadrant int) ([]domain.TechnologyDomain, *dto.ErrorResponse) {
	res, err := f.repository.GetTechnologyByTeamAndQuadrants(team, quadrant)

	if err != nil {
		log.Println(err)
		return nil, dto.DefaultError()
	}

	return res, validateArrayContent(res, team, strconv.Itoa(quadrant))
}

func (f *FindTechnologyUseCase) GetTechnologyByTeamAndTitle(team string, friendlyTitle string) (*domain.TechnologyDomain, *dto.ErrorResponse) {
	res, err := f.repository.GetTechnologyByTeamAndTitle(team, friendlyTitle)

	if err != nil {
		log.Println(err)
		return nil, dto.DefaultError()
	}

	if res == nil {
		return nil, &dto.ErrorResponse{
			StatusCode: 404,
			Body: dto.ErrorBodyDTO{
				Messages: []string{fmt.Sprintf("Resources not found to team %s title %s", team, friendlyTitle)},
			},
		}
	}

	techObj := res[0]
	return &techObj, validateArrayContent(res, team)
}
