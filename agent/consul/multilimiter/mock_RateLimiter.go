// Code generated by mockery v2.20.0. DO NOT EDIT.

package multilimiter

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRateLimiter is an autogenerated mock type for the RateLimiter type
type MockRateLimiter struct {
	mock.Mock
}

// Allow provides a mock function with given fields: entity
func (_m *MockRateLimiter) Allow(entity LimitedEntity) bool {
	ret := _m.Called(entity)

	var r0 bool
	if rf, ok := ret.Get(0).(func(LimitedEntity) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeleteConfig provides a mock function with given fields: prefix
func (_m *MockRateLimiter) DeleteConfig(prefix []byte) {
	_m.Called(prefix)
}

// Run provides a mock function with given fields: ctx
func (_m *MockRateLimiter) Run(ctx context.Context) {
	_m.Called(ctx)
}

// UpdateConfig provides a mock function with given fields: c, prefix
func (_m *MockRateLimiter) UpdateConfig(c LimiterConfig, prefix []byte) {
	_m.Called(c, prefix)
}

type mockConstructorTestingTNewMockRateLimiter interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRateLimiter creates a new instance of MockRateLimiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRateLimiter(t mockConstructorTestingTNewMockRateLimiter) *MockRateLimiter {
	mock := &MockRateLimiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
