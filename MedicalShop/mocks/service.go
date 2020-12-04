package mocks

import (
	apperrors "github.com/aayush1512jais/go_prog/MedicalShop/apperrors"
	mock "github.com/stretchr/testify/mock"

	model "github.com/aayush1512jais/go_prog/MedicalShop/model"
)

type MockService struct {
	mock.Mock
}

func NewMockService() *MockService {
	return &MockService{}
}
func (mock *MockService) Add(medicine model.Medicine) (int, apperrors.ErrorModel) {
	args := mock.Called(medicine)

	var result int
	if rf, ok := args.Get(0).(func(model.Medicine) int); ok {
		result = rf(medicine)
	} else {
		result = args.Get(0).(int)
	}

	var result1 apperrors.ErrorModel
	if rf, ok := args.Get(1).(func(model.Medicine) apperrors.ErrorModel); ok {
		result1 = rf(medicine)
	} else {
		result1 = args.Get(1).(apperrors.ErrorModel)
	}

	return result, result1
}

func (mock *MockService) Delete(id int) (bool, apperrors.ErrorModel) {
	args := mock.Called(id)

	var result bool
	if rf, ok := args.Get(0).(func(int) bool); ok {
		result = rf(id)
	} else {
		result = args.Get(0).(bool)
	}

	var result1 apperrors.ErrorModel
	if rf, ok := args.Get(1).(func(int) apperrors.ErrorModel); ok {
		result1 = rf(id)
	} else {
		result1 = args.Get(1).(apperrors.ErrorModel)
	}

	return result, result1
}

func (mock *MockService) Get(id int) (model.Medicine, apperrors.ErrorModel) {
	args := mock.Called(id)

	var result model.Medicine
	if rf, ok := args.Get(0).(func(int) model.Medicine); ok {
		result = rf(id)
	} else {
		result = args.Get(0).(model.Medicine)
	}

	var result1 apperrors.ErrorModel
	if rf, ok := args.Get(1).(func(int) apperrors.ErrorModel); ok {
		result1 = rf(id)
	} else {
		result1 = args.Get(1).(apperrors.ErrorModel)
	}

	return result, result1
}

func (mock *MockService) GetAll() ([]model.Medicine, apperrors.ErrorModel) {
	args := mock.Called()

	var result []model.Medicine
	if rf, ok := args.Get(0).(func() []model.Medicine); ok {
		result = rf()
	} else {
		if args.Get(0) != nil {
			result = args.Get(0).([]model.Medicine)
		}
	}

	var result1 apperrors.ErrorModel
	if rf, ok := args.Get(1).(func() apperrors.ErrorModel); ok {
		result1 = rf()
	} else {
		result1 = args.Get(1).(apperrors.ErrorModel)
	}

	return result, result1
}

func (mock *MockService) Update(medicine model.Medicine) (bool, apperrors.ErrorModel) {
	args := mock.Called(medicine)

	var result bool
	if rf, ok := args.Get(0).(func(model.Medicine) bool); ok {
		result = rf(medicine)
	} else {
		result = args.Get(0).(bool)
	}

	var result1 apperrors.ErrorModel
	if rf, ok := args.Get(1).(func(model.Medicine) apperrors.ErrorModel); ok {
		result1 = rf(medicine)
	} else {
		result1 = args.Get(1).(apperrors.ErrorModel)
	}

	return result, result1
}
