// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/stashapp/stash/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// GalleryReaderWriter is an autogenerated mock type for the GalleryReaderWriter type
type GalleryReaderWriter struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *GalleryReaderWriter) All(ctx context.Context) ([]*models.Gallery, error) {
	ret := _m.Called(ctx)

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Gallery); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
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
func (_m *GalleryReaderWriter) Count(ctx context.Context) (int, error) {
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

// Create provides a mock function with given fields: ctx, newGallery
func (_m *GalleryReaderWriter) Create(ctx context.Context, newGallery models.Gallery) (*models.Gallery, error) {
	ret := _m.Called(ctx, newGallery)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, models.Gallery) *models.Gallery); ok {
		r0 = rf(ctx, newGallery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Gallery) error); ok {
		r1 = rf(ctx, newGallery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destroy provides a mock function with given fields: ctx, id
func (_m *GalleryReaderWriter) Destroy(ctx context.Context, id int) error {
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
func (_m *GalleryReaderWriter) Find(ctx context.Context, id int) (*models.Gallery, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Gallery); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
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

// FindByChecksum provides a mock function with given fields: ctx, checksum
func (_m *GalleryReaderWriter) FindByChecksum(ctx context.Context, checksum string) (*models.Gallery, error) {
	ret := _m.Called(ctx, checksum)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Gallery); ok {
		r0 = rf(ctx, checksum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, checksum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByChecksums provides a mock function with given fields: ctx, checksums
func (_m *GalleryReaderWriter) FindByChecksums(ctx context.Context, checksums []string) ([]*models.Gallery, error) {
	ret := _m.Called(ctx, checksums)

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*models.Gallery); ok {
		r0 = rf(ctx, checksums)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, checksums)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByImageID provides a mock function with given fields: ctx, imageID
func (_m *GalleryReaderWriter) FindByImageID(ctx context.Context, imageID int) ([]*models.Gallery, error) {
	ret := _m.Called(ctx, imageID)

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, int) []*models.Gallery); ok {
		r0 = rf(ctx, imageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, imageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByPath provides a mock function with given fields: ctx, path
func (_m *GalleryReaderWriter) FindByPath(ctx context.Context, path string) (*models.Gallery, error) {
	ret := _m.Called(ctx, path)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Gallery); ok {
		r0 = rf(ctx, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBySceneID provides a mock function with given fields: ctx, sceneID
func (_m *GalleryReaderWriter) FindBySceneID(ctx context.Context, sceneID int) ([]*models.Gallery, error) {
	ret := _m.Called(ctx, sceneID)

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, int) []*models.Gallery); ok {
		r0 = rf(ctx, sceneID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, sceneID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMany provides a mock function with given fields: ctx, ids
func (_m *GalleryReaderWriter) FindMany(ctx context.Context, ids []int) ([]*models.Gallery, error) {
	ret := _m.Called(ctx, ids)

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, []int) []*models.Gallery); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
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

// GetImageIDs provides a mock function with given fields: ctx, galleryID
func (_m *GalleryReaderWriter) GetImageIDs(ctx context.Context, galleryID int) ([]int, error) {
	ret := _m.Called(ctx, galleryID)

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, int) []int); ok {
		r0 = rf(ctx, galleryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, galleryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPerformerIDs provides a mock function with given fields: ctx, galleryID
func (_m *GalleryReaderWriter) GetPerformerIDs(ctx context.Context, galleryID int) ([]int, error) {
	ret := _m.Called(ctx, galleryID)

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, int) []int); ok {
		r0 = rf(ctx, galleryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, galleryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSceneIDs provides a mock function with given fields: ctx, galleryID
func (_m *GalleryReaderWriter) GetSceneIDs(ctx context.Context, galleryID int) ([]int, error) {
	ret := _m.Called(ctx, galleryID)

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, int) []int); ok {
		r0 = rf(ctx, galleryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, galleryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagIDs provides a mock function with given fields: ctx, galleryID
func (_m *GalleryReaderWriter) GetTagIDs(ctx context.Context, galleryID int) ([]int, error) {
	ret := _m.Called(ctx, galleryID)

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, int) []int); ok {
		r0 = rf(ctx, galleryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, galleryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: ctx, galleryFilter, findFilter
func (_m *GalleryReaderWriter) Query(ctx context.Context, galleryFilter *models.GalleryFilterType, findFilter *models.FindFilterType) ([]*models.Gallery, int, error) {
	ret := _m.Called(ctx, galleryFilter, findFilter)

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, *models.GalleryFilterType, *models.FindFilterType) []*models.Gallery); ok {
		r0 = rf(ctx, galleryFilter, findFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, *models.GalleryFilterType, *models.FindFilterType) int); ok {
		r1 = rf(ctx, galleryFilter, findFilter)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *models.GalleryFilterType, *models.FindFilterType) error); ok {
		r2 = rf(ctx, galleryFilter, findFilter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// QueryCount provides a mock function with given fields: ctx, galleryFilter, findFilter
func (_m *GalleryReaderWriter) QueryCount(ctx context.Context, galleryFilter *models.GalleryFilterType, findFilter *models.FindFilterType) (int, error) {
	ret := _m.Called(ctx, galleryFilter, findFilter)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *models.GalleryFilterType, *models.FindFilterType) int); ok {
		r0 = rf(ctx, galleryFilter, findFilter)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.GalleryFilterType, *models.FindFilterType) error); ok {
		r1 = rf(ctx, galleryFilter, findFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, updatedGallery
func (_m *GalleryReaderWriter) Update(ctx context.Context, updatedGallery models.Gallery) (*models.Gallery, error) {
	ret := _m.Called(ctx, updatedGallery)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, models.Gallery) *models.Gallery); ok {
		r0 = rf(ctx, updatedGallery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Gallery) error); ok {
		r1 = rf(ctx, updatedGallery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFileModTime provides a mock function with given fields: ctx, id, modTime
func (_m *GalleryReaderWriter) UpdateFileModTime(ctx context.Context, id int, modTime models.NullSQLiteTimestamp) error {
	ret := _m.Called(ctx, id, modTime)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, models.NullSQLiteTimestamp) error); ok {
		r0 = rf(ctx, id, modTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateImages provides a mock function with given fields: ctx, galleryID, imageIDs
func (_m *GalleryReaderWriter) UpdateImages(ctx context.Context, galleryID int, imageIDs []int) error {
	ret := _m.Called(ctx, galleryID, imageIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []int) error); ok {
		r0 = rf(ctx, galleryID, imageIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePartial provides a mock function with given fields: ctx, updatedGallery
func (_m *GalleryReaderWriter) UpdatePartial(ctx context.Context, updatedGallery models.GalleryPartial) (*models.Gallery, error) {
	ret := _m.Called(ctx, updatedGallery)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(context.Context, models.GalleryPartial) *models.Gallery); ok {
		r0 = rf(ctx, updatedGallery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.GalleryPartial) error); ok {
		r1 = rf(ctx, updatedGallery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePerformers provides a mock function with given fields: ctx, galleryID, performerIDs
func (_m *GalleryReaderWriter) UpdatePerformers(ctx context.Context, galleryID int, performerIDs []int) error {
	ret := _m.Called(ctx, galleryID, performerIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []int) error); ok {
		r0 = rf(ctx, galleryID, performerIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateScenes provides a mock function with given fields: ctx, galleryID, sceneIDs
func (_m *GalleryReaderWriter) UpdateScenes(ctx context.Context, galleryID int, sceneIDs []int) error {
	ret := _m.Called(ctx, galleryID, sceneIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []int) error); ok {
		r0 = rf(ctx, galleryID, sceneIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTags provides a mock function with given fields: ctx, galleryID, tagIDs
func (_m *GalleryReaderWriter) UpdateTags(ctx context.Context, galleryID int, tagIDs []int) error {
	ret := _m.Called(ctx, galleryID, tagIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []int) error); ok {
		r0 = rf(ctx, galleryID, tagIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
