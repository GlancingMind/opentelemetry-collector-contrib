// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/GlancingMind/opentelemetry-collector-contrib/receiver/bigipreceiver/internal/models"
)

// client is an autogenerated mock type for the client type
type MockClient struct {
	mock.Mock
}

// GetNewToken provides a mock function with given fields: ctx
func (_m *MockClient) GetNewToken(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNodes provides a mock function with given fields: ctx
func (_m *MockClient) GetNodes(ctx context.Context) (*models.Nodes, error) {
	ret := _m.Called(ctx)

	var r0 *models.Nodes
	if rf, ok := ret.Get(0).(func(context.Context) *models.Nodes); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Nodes)
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

// GetPoolMembers provides a mock function with given fields: ctx, pools
func (_m *MockClient) GetPoolMembers(ctx context.Context, pools *models.Pools) (*models.PoolMembers, error) {
	ret := _m.Called(ctx, pools)

	var r0 *models.PoolMembers
	if rf, ok := ret.Get(0).(func(context.Context, *models.Pools) *models.PoolMembers); ok {
		r0 = rf(ctx, pools)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PoolMembers)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Pools) error); ok {
		r1 = rf(ctx, pools)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPools provides a mock function with given fields: ctx
func (_m *MockClient) GetPools(ctx context.Context) (*models.Pools, error) {
	ret := _m.Called(ctx)

	var r0 *models.Pools
	if rf, ok := ret.Get(0).(func(context.Context) *models.Pools); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Pools)
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

// GetVirtualServers provides a mock function with given fields: ctx
func (_m *MockClient) GetVirtualServers(ctx context.Context) (*models.VirtualServers, error) {
	ret := _m.Called(ctx)

	var r0 *models.VirtualServers
	if rf, ok := ret.Get(0).(func(context.Context) *models.VirtualServers); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.VirtualServers)
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

// HasToken provides a mock function with given fields:
func (_m *MockClient) HasToken() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
