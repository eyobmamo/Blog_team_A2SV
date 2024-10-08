// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	Domain "blogapp/Domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// TagRepository is an autogenerated mock type for the TagRepository type
type TagRepository struct {
	mock.Mock
}

// CreateTag provides a mock function with given fields: ctx, tag
func (_m *TagRepository) CreateTag(ctx context.Context, tag *Domain.Tag) (error, int) {
	ret := _m.Called(ctx, tag)

	if len(ret) == 0 {
		panic("no return value specified for CreateTag")
	}

	var r0 error
	var r1 int
	if rf, ok := ret.Get(0).(func(context.Context, *Domain.Tag) (error, int)); ok {
		return rf(ctx, tag)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Domain.Tag) error); ok {
		r0 = rf(ctx, tag)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Domain.Tag) int); ok {
		r1 = rf(ctx, tag)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// DeleteTag provides a mock function with given fields: ctx, id
func (_m *TagRepository) DeleteTag(ctx context.Context, id primitive.ObjectID) (error, int) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTag")
	}

	var r0 error
	var r1 int
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (error, int)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) int); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// GetAllTags provides a mock function with given fields: ctx
func (_m *TagRepository) GetAllTags(ctx context.Context) ([]*Domain.Tag, error, int) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTags")
	}

	var r0 []*Domain.Tag
	var r1 error
	var r2 int
	if rf, ok := ret.Get(0).(func(context.Context) ([]*Domain.Tag, error, int)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*Domain.Tag); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Domain.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	if rf, ok := ret.Get(2).(func(context.Context) int); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// GetTagBySlug provides a mock function with given fields: ctx, slug
func (_m *TagRepository) GetTagBySlug(ctx context.Context, slug string) (*Domain.Tag, error, int) {
	ret := _m.Called(ctx, slug)

	if len(ret) == 0 {
		panic("no return value specified for GetTagBySlug")
	}

	var r0 *Domain.Tag
	var r1 error
	var r2 int
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Domain.Tag, error, int)); ok {
		return rf(ctx, slug)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Domain.Tag); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Domain.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) int); ok {
		r2 = rf(ctx, slug)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// NewTagRepository creates a new instance of TagRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTagRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TagRepository {
	mock := &TagRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
