// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/stashapp/stash/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// MovieReaderWriter is an autogenerated mock type for the MovieReaderWriter type
type MovieReaderWriter struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *MovieReaderWriter) All(ctx context.Context) ([]*models.Movie, error) {
	ret := _m.Called(ctx)

	var r0 []*models.Movie
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Movie); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Movie)
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

// Count provides a mock function with given fields: ctx
func (_m *MovieReaderWriter) Count(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountByPerformerID provides a mock function with given fields: ctx, performerID
func (_m *MovieReaderWriter) CountByPerformerID(ctx context.Context, performerID int) (int, error) {
	ret := _m.Called(ctx, performerID)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, int) int); ok {
		r0 = rf(ctx, performerID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, performerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountByStudioID provides a mock function with given fields: ctx, studioID
func (_m *MovieReaderWriter) CountByStudioID(ctx context.Context, studioID int) (int, error) {
	ret := _m.Called(ctx, studioID)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, int) int); ok {
		r0 = rf(ctx, studioID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, studioID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, newMovie
func (_m *MovieReaderWriter) Create(ctx context.Context, newMovie *models.Movie) error {
	ret := _m.Called(ctx, newMovie)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Movie) error); ok {
		r0 = rf(ctx, newMovie)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Destroy provides a mock function with given fields: ctx, id
func (_m *MovieReaderWriter) Destroy(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, id
func (_m *MovieReaderWriter) Find(ctx context.Context, id int) (*models.Movie, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Movie
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Movie); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: ctx, name, nocase
func (_m *MovieReaderWriter) FindByName(ctx context.Context, name string, nocase bool) (*models.Movie, error) {
	ret := _m.Called(ctx, name, nocase)

	var r0 *models.Movie
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *models.Movie); ok {
		r0 = rf(ctx, name, nocase)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, name, nocase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByNames provides a mock function with given fields: ctx, names, nocase
func (_m *MovieReaderWriter) FindByNames(ctx context.Context, names []string, nocase bool) ([]*models.Movie, error) {
	ret := _m.Called(ctx, names, nocase)

	var r0 []*models.Movie
	if rf, ok := ret.Get(0).(func(context.Context, []string, bool) []*models.Movie); ok {
		r0 = rf(ctx, names, nocase)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, bool) error); ok {
		r1 = rf(ctx, names, nocase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByPerformerID provides a mock function with given fields: ctx, performerID
func (_m *MovieReaderWriter) FindByPerformerID(ctx context.Context, performerID int) ([]*models.Movie, error) {
	ret := _m.Called(ctx, performerID)

	var r0 []*models.Movie
	if rf, ok := ret.Get(0).(func(context.Context, int) []*models.Movie); ok {
		r0 = rf(ctx, performerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, performerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByStudioID provides a mock function with given fields: ctx, studioID
func (_m *MovieReaderWriter) FindByStudioID(ctx context.Context, studioID int) ([]*models.Movie, error) {
	ret := _m.Called(ctx, studioID)

	var r0 []*models.Movie
	if rf, ok := ret.Get(0).(func(context.Context, int) []*models.Movie); ok {
		r0 = rf(ctx, studioID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, studioID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMany provides a mock function with given fields: ctx, ids
func (_m *MovieReaderWriter) FindMany(ctx context.Context, ids []int) ([]*models.Movie, error) {
	ret := _m.Called(ctx, ids)

	var r0 []*models.Movie
	if rf, ok := ret.Get(0).(func(context.Context, []int) []*models.Movie); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []int) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackImage provides a mock function with given fields: ctx, movieID
func (_m *MovieReaderWriter) GetBackImage(ctx context.Context, movieID int) ([]byte, error) {
	ret := _m.Called(ctx, movieID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, int) []byte); ok {
		r0 = rf(ctx, movieID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, movieID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFrontImage provides a mock function with given fields: ctx, movieID
func (_m *MovieReaderWriter) GetFrontImage(ctx context.Context, movieID int) ([]byte, error) {
	ret := _m.Called(ctx, movieID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, int) []byte); ok {
		r0 = rf(ctx, movieID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, movieID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetURLs provides a mock function with given fields: ctx, relatedID
func (_m *MovieReaderWriter) GetURLs(ctx context.Context, relatedID int) ([]string, error) {
	ret := _m.Called(ctx, relatedID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, int) []string); ok {
		r0 = rf(ctx, relatedID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, relatedID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasBackImage provides a mock function with given fields: ctx, movieID
func (_m *MovieReaderWriter) HasBackImage(ctx context.Context, movieID int) (bool, error) {
	ret := _m.Called(ctx, movieID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int) bool); ok {
		r0 = rf(ctx, movieID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, movieID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasFrontImage provides a mock function with given fields: ctx, movieID
func (_m *MovieReaderWriter) HasFrontImage(ctx context.Context, movieID int) (bool, error) {
	ret := _m.Called(ctx, movieID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int) bool); ok {
		r0 = rf(ctx, movieID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, movieID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: ctx, movieFilter, findFilter
func (_m *MovieReaderWriter) Query(ctx context.Context, movieFilter *models.MovieFilterType, findFilter *models.FindFilterType) ([]*models.Movie, int, error) {
	ret := _m.Called(ctx, movieFilter, findFilter)

	var r0 []*models.Movie
	if rf, ok := ret.Get(0).(func(context.Context, *models.MovieFilterType, *models.FindFilterType) []*models.Movie); ok {
		r0 = rf(ctx, movieFilter, findFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Movie)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, *models.MovieFilterType, *models.FindFilterType) int); ok {
		r1 = rf(ctx, movieFilter, findFilter)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *models.MovieFilterType, *models.FindFilterType) error); ok {
		r2 = rf(ctx, movieFilter, findFilter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// QueryCount provides a mock function with given fields: ctx, movieFilter, findFilter
func (_m *MovieReaderWriter) QueryCount(ctx context.Context, movieFilter *models.MovieFilterType, findFilter *models.FindFilterType) (int, error) {
	ret := _m.Called(ctx, movieFilter, findFilter)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *models.MovieFilterType, *models.FindFilterType) int); ok {
		r0 = rf(ctx, movieFilter, findFilter)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.MovieFilterType, *models.FindFilterType) error); ok {
		r1 = rf(ctx, movieFilter, findFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, updatedMovie
func (_m *MovieReaderWriter) Update(ctx context.Context, updatedMovie *models.Movie) error {
	ret := _m.Called(ctx, updatedMovie)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Movie) error); ok {
		r0 = rf(ctx, updatedMovie)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBackImage provides a mock function with given fields: ctx, movieID, backImage
func (_m *MovieReaderWriter) UpdateBackImage(ctx context.Context, movieID int, backImage []byte) error {
	ret := _m.Called(ctx, movieID, backImage)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []byte) error); ok {
		r0 = rf(ctx, movieID, backImage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFrontImage provides a mock function with given fields: ctx, movieID, frontImage
func (_m *MovieReaderWriter) UpdateFrontImage(ctx context.Context, movieID int, frontImage []byte) error {
	ret := _m.Called(ctx, movieID, frontImage)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []byte) error); ok {
		r0 = rf(ctx, movieID, frontImage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePartial provides a mock function with given fields: ctx, id, updatedMovie
func (_m *MovieReaderWriter) UpdatePartial(ctx context.Context, id int, updatedMovie models.MoviePartial) (*models.Movie, error) {
	ret := _m.Called(ctx, id, updatedMovie)

	var r0 *models.Movie
	if rf, ok := ret.Get(0).(func(context.Context, int, models.MoviePartial) *models.Movie); ok {
		r0 = rf(ctx, id, updatedMovie)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, models.MoviePartial) error); ok {
		r1 = rf(ctx, id, updatedMovie)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
