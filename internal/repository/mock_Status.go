// Code generated by mockery v2.39.1. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockStatus is an autogenerated mock type for the Status type
type MockStatus struct {
	mock.Mock
}

type MockStatus_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStatus) EXPECT() *MockStatus_Expecter {
	return &MockStatus_Expecter{mock: &_m.Mock}
}

// GetStatus provides a mock function with given fields:
func (_m *MockStatus) GetStatus() (bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStatus")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStatus_GetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatus'
type MockStatus_GetStatus_Call struct {
	*mock.Call
}

// GetStatus is a helper method to define mock.On call
func (_e *MockStatus_Expecter) GetStatus() *MockStatus_GetStatus_Call {
	return &MockStatus_GetStatus_Call{Call: _e.mock.On("GetStatus")}
}

func (_c *MockStatus_GetStatus_Call) Run(run func()) *MockStatus_GetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStatus_GetStatus_Call) Return(status bool, err error) *MockStatus_GetStatus_Call {
	_c.Call.Return(status, err)
	return _c
}

func (_c *MockStatus_GetStatus_Call) RunAndReturn(run func() (bool, error)) *MockStatus_GetStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStatus creates a new instance of MockStatus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStatus(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStatus {
	mock := &MockStatus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
