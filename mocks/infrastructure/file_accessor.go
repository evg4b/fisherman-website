// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// FileAccessor is an autogenerated mock type for the FileAccessor type
type FileAccessor struct {
	mock.Mock
}

// Delete provides a mock function with given fields: path
func (_m *FileAccessor) Delete(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exist provides a mock function with given fields: path
func (_m *FileAccessor) Exist(path string) bool {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Read provides a mock function with given fields: path
func (_m *FileAccessor) Read(path string) (string, error) {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reader provides a mock function with given fields: path
func (_m *FileAccessor) Reader(path string) (io.Reader, error) {
	ret := _m.Called(path)

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func(string) io.Reader); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Write provides a mock function with given fields: path, content
func (_m *FileAccessor) Write(path string, content string) error {
	ret := _m.Called(path, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(path, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
