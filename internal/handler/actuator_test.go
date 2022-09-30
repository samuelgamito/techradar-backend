package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

func TestRegisterActuatorRoutes(t *testing.T) {

	_, mockedGinEngine := gin.CreateTestContext(httptest.NewRecorder())
	mockedHandler := Handler{
		Gin: mockedGinEngine,
	}
	registerActuatorRoutes(&mockedHandler)

	req := httptest.NewRequest("GET", "/health", nil)

	w := httptest.NewRecorder()

	mockedGinEngine.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), "{\"message\":\"Health OK\"}")

}
