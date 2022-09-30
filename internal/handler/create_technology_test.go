package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"strings"
	"techradar-backend/internal/handler/dto"
	"techradar-backend/internal/handler/mocks"
	"testing"
)

func setUp() (*httptest.ResponseRecorder, *gin.Engine) {
	mockUpsertTechnologyUseCase := &mocks.UpsertTechnologyUseCase{}
	httpRecorder := httptest.NewRecorder()
	_, mockGinEngine := gin.CreateTestContext(httpRecorder)

	mockHandler := Handler{
		Gin: mockGinEngine,
	}
	createTechnology := NewCreateResource(mockUpsertTechnologyUseCase)

	registerCreateTechnologyRoutes(createTechnology, &mockHandler)

	return httpRecorder, mockGinEngine
}

func TestCreateTechnologyController(t *testing.T) {

	t.Run("Json Parsing Error", func(t *testing.T) {

		httpRecorder, mockGinEngine := setUp()

		req := httptest.NewRequest("POST", "/technologies", nil)

		mockGinEngine.ServeHTTP(httpRecorder, req)

		assert.Equal(t, httpRecorder.Code, 400)
		assert.Equal(t, httpRecorder.Body.String(), "")
	})

	t.Run("Json Not Valid Error", func(t *testing.T) {

		createRequestMock := dto.CreateTechnologyDTO{}

		expectedErrorBody := dto.ErrorBodyDTO{
			Messages: []string{
				"Field Title is required.",
				"Field Description is required.",
				"Field Team is required.",
				"Score min values is 0, actual value is 0.",
				"Field Quadrant is required.",
			},
		}

		httpRecorder, mockGinEngine := setUp()

		createRequestJsonMock, _ := json.Marshal(createRequestMock)
		req := httptest.NewRequest("POST", "/technologies", strings.NewReader(string(createRequestJsonMock)))

		mockGinEngine.ServeHTTP(httpRecorder, req)

		assert.Equal(t, httpRecorder.Code, 400)
		errorBodyResponse := dto.ErrorBodyDTO{}
		_ = json.Unmarshal(httpRecorder.Body.Bytes(), &errorBodyResponse)
		assert.Equal(t, errorBodyResponse, expectedErrorBody)
	})

}
