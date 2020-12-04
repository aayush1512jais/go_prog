package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aayush1512jais/go_prog/MedicalShop/mocks"

	"github.com/aayush1512jais/go_prog/MedicalShop/model"
)

func TestService_Add(t *testing.T) {
	//mockRepo := mocks.NewMockRepository()
	mockRepo := new(mocks.MockRepository)
	medicine := model.Medicine{
		MedicineID: 1,
		Name:       AYQ,
		Company:    Cipla,
		IsExpired:  false,
	}
	mockRepo.On("AddMedicine").Return(1, nil)
	testService := NewMedicineService(mockRepo)
	res, err := testService.Add(medicine)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, res)
	assert.Nil(t, err)
}
