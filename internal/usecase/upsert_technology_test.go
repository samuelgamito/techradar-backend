package usecase

import (
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"techradar-backend/internal/domain"
	"techradar-backend/internal/handler/dto"
	"techradar-backend/internal/usecase/mocks"
	"testing"
)

func TestCreateTechnology(t *testing.T) {

	mockQuadrantId := 1
	defaultError := dto.DefaultError()

	t.Run("Create technology quadrant not found", func(t *testing.T) {

		mockRequest := domain.TechnologyDomain{
			Quadrant: 1,
		}
		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockInfoRepository.On("GetQuadrantInfoById", mockQuadrantId).Return(nil, nil)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		err := upsertTechnologyUseCase.CreateTechnology(&mockRequest)

		assert.Equal(t, err.StatusCode, 400)
		assert.Equal(t, err.Body.Messages[0], "Quadrant id 1 not created")
	})

	t.Run("Create technology get quadrant error", func(t *testing.T) {

		mockRequest := domain.TechnologyDomain{
			Quadrant: 1,
		}
		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockInfoRepository.On("GetQuadrantInfoById", mockQuadrantId).Return(nil, defaultError)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		err := upsertTechnologyUseCase.CreateTechnology(&mockRequest)

		assert.Equal(t, err, defaultError)
	})

	t.Run("Create technology get technology error", func(t *testing.T) {

		mockRequest := domain.TechnologyDomain{
			Quadrant:      1,
			Team:          "team1",
			FriendlyTitle: "tech-1",
		}
		mockQuadrant := domain.Quadrant{
			Id:    1,
			Title: "Language",
		}
		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockInfoRepository.On("GetQuadrantInfoById", mockQuadrantId).Return(&mockQuadrant, nil)
		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return(nil, defaultError)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		err := upsertTechnologyUseCase.CreateTechnology(&mockRequest)

		assert.Equal(t, err, defaultError)
	})

	t.Run("Create technology to existing pair team and friendly title", func(t *testing.T) {

		mockRequest := domain.TechnologyDomain{
			Quadrant:      1,
			Team:          "team1",
			FriendlyTitle: "tech-1",
		}
		mockQuadrant := domain.Quadrant{
			Id:    1,
			Title: "Language",
		}
		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockInfoRepository.On("GetQuadrantInfoById", mockQuadrantId).Return(&mockQuadrant, nil)
		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return([]domain.TechnologyDomain{mockRequest}, nil)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		err := upsertTechnologyUseCase.CreateTechnology(&mockRequest)

		assert.Equal(t, err.StatusCode, 409)
		assert.Equal(t, err.Body.Messages[0], "Team team1 already has the technology tech-1 included")
	})

	t.Run("Create technology with error", func(t *testing.T) {

		mockRequest := domain.TechnologyDomain{
			Quadrant:      1,
			Team:          "team1",
			FriendlyTitle: "tech-1",
		}
		mockQuadrant := domain.Quadrant{
			Id:    1,
			Title: "Language",
		}
		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockInfoRepository.On("GetQuadrantInfoById", mockQuadrantId).Return(&mockQuadrant, nil)
		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return(nil, nil)
		mockTechnologyRepository.On("CreateTechnology", &mockRequest).Return(defaultError)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		err := upsertTechnologyUseCase.CreateTechnology(&mockRequest)

		assert.Equal(t, err, defaultError)
	})

	t.Run("Create technology without error", func(t *testing.T) {

		mockRequest := domain.TechnologyDomain{
			Quadrant:      1,
			Team:          "team1",
			FriendlyTitle: "tech-1",
		}
		mockQuadrant := domain.Quadrant{
			Id:    1,
			Title: "Language",
		}
		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockInfoRepository.On("GetQuadrantInfoById", mockQuadrantId).Return(&mockQuadrant, nil)
		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return(nil, nil)
		mockTechnologyRepository.On("CreateTechnology", &mockRequest).Return(nil)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		err := upsertTechnologyUseCase.CreateTechnology(&mockRequest)

		assert.Equal(t, err, nil)
	})
}

func TestUpdateTechnology(t *testing.T) {

	mockQuadrantId := 1
	defaultError := dto.DefaultError()

	t.Run("Update technology error get technology", func(t *testing.T) {

		mockRequest := domain.UpdateTechnologyDomain{
			Quadrant: &mockQuadrantId,
		}

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return(nil, defaultError)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		resp, err := upsertTechnologyUseCase.UpdateTechnology("team1", "tech-1", &mockRequest)

		assert.Equal(t, resp, nil)
		assert.Equal(t, err, defaultError)
	})

	t.Run("Update technology not found", func(t *testing.T) {

		mockRequest := domain.UpdateTechnologyDomain{
			Quadrant: &mockQuadrantId,
		}

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return(nil, nil)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		resp, err := upsertTechnologyUseCase.UpdateTechnology("team1", "tech-1", &mockRequest)

		assert.Equal(t, resp, nil)
		assert.Equal(t, err.StatusCode, 404)
		assert.Equal(t, err.Body.Messages[0], "Resources not found to team team1 title tech-1")
	})

	t.Run("Update technology error during update", func(t *testing.T) {

		mockRequest := domain.UpdateTechnologyDomain{
			Quadrant: &mockQuadrantId,
		}

		mockTechnology := domain.TechnologyDomain{
			Quadrant:      1,
			Team:          "team1",
			FriendlyTitle: "tech-1",
		}

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return([]domain.TechnologyDomain{mockTechnology}, nil)
		mockTechnologyRepository.On("UpdateTechnology", &mockTechnology).Return(defaultError)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		resp, err := upsertTechnologyUseCase.UpdateTechnology("team1", "tech-1", &mockRequest)

		assert.Equal(t, resp, nil)
		assert.Equal(t, err, defaultError)
	})

	t.Run("Update technology", func(t *testing.T) {

		mockNewDescription := "new description"
		mockNewActive := false
		mockRequest := domain.UpdateTechnologyDomain{
			Quadrant:    &mockQuadrantId,
			Description: &mockNewDescription,
			Active:      &mockNewActive,
		}

		mockTechnology := domain.TechnologyDomain{
			Quadrant:      0,
			Team:          "team1",
			FriendlyTitle: "tech-1",
			Description:   "old description",
			Active:        true,
		}

		expectedTechnology := domain.TechnologyDomain{
			Quadrant:      1,
			Team:          "team1",
			FriendlyTitle: "tech-1",
			Description:   "new description",
			Active:        false,
		}

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return([]domain.TechnologyDomain{mockTechnology}, nil)
		mockTechnologyRepository.On("UpdateTechnology", &expectedTechnology).Return(nil)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		resp, err := upsertTechnologyUseCase.UpdateTechnology("team1", "tech-1", &mockRequest)

		assert.Equal(t, err, nil)
		assert.Equal(t, resp, expectedTechnology)
	})
}

func TestMoveTechnology(t *testing.T) {

	defaultError := dto.DefaultError()

	t.Run("Move technology error get technology", func(t *testing.T) {

		mockRequest := domain.MoveTechnologyDomain{}

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return(nil, defaultError)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		err := upsertTechnologyUseCase.MoveTechnology("team1", "tech-1", &mockRequest)

		assert.Equal(t, err, defaultError)
	})

	t.Run("Move technology not found", func(t *testing.T) {

		mockRequest := domain.MoveTechnologyDomain{}

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return(nil, nil)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		err := upsertTechnologyUseCase.MoveTechnology("team1", "tech-1", &mockRequest)

		assert.Equal(t, err.StatusCode, 404)
		assert.Equal(t, err.Body.Messages[0], "Resources not found to team team1 title tech-1")
	})

	t.Run("Move technology error during update", func(t *testing.T) {

		mockComments := "moved"
		mockRequest := domain.MoveTechnologyDomain{
			Comments: &mockComments,
		}

		mockTechnology := domain.TechnologyDomain{
			Quadrant:      1,
			Score:         1,
			Moved:         0,
			Team:          "team1",
			FriendlyTitle: "tech-1",
		}

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return([]domain.TechnologyDomain{mockTechnology}, nil)
		mockTechnologyRepository.On("UpdateTechnology", mock.Anything).Return(defaultError)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		err := upsertTechnologyUseCase.MoveTechnology("team1", "tech-1", &mockRequest)

		assert.Equal(t, err, defaultError)
	})

	t.Run("Move technology", func(t *testing.T) {

		mockComments := "moved"
		mockScore := 2
		mockRequest := domain.MoveTechnologyDomain{
			Comments: &mockComments,
			Score:    &mockScore,
		}

		mockTechnology := domain.TechnologyDomain{
			Quadrant:      1,
			Score:         1,
			Moved:         0,
			Team:          "team1",
			FriendlyTitle: "tech-1",
		}

		mockTechnologyRepository := &mocks.TechnologyRepository{}
		mockInfoRepository := &mocks.InfoRepository{}

		mockTechnologyRepository.On("GetTechnologyByTeamAndTitle", "team1", "tech-1").Return([]domain.TechnologyDomain{mockTechnology}, nil)
		mockTechnologyRepository.On("UpdateTechnology", mock.Anything).Return(nil)

		upsertTechnologyUseCase := NewUpsertTechnology(mockTechnologyRepository, mockInfoRepository)

		err := upsertTechnologyUseCase.MoveTechnology("team1", "tech-1", &mockRequest)

		movedTechnologyInterface := mockTechnologyRepository.Calls[1].Arguments[0]

		movedTechnology := movedTechnologyInterface.(*domain.TechnologyDomain)
		assert.Equal(t, movedTechnology.Score, mockScore)
		assert.Equal(t, movedTechnology.Moved, 1)
		assert.Equal(t, movedTechnology.History[0].PreviousScore, mockTechnology.Score)
		assert.Equal(t, movedTechnology.History[0].Score, mockScore)
		assert.Equal(t, movedTechnology.History[0].Comments, mockComments)
		assert.Equal(t, err, nil)
	})
}
