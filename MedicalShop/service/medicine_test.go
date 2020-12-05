package service

import (
	"testing"

	"github.com/aayush1512jais/go_prog/MedicalShop/mocks"
	"github.com/aayush1512jais/go_prog/MedicalShop/model"
	"github.com/stretchr/testify/assert"
)

func TestService_Add(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	Medicine := model.Medicine{
		MedicineID: 1,
		Name:       "AYQ",
		Company:    "Cipla",
		IsExpired:  false,
	}
	mockRepo.On("AddMedicine").Return(Medicine.MedicineID, nil)
	testService := NewMedicineService(mockRepo)
	result, err := testService.Add(Medicine)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestService_Delete(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	id := 14
	mockRepo.On("DeleteMedicine").Return(nil)
	testService := NewMedicineService(mockRepo)
	err := testService.Delete(id)

	mockRepo.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestService_GetAll(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	medicine := model.Medicine{
		MedicineID: 1,
		Name:       "AYQ",
		Company:    "Cipla",
		IsExpired:  false,
	}
	mockRepo.On("GetAllMedicine").Return([]model.Medicine{medicine}, nil)
	testService := NewMedicineService(mockRepo)
	result, err := testService.GetAll()

	mockRepo.AssertExpectations(t)
	assert.Equal(t, "AYQ", result[0].Name)
	assert.Equal(t, "Cipla", result[0].Company)
	assert.Equal(t, false, result[0].IsExpired)
	//assert.NotNil(t, result)
	assert.Nil(t, err)

}
