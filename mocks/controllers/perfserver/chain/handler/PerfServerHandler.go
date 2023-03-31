// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	v1 "github.com/epam/edp-perf-operator/v2/api/v1"
	mock "github.com/stretchr/testify/mock"
)

// PerfServerHandler is an autogenerated mock type for the PerfServerHandler type
type PerfServerHandler struct {
	mock.Mock
}

// ServeRequest provides a mock function with given fields: server
func (_m *PerfServerHandler) ServeRequest(server *v1.PerfServer) error {
	ret := _m.Called(server)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.PerfServer) error); ok {
		r0 = rf(server)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPerfServerHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewPerfServerHandler creates a new instance of PerfServerHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPerfServerHandler(t mockConstructorTestingTNewPerfServerHandler) *PerfServerHandler {
	mock := &PerfServerHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}