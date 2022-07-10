// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/arielizuardi/golang-backend-blog/model"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateArticle provides a mock function with given fields: a
func (_m *Repository) CreateArticle(a *model.Article) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Article) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllArticle provides a mock function with given fields:
func (_m *Repository) GetAllArticle() ([]*model.Article, error) {
	ret := _m.Called()

	var r0 []*model.Article
	if rf, ok := ret.Get(0).(func() []*model.Article); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticleByID provides a mock function with given fields: id
func (_m *Repository) GetArticleByID(id int) (*model.Article, error) {
	ret := _m.Called(id)

	var r0 *model.Article
	if rf, ok := ret.Get(0).(func(int) *model.Article); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
