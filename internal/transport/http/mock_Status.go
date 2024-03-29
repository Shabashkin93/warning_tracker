// Code generated by mockery v2.39.1. DO NOT EDIT.

package transport

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

// Register provides a mock function with given fields: _a0, _a1
func (_m *MockStatus) Register(_a0 string, _a1 interface{}) {
	_m.Called(_a0, _a1)
}

// MockStatus_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockStatus_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - _a0 string
//   - _a1 interface{}
func (_e *MockStatus_Expecter) Register(_a0 interface{}, _a1 interface{}) *MockStatus_Register_Call {
	return &MockStatus_Register_Call{Call: _e.mock.On("Register", _a0, _a1)}
}

func (_c *MockStatus_Register_Call) Run(run func(_a0 string, _a1 interface{})) *MockStatus_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *MockStatus_Register_Call) Return() *MockStatus_Register_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStatus_Register_Call) RunAndReturn(run func(string, interface{})) *MockStatus_Register_Call {
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
