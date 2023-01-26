// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "vtrial/pkg/Models"

	mock "github.com/stretchr/testify/mock"
)

// ISearchService is an autogenerated mock type for the ISearchService type
type ISearchService struct {
	mock.Mock
}

// ComputeResult provides a mock function with given fields: _a0, _a1
func (_m *ISearchService) ComputeResult(_a0 context.Context, _a1 models.Keywords) ([]string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, models.Keywords) []string); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Keywords) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SavePage provides a mock function with given fields: _a0, _a1
func (_m *ISearchService) SavePage(_a0 context.Context, _a1 models.Page) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Page) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewISearchService interface {
	mock.TestingT
	Cleanup(func())
}

// NewISearchService creates a new instance of ISearchService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewISearchService(t mockConstructorTestingTNewISearchService) *ISearchService {
	mock := &ISearchService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
