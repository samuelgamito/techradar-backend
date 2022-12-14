// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "techradar-backend/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// TechnologyRepository is an autogenerated mock type for the TechnologyRepository type
type TechnologyRepository struct {
	mock.Mock
}

// CreateTechnology provides a mock function with given fields: tech
func (_m *TechnologyRepository) CreateTechnology(tech *domain.TechnologyDomain) error {
	ret := _m.Called(tech)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.TechnologyDomain) error); ok {
		r0 = rf(tech)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTechnologyByTeam provides a mock function with given fields: team
func (_m *TechnologyRepository) GetTechnologyByTeam(team string) ([]domain.TechnologyDomain, error) {
	ret := _m.Called(team)

	var r0 []domain.TechnologyDomain
	if rf, ok := ret.Get(0).(func(string) []domain.TechnologyDomain); ok {
		r0 = rf(team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.TechnologyDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(team)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTechnologyByTeamAndQuadrants provides a mock function with given fields: team, quadrant
func (_m *TechnologyRepository) GetTechnologyByTeamAndQuadrants(team string, quadrant int) ([]domain.TechnologyDomain, error) {
	ret := _m.Called(team, quadrant)

	var r0 []domain.TechnologyDomain
	if rf, ok := ret.Get(0).(func(string, int) []domain.TechnologyDomain); ok {
		r0 = rf(team, quadrant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.TechnologyDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(team, quadrant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTechnologyByTeamAndTitle provides a mock function with given fields: name, teamName
func (_m *TechnologyRepository) GetTechnologyByTeamAndTitle(name string, teamName string) ([]domain.TechnologyDomain, error) {
	ret := _m.Called(name, teamName)

	var r0 []domain.TechnologyDomain
	if rf, ok := ret.Get(0).(func(string, string) []domain.TechnologyDomain); ok {
		r0 = rf(name, teamName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.TechnologyDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, teamName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTechnology provides a mock function with given fields: tech
func (_m *TechnologyRepository) UpdateTechnology(tech *domain.TechnologyDomain) error {
	ret := _m.Called(tech)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.TechnologyDomain) error); ok {
		r0 = rf(tech)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTechnologyRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTechnologyRepository creates a new instance of TechnologyRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTechnologyRepository(t mockConstructorTestingTNewTechnologyRepository) *TechnologyRepository {
	mock := &TechnologyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
