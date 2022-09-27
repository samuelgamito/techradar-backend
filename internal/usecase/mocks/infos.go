// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "techradar-backend/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// InfoRepository is an autogenerated mock type for the InfoRepository type
type InfoRepository struct {
	mock.Mock
}

// GetDistinctTeam provides a mock function with given fields:
func (_m *InfoRepository) GetDistinctTeam() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQuadrantInfoById provides a mock function with given fields: id
func (_m *InfoRepository) GetQuadrantInfoById(id int) (*domain.Quadrant, error) {
	ret := _m.Called(id)

	var r0 *domain.Quadrant
	if rf, ok := ret.Get(0).(func(int) *domain.Quadrant); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Quadrant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQuadrantsInfo provides a mock function with given fields:
func (_m *InfoRepository) GetQuadrantsInfo() ([]domain.Quadrant, error) {
	ret := _m.Called()

	var r0 []domain.Quadrant
	if rf, ok := ret.Get(0).(func() []domain.Quadrant); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Quadrant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertQuadrantsInfo provides a mock function with given fields: quadrant
func (_m *InfoRepository) UpsertQuadrantsInfo(quadrant domain.Quadrant) error {
	ret := _m.Called(quadrant)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Quadrant) error); ok {
		r0 = rf(quadrant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewInfoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewInfoRepository creates a new instance of InfoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInfoRepository(t mockConstructorTestingTNewInfoRepository) *InfoRepository {
	mock := &InfoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
