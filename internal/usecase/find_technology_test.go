package usecase

import (
	"github.com/go-playground/assert/v2"
	"techradar-backend/internal/domain"
	"techradar-backend/internal/handler/dto"
	"techradar-backend/internal/usecase/mocks"
	"testing"
)

func TestGetTechnologyByTeam(t *testing.T) {
	defaultError := dto.DefaultError()
	mockTeam := "team1"

	technologyArr := []domain.TechnologyDomain{
		{
			Team:          "team1",
			Title:         "tech 1",
			FriendlyTitle: "tech_1",
			Description:   "Tech 1 description",
			Moved:         -1,
			Score:         1,
			Quadrant:      1,
			Active:        true,
			History:       nil,
		},
		{
			Team:          "team1",
			Title:         "tech 2",
			FriendlyTitle: "tech_2",
			Description:   "Tech 2 description",
			Moved:         -1,
			Score:         1,
			Quadrant:      2,
			Active:        true,
			History:       nil,
		},
	}

	t.Run("Get technology by team not found", func(t *testing.T) {

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockTechnologyRepository.On("GetTechnologyByTeam", mockTeam).Return(nil, nil)

		findTechnologyUseCase := NewFindTechnology(mockTechnologyRepository)

		resp, errResp := findTechnologyUseCase.GetTechnologyByTeam(mockTeam)

		assert.Equal(t, resp, nil)
		assert.Equal(t, errResp.StatusCode, 404)
	})

	t.Run("Get technology by team with error", func(t *testing.T) {

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockTechnologyRepository.On("GetTechnologyByTeam", mockTeam).Return(nil, defaultError)

		findTechnologyUseCase := NewFindTechnology(mockTechnologyRepository)

		resp, errResp := findTechnologyUseCase.GetTechnologyByTeam(mockTeam)

		assert.Equal(t, resp, nil)
		assert.Equal(t, errResp, defaultError)
	})

	t.Run("Get technology by team without error", func(t *testing.T) {

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockTechnologyRepository.On("GetTechnologyByTeam", mockTeam).Return(technologyArr, nil)

		findTechnologyUseCase := NewFindTechnology(mockTechnologyRepository)

		resp, errResp := findTechnologyUseCase.GetTechnologyByTeam(mockTeam)

		assert.Equal(t, resp, technologyArr)
		assert.Equal(t, errResp, nil)
	})
}

func TestGetTechnologyByTeamAndQuadrants(t *testing.T) {
	defaultError := dto.DefaultError()
	mockTeam := "team1"
	mockQuadrant := 0

	technologyArr := []domain.TechnologyDomain{
		{
			Team:          "team1",
			Title:         "tech 1",
			FriendlyTitle: "tech_1",
			Description:   "Tech 1 description",
			Moved:         -1,
			Score:         1,
			Quadrant:      1,
			Active:        true,
			History:       nil,
		},
		{
			Team:          "team1",
			Title:         "tech 2",
			FriendlyTitle: "tech_2",
			Description:   "Tech 2 description",
			Moved:         -1,
			Score:         1,
			Quadrant:      1,
			Active:        true,
			History:       nil,
		},
	}

	t.Run("Get technology by team and quadrant not found", func(t *testing.T) {

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockTechnologyRepository.On("GetTechnologyByTeamAndQuadrants", mockTeam, mockQuadrant).Return(nil, nil)

		findTechnologyUseCase := NewFindTechnology(mockTechnologyRepository)

		resp, errResp := findTechnologyUseCase.GetTechnologyByTeamAndQuadrants(mockTeam, mockQuadrant)

		assert.Equal(t, resp, nil)
		assert.Equal(t, errResp.StatusCode, 404)
	})

	t.Run("Get technology by team and quadrant with error", func(t *testing.T) {

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockTechnologyRepository.On("GetTechnologyByTeamAndQuadrants", mockTeam, mockQuadrant).Return(nil, defaultError)

		findTechnologyUseCase := NewFindTechnology(mockTechnologyRepository)

		resp, errResp := findTechnologyUseCase.GetTechnologyByTeamAndQuadrants(mockTeam, mockQuadrant)

		assert.Equal(t, resp, nil)
		assert.Equal(t, errResp, defaultError)
	})

	t.Run("Get technology by team and quadrant without error", func(t *testing.T) {

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockTechnologyRepository.On("GetTechnologyByTeamAndQuadrants", mockTeam, mockQuadrant).Return(technologyArr, nil)

		findTechnologyUseCase := NewFindTechnology(mockTechnologyRepository)

		resp, errResp := findTechnologyUseCase.GetTechnologyByTeamAndQuadrants(mockTeam, mockQuadrant)

		assert.Equal(t, resp, technologyArr)
		assert.Equal(t, errResp, nil)
	})
}

func TestGetTechnologyByTeamAndTitle(t *testing.T) {
	defaultError := dto.DefaultError()
	mockTeam := "team1"
	mockFriendlyTitle := "four-key-metrics"
	technologyArr := []domain.TechnologyDomain{
		{
			Team:          "team1",
			Title:         "tech 1",
			FriendlyTitle: "tech_1",
			Description:   "Tech 1 description",
			Moved:         -1,
			Score:         1,
			Quadrant:      1,
			Active:        true,
			History:       nil,
		},
	}
	t.Run("Get technology by team and title not found", func(t *testing.T) {

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", mockTeam, mockFriendlyTitle).Return(nil, nil)

		findTechnologyUseCase := NewFindTechnology(mockTechnologyRepository)

		resp, errResp := findTechnologyUseCase.GetTechnologyByTeamAndTitle(mockTeam, mockFriendlyTitle)

		assert.Equal(t, resp, nil)
		assert.Equal(t, errResp.StatusCode, 404)
	})

	t.Run("Get technology by team and title with error", func(t *testing.T) {

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", mockTeam, mockFriendlyTitle).Return(nil, defaultError)

		findTechnologyUseCase := NewFindTechnology(mockTechnologyRepository)

		resp, errResp := findTechnologyUseCase.GetTechnologyByTeamAndTitle(mockTeam, mockFriendlyTitle)

		assert.Equal(t, resp, nil)
		assert.Equal(t, errResp, defaultError)
	})

	t.Run("Get technology by team and title with error", func(t *testing.T) {

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", mockTeam, mockFriendlyTitle).Return(technologyArr, nil)

		findTechnologyUseCase := NewFindTechnology(mockTechnologyRepository)

		resp, errResp := findTechnologyUseCase.GetTechnologyByTeamAndTitle(mockTeam, mockFriendlyTitle)

		assert.Equal(t, resp, technologyArr[0])
		assert.Equal(t, errResp, nil)
	})
}
