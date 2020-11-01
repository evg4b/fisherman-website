// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Shell is an autogenerated mock type for the Shell type
type Shell struct {
	mock.Mock
}

// Exec provides a mock function with given fields: commands, env
func (_m *Shell) Exec(commands []string, env *map[string]string) (string, string, error) {
	ret := _m.Called(commands, env)

	var r0 string
	if rf, ok := ret.Get(0).(func([]string, *map[string]string) string); ok {
		r0 = rf(commands, env)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func([]string, *map[string]string) string); ok {
		r1 = rf(commands, env)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]string, *map[string]string) error); ok {
		r2 = rf(commands, env)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
