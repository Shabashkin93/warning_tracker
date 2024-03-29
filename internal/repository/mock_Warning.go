// Code generated by mockery v2.39.1. DO NOT EDIT.

package repository

import (
	warning "github.com/Shabashkin93/warning_tracker/internal/domain/warning"
	mock "github.com/stretchr/testify/mock"
)

// MockWarning is an autogenerated mock type for the Warning type
type MockWarning struct {
	mock.Mock
}

type MockWarning_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWarning) EXPECT() *MockWarning_Expecter {
	return &MockWarning_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: in
func (_m *MockWarning) Create(in *warning.WarningCreate) (string, error) {
	ret := _m.Called(in)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*warning.WarningCreate) (string, error)); ok {
		return rf(in)
	}
	if rf, ok := ret.Get(0).(func(*warning.WarningCreate) string); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*warning.WarningCreate) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWarning_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockWarning_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - in *warning.WarningCreate
func (_e *MockWarning_Expecter) Create(in interface{}) *MockWarning_Create_Call {
	return &MockWarning_Create_Call{Call: _e.mock.On("Create", in)}
}

func (_c *MockWarning_Create_Call) Run(run func(in *warning.WarningCreate)) *MockWarning_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*warning.WarningCreate))
	})
	return _c
}

func (_c *MockWarning_Create_Call) Return(id string, err error) *MockWarning_Create_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockWarning_Create_Call) RunAndReturn(run func(*warning.WarningCreate) (string, error)) *MockWarning_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetOne provides a mock function with given fields: out
func (_m *MockWarning) GetOne(out *warning.WarningCreate) error {
	ret := _m.Called(out)

	if len(ret) == 0 {
		panic("no return value specified for GetOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*warning.WarningCreate) error); ok {
		r0 = rf(out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWarning_GetOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOne'
type MockWarning_GetOne_Call struct {
	*mock.Call
}

// GetOne is a helper method to define mock.On call
//   - out *warning.WarningCreate
func (_e *MockWarning_Expecter) GetOne(out interface{}) *MockWarning_GetOne_Call {
	return &MockWarning_GetOne_Call{Call: _e.mock.On("GetOne", out)}
}

func (_c *MockWarning_GetOne_Call) Run(run func(out *warning.WarningCreate)) *MockWarning_GetOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*warning.WarningCreate))
	})
	return _c
}

func (_c *MockWarning_GetOne_Call) Return(err error) *MockWarning_GetOne_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockWarning_GetOne_Call) RunAndReturn(run func(*warning.WarningCreate) error) *MockWarning_GetOne_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWarning creates a new instance of MockWarning. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWarning(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWarning {
	mock := &MockWarning{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
