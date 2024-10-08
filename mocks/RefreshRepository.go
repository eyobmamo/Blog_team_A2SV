// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// RefreshRepository is an autogenerated mock type for the RefreshRepository type
type RefreshRepository struct {
	mock.Mock
}

// DeleteToken provides a mock function with given fields: ctx, userid
func (_m *RefreshRepository) DeleteToken(ctx context.Context, userid primitive.ObjectID) (error, int) {
	ret := _m.Called(ctx, userid)

	if len(ret) == 0 {
		panic("no return value specified for DeleteToken")
	}

	var r0 error
	var r1 int
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (error, int)); ok {
		return rf(ctx, userid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, userid)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) int); ok {
		r1 = rf(ctx, userid)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// FindToken provides a mock function with given fields: ctx, userid
func (_m *RefreshRepository) FindToken(ctx context.Context, userid primitive.ObjectID) (string, error, int) {
	ret := _m.Called(ctx, userid)

	if len(ret) == 0 {
		panic("no return value specified for FindToken")
	}

	var r0 string
	var r1 error
	var r2 int
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (string, error, int)); ok {
		return rf(ctx, userid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) string); ok {
		r0 = rf(ctx, userid)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, userid)
	} else {
		r1 = ret.Error(1)
	}

	if rf, ok := ret.Get(2).(func(context.Context, primitive.ObjectID) int); ok {
		r2 = rf(ctx, userid)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// StoreToken provides a mock function with given fields: ctx, userid, refreshToken
func (_m *RefreshRepository) StoreToken(ctx context.Context, userid primitive.ObjectID, refreshToken string) (error, int) {
	ret := _m.Called(ctx, userid, refreshToken)

	if len(ret) == 0 {
		panic("no return value specified for StoreToken")
	}

	var r0 error
	var r1 int
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, string) (error, int)); ok {
		return rf(ctx, userid, refreshToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, string) error); ok {
		r0 = rf(ctx, userid, refreshToken)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID, string) int); ok {
		r1 = rf(ctx, userid, refreshToken)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// UpdateToken provides a mock function with given fields: ctx, refreshToken, userid
func (_m *RefreshRepository) UpdateToken(ctx context.Context, refreshToken string, userid primitive.ObjectID) (error, int) {
	ret := _m.Called(ctx, refreshToken, userid)

	if len(ret) == 0 {
		panic("no return value specified for UpdateToken")
	}

	var r0 error
	var r1 int
	if rf, ok := ret.Get(0).(func(context.Context, string, primitive.ObjectID) (error, int)); ok {
		return rf(ctx, refreshToken, userid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, primitive.ObjectID) error); ok {
		r0 = rf(ctx, refreshToken, userid)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, primitive.ObjectID) int); ok {
		r1 = rf(ctx, refreshToken, userid)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// NewRefreshRepository creates a new instance of RefreshRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRefreshRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RefreshRepository {
	mock := &RefreshRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
