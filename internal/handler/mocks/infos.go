// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "techradar-backend/internal/domain"
	dto "techradar-backend/internal/handler/dto"

	mock "github.com/stretchr/testify/mock"
)

// InfosUseCase is an autogenerated mock type for the InfosUseCase type
type InfosUseCase struct {
	mock.Mock
}

// GetListOfTeams provides a mock function with given fields:
func (_m *InfosUseCase) GetListOfTeams() ([]string, *dto.ErrorResponse) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 *dto.ErrorResponse
	if rf, ok := ret.Get(1).(func() *dto.ErrorResponse); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*dto.ErrorResponse)
		}
	}

	return r0, r1
}

// GetQuadrantsInfo provides a mock function with given fields:
func (_m *InfosUseCase) GetQuadrantsInfo() ([]domain.Quadrant, *dto.ErrorResponse) {
	ret := _m.Called()

	var r0 []domain.Quadrant
	if rf, ok := ret.Get(0).(func() []domain.Quadrant); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Quadrant)
		}
	}

	var r1 *dto.ErrorResponse
	if rf, ok := ret.Get(1).(func() *dto.ErrorResponse); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*dto.ErrorResponse)
		}
	}

	return r0, r1
}

// UpsertQuadrantsInfo provides a mock function with given fields: quadrant
func (_m *InfosUseCase) UpsertQuadrantsInfo(quadrant domain.Quadrant) *dto.ErrorResponse {
	ret := _m.Called(quadrant)

	var r0 *dto.ErrorResponse
	if rf, ok := ret.Get(0).(func(domain.Quadrant) *dto.ErrorResponse); ok {
		r0 = rf(quadrant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ErrorResponse)
		}
	}

	return r0
}

type mockConstructorTestingTNewInfosUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewInfosUseCase creates a new instance of InfosUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInfosUseCase(t mockConstructorTestingTNewInfosUseCase) *InfosUseCase {
	mock := &InfosUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}