// Code generated by mockery v2.53.2. DO NOT EDIT.

package mocks

import (
	producer "ais_service/internal/dataaccess/mq/producer"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AccountCreatedHandler is an autogenerated mock type for the AccountCreatedHandler type
type AccountCreatedHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, event
func (_m *AccountCreatedHandler) Handle(ctx context.Context, event producer.AccountEvent) error {
	ret := _m.Called(ctx, event)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, producer.AccountEvent) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAccountCreatedHandler creates a new instance of AccountCreatedHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountCreatedHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountCreatedHandler {
	mock := &AccountCreatedHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
