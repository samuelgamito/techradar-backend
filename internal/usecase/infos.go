package usecase

import (
	"backend/internal/domain"
	"backend/internal/handler/dto"
	"log"
	"sort"
)

type InfosUseCase struct {
	repository InfoRepository
}

func NewInfoUseCase(repository InfoRepository) *InfosUseCase {
	return &InfosUseCase{
		repository: repository,
	}
}

func (i *InfosUseCase) GetListOfTeams() ([]string, *dto.ErrorResponse) {
	resp, err := i.repository.GetDistinctTeam()

	if err != nil {
		log.Println(err)
		return nil, dto.DefaultError()
	}

	sort.Strings(resp)
	return resp, nil
}

func (i *InfosUseCase) GetQuadrantsInfo() ([]domain.Quadrant, *dto.ErrorResponse) {
	resp, err := i.repository.GetQuadrantsInfo()

	if err != nil {
		log.Println(err)
		return nil, dto.DefaultError()
	}

	return resp, nil
}

func (i *InfosUseCase) UpsertQuadrantsInfo(quadrant domain.Quadrant) *dto.ErrorResponse {
	err := i.repository.UpsertQuadrantsInfo(quadrant)

	if err != nil {
		log.Println(err)
		return dto.DefaultError()
	}

	return nil
}
