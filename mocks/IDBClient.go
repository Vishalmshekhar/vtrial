// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "vtrial/pkg/Models"

	mock "github.com/stretchr/testify/mock"
)

// IDBClient is an autogenerated mock type for the IDBClient type
type IDBClient struct {
	mock.Mock
}

// GetAllCollection provides a mock function with given fields:
func (_m *IDBClient) GetAllCollection() []models.Page {
	ret := _m.Called()

	var r0 []models.Page
	if rf, ok := ret.Get(0).(func() []models.Page); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Page)
		}
	}

	return r0
}

// InsertOnePage provides a mock function with given fields: _a0, _a1
func (_m *IDBClient) InsertOnePage(_a0 context.Context, _a1 models.Page) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Page) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIDBClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewIDBClient creates a new instance of IDBClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDBClient(t mockConstructorTestingTNewIDBClient) *IDBClient {
	mock := &IDBClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
