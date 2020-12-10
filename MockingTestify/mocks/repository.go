package mocks

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Get(id int) (int, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(int), args.Error(1)
}
func (mock *MockRepository) Add(id int) error {
	args := mock.Called()
	//result := args.get(0)
	return args.Error(0)
}
