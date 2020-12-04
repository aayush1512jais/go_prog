package mocks

import (
	model "github.com/aayush1512jais/go_prog/MedicalShop/model"
	mock "github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (mock *MockService) Add(medicine model.Medicine) (int, error) {
	args := mock.Called(medicine)

	var result int
	if rf, ok := args.Get(0).(func(model.Medicine) int); ok {
		result = rf(medicine)
	} else {
		result = args.Get(0).(int)
	}

	var result1 error
	if rf, ok := args.Get(1).(func(model.Medicine) error); ok {
		result1 = rf(medicine)
	} else {
		result1 = args.Error(1)
	}

	return result, result1
}

func (mock *MockService) Delete(id int) error {
	args := mock.Called(id)

	var result error
	if rf, ok := args.Get(0).(func(int) error); ok {
		result = rf(id)
	} else {
		result = args.Error(0)
	}

	return result
}

func (mock *MockService) Get(id int) (model.Medicine, error) {
	args := mock.Called(id)

	var result model.Medicine
	if rf, ok := args.Get(0).(func(int) model.Medicine); ok {
		result = rf(id)
	} else {
		result = args.Get(0).(model.Medicine)
	}

	var result1 error
	if rf, ok := args.Get(1).(func(int) error); ok {
		result1 = rf(id)
	} else {
		result1 = args.Error(1)
	}

	return result, result1
}

func (mock *MockService) GetAll() ([]model.Medicine, error) {
	args := mock.Called()

	var result []model.Medicine
	if rf, ok := args.Get(0).(func() []model.Medicine); ok {
		result = rf()
	} else {
		if args.Get(0) != nil {
			result = args.Get(0).([]model.Medicine)
		}
	}

	var result1 error
	if rf, ok := args.Get(1).(func() error); ok {
		result1 = rf()
	} else {
		result1 = args.Error(1)
	}

	return result, result1
}

func (mock *MockService) Update(medicine model.Medicine) error {
	args := mock.Called(medicine)

	var result error
	if rf, ok := args.Get(0).(func(model.Medicine) error); ok {
		result = rf(medicine)
	} else {
		result = args.Error(0)
	}

	return result
}
