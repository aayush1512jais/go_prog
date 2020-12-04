package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	//"github.com/aayush1512jais/go_prog/MedicalShop/mocks"

	mocks "github.com/aayush1512jais/go_prog/MedicalShop/mocks/db"
	"github.com/aayush1512jais/go_prog/MedicalShop/model"
)

func TestService_Add(t *testing.T) {
	//mockRepo := mocks.NewMockRepository()
	mockRepo := new(mocks.RepositoryInterface)
	Medicine := model.Medicine{
		MedicineID: 1,
		Name:       "AYQ",
		Company:    "Cipla",
		IsExpired:  false,
	}
	mockRepo.On("AddMedicine").Return(nil)
	testService := NewMedicineService(mockRepo)
	err := testService.Add(Medicine)

	mockRepo.AssertExpectations(t)
	// var val int = 1
	// assert.Equal(t, val, res)
	assert.Nil(t, err)
}
