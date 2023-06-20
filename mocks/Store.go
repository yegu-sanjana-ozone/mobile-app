// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/model"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// CreateMobile provides a mock function with given fields: mobile
func (_m *Store) CreateMobile(mobile Model.Mobile) error {
	ret := _m.Called(mobile)

	var r0 error
	if rf, ok := ret.Get(0).(func(Model.Mobile) error); ok {
		r0 = rf(mobile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: id
func (_m *Store) DeleteByID(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllMobiles provides a mock function with given fields:
func (_m *Store) GetAllMobiles() ([]Model.Mobile, error) {
	ret := _m.Called()

	var r0 []Model.Mobile
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]Model.Mobile, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []Model.Mobile); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Model.Mobile)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *Store) GetByID(id int) Model.Mobile {
	ret := _m.Called(id)

	var r0 Model.Mobile
	if rf, ok := ret.Get(0).(func(int) Model.Mobile); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(Model.Mobile)
	}

	return r0
}

// UpdateByID provides a mock function with given fields: id, brand
func (_m *Store) UpdateByID(id int, brand string) error {
	ret := _m.Called(id, brand)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(id, brand)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
