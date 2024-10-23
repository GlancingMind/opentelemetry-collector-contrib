// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/GlancingMind/opentelemetry-collector-contrib/receiver/rabbitmqreceiver/internal/models"
)

// MockClient is an autogenerated mock type for the client type
type MockClient struct {
	mock.Mock
}

// GetQueues provides a mock function with given fields: ctx
func (_m *MockClient) GetQueues(ctx context.Context) ([]*models.Queue, error) {
	ret := _m.Called(ctx)

	var r0 []*models.Queue
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Queue); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Queue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
