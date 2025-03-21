// Code generated by mockery v2.53.2. DO NOT EDIT.

package mocks

import (
	logic "ais_service/internal/logic"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// PublisherLogic is an autogenerated mock type for the PublisherLogic type
type PublisherLogic struct {
	mock.Mock
}

// PublishAisAccount provides a mock function with given fields: ctx, params
func (_m *PublisherLogic) PublishAisAccount(ctx context.Context, params logic.PublishAisAccountParams) (logic.PublishAisAccountOutput, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for PublishAisAccount")
	}

	var r0 logic.PublishAisAccountOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, logic.PublishAisAccountParams) (logic.PublishAisAccountOutput, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, logic.PublishAisAccountParams) logic.PublishAisAccountOutput); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(logic.PublishAisAccountOutput)
	}

	if rf, ok := ret.Get(1).(func(context.Context, logic.PublishAisAccountParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPublisherLogic creates a new instance of PublisherLogic. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPublisherLogic(t interface {
	mock.TestingT
	Cleanup(func())
}) *PublisherLogic {
	mock := &PublisherLogic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
