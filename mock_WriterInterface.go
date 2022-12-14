// Code generated by mockery v2.14.1. DO NOT EDIT.

package events

import (
	context "context"

	kafka "github.com/segmentio/kafka-go"
	mock "github.com/stretchr/testify/mock"
)

// MockWriterInterface is an autogenerated mock type for the WriterInterface type
type MockWriterInterface struct {
	mock.Mock
}

// WriteMessages provides a mock function with given fields: ctx, msgs
func (_m *MockWriterInterface) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	_va := make([]interface{}, len(msgs))
	for _i := range msgs {
		_va[_i] = msgs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...kafka.Message) error); ok {
		r0 = rf(ctx, msgs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockWriterInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockWriterInterface creates a new instance of MockWriterInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockWriterInterface(t mockConstructorTestingTNewMockWriterInterface) *MockWriterInterface {
	mock := &MockWriterInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
