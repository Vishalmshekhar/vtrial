// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// I_Controller is an autogenerated mock type for the I_Controller type
type I_Controller struct {
	mock.Mock
}

// Routes provides a mock function with given fields:
func (_m *I_Controller) Routes() {
	_m.Called()
}

// UseComputeResult provides a mock function with given fields:
func (_m *I_Controller) UseComputeResult() {
	_m.Called()
}

// UseSavePage provides a mock function with given fields:
func (_m *I_Controller) UseSavePage() {
	_m.Called()
}

// healthCheck provides a mock function with given fields:
func (_m *I_Controller) healthCheck() {
	_m.Called()
}

type mockConstructorTestingTNewI_Controller interface {
	mock.TestingT
	Cleanup(func())
}

// NewI_Controller creates a new instance of I_Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewI_Controller(t mockConstructorTestingTNewI_Controller) *I_Controller {
	mock := &I_Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
