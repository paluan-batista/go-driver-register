// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	entities "go-driver-register/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// DriverRepository is an autogenerated mock type for the DriverRepository type
type DriverRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: driver
func (_m *DriverRepository) Create(driver *entities.Driver) error {
	ret := _m.Called(driver)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Driver) error); ok {
		r0 = rf(driver)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *DriverRepository) Delete(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *DriverRepository) GetAll() ([]entities.Driver, error) {
	ret := _m.Called()

	var r0 []entities.Driver
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entities.Driver, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entities.Driver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Driver)
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
func (_m *DriverRepository) GetByID(id uuid.UUID) (entities.Driver, error) {
	ret := _m.Called(id)

	var r0 entities.Driver
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (entities.Driver, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) entities.Driver); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.Driver)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: driver
func (_m *DriverRepository) Update(driver *entities.Driver) error {
	ret := _m.Called(driver)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Driver) error); ok {
		r0 = rf(driver)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDriverRepository creates a new instance of DriverRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDriverRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *DriverRepository {
	mock := &DriverRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}