package mocks

import (
	"github.com/aayush1512jais/go_prog/MedicalShop/model"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) AddMedicine(medicine model.Medicine) (int, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(int), args.Error(1)
}
func (mock *MockRepository) GetMedicine(id int) (model.Medicine, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(model.Medicine), args.Error(1)
}
func (mock *MockRepository) GetAllMedicine() ([]model.Medicine, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Medicine), args.Error(1)
}
func (mock *MockRepository) UpdateMedicine(newMedicine model.Medicine) error {
	args := mock.Called()
	//result := args.Get(0)
	return args.Error(0)
}
func (mock *MockRepository) DeleteMedicine(id int) error {
	args := mock.Called()
	//result := args.Get(0)
	return args.Error(0)
}
