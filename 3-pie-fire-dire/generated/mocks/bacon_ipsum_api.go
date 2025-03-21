// Code generated by mockery v2.53.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// BaconIpsumAPI is an autogenerated mock type for the BaconIpsumAPI type
type BaconIpsumAPI struct {
	mock.Mock
}

// GetText provides a mock function with given fields: ctx
func (_m *BaconIpsumAPI) GetText(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetText")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBaconIpsumAPI creates a new instance of BaconIpsumAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBaconIpsumAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *BaconIpsumAPI {
	mock := &BaconIpsumAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
