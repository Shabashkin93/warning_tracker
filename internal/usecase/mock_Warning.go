// Code generated by mockery v2.39.1. DO NOT EDIT.

package usecase

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
func (_m *MockWarning) Create(in *warning.WarningCreate) (warning.WarningResponse, error) {
	ret := _m.Called(in)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 warning.WarningResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*warning.WarningCreate) (warning.WarningResponse, error)); ok {
		return rf(in)
	}
	if rf, ok := ret.Get(0).(func(*warning.WarningCreate) warning.WarningResponse); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(warning.WarningResponse)
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

func (_c *MockWarning_Create_Call) Return(result warning.WarningResponse, err error) *MockWarning_Create_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *MockWarning_Create_Call) RunAndReturn(run func(*warning.WarningCreate) (warning.WarningResponse, error)) *MockWarning_Create_Call {
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
