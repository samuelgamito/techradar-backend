package usecase

import (
	"github.com/go-playground/assert/v2"
	"techradar-backend/internal/domain"
	"techradar-backend/internal/handler/dto"
	"techradar-backend/internal/usecase/mocks"
	"testing"
)

func TestGetListOfTeams(t *testing.T) {

	defaultError := dto.DefaultError()

	t.Run("Get List of teams without error", func(t *testing.T) {

		mockInfoRepository := &mocks.InfoRepository{}

		mockInfoRepository.On("GetDistinctTeam").Return([]string{"team 3", "team 2", "team 5", "team 1"}, nil)

		infosUseCase := NewInfoUseCase(mockInfoRepository)

		resp, err := infosUseCase.GetListOfTeams()

		assert.Equal(t, err, nil)
		assert.Equal(t, resp, []string{"team 1", "team 2", "team 3", "team 5"})
	})

	t.Run("Get List of teams with error", func(t *testing.T) {

		mockInfoRepository := &mocks.InfoRepository{}

		mockInfoRepository.On("GetDistinctTeam").Return(nil, dto.DefaultError())

		infosUseCase := NewInfoUseCase(mockInfoRepository)

		resp, err := infosUseCase.GetListOfTeams()

		assert.Equal(t, err, defaultError)
		assert.Equal(t, resp, nil)
	})
}

func TestGetQuadrantsInfo(t *testing.T) {

	defaultError := dto.DefaultError()

	t.Run("Get List of quadrants without error", func(t *testing.T) {

		mockInfoRepository := &mocks.InfoRepository{}
		mockQuadrantsList := []domain.Quadrant{
			{
				Id:    0,
				Title: "Technology",
			},
			{
				Id:    1,
				Title: "Tools",
			},
			{
				Id:    2,
				Title: "Languages",
			},
			{
				Id:    3,
				Title: "Frameworks",
			},
		}
		expectedQuadrantsList := []domain.Quadrant{
			{
				Id:    0,
				Title: "Technology",
			},
			{
				Id:    1,
				Title: "Tools",
			},
			{
				Id:    2,
				Title: "Languages",
			},
			{
				Id:    3,
				Title: "Frameworks",
			},
		}
		mockInfoRepository.On("GetQuadrantsInfo").Return(mockQuadrantsList, nil)

		infosUseCase := NewInfoUseCase(mockInfoRepository)

		resp, err := infosUseCase.GetQuadrantsInfo()

		assert.Equal(t, err, nil)
		assert.Equal(t, resp, expectedQuadrantsList)
	})

	t.Run("Get List of quadrants with error", func(t *testing.T) {

		mockInfoRepository := &mocks.InfoRepository{}

		mockInfoRepository.On("GetQuadrantsInfo").Return(nil, dto.DefaultError())

		infosUseCase := NewInfoUseCase(mockInfoRepository)

		resp, err := infosUseCase.GetQuadrantsInfo()

		assert.Equal(t, err, defaultError)
		assert.Equal(t, resp, nil)
	})
}

func TestUpsertQuadrantsInfo(t *testing.T) {

	defaultError := dto.DefaultError()

	t.Run("Get List of quadrants without error", func(t *testing.T) {

		mockInfoRepository := &mocks.InfoRepository{}
		mockQuadrant := domain.Quadrant{
			Id:    0,
			Title: "Technology",
		}
		mockInfoRepository.On("UpsertQuadrantsInfo", mockQuadrant).Return(nil)

		infosUseCase := NewInfoUseCase(mockInfoRepository)

		err := infosUseCase.UpsertQuadrantsInfo(mockQuadrant)

		assert.Equal(t, err, nil)
	})

	t.Run("Get List of quadrants with error", func(t *testing.T) {

		mockInfoRepository := &mocks.InfoRepository{}
		mockQuadrant := domain.Quadrant{
			Id:    0,
			Title: "Technology",
		}
		mockInfoRepository.On("UpsertQuadrantsInfo", mockQuadrant).Return(dto.DefaultError())

		infosUseCase := NewInfoUseCase(mockInfoRepository)

		err := infosUseCase.UpsertQuadrantsInfo(mockQuadrant)

		assert.Equal(t, err, defaultError)
	})
}
