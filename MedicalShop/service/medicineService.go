package service

import (
	"github.com/aayush1512jais/go_prog/MedicalShop/model"
)

type MedicineService interface {
	Add(medicine model.Medicine) bool
	Update(medicine model.Medicine) bool
	Delete(id int) bool
	Get(id int) (model.Medicine, bool)
	GetAll() []model.Medicine
}

type medicineService struct {
	medicines []model.Medicine
}

func New() MedicineService {
	return &medicineService{}
}

func findMedicineByID(id int, medicines []model.Medicine) (model.Medicine, int) {
	for index, item := range medicines {
		if item.MedicineID == id {
			return item, index
		}
	}
	return model.Medicine{}, -1
}

func (service *medicineService) Add(medicine model.Medicine) bool {
	service.medicines = append(service.medicines, medicine)
	return true
}

func (service *medicineService) Update(medicine model.Medicine) bool {
	if _, index := findMedicineByID(medicine.MedicineID, service.medicines); index != -1 {
		service.medicines = append(service.medicines[:index], service.medicines[index+1:]...)
		service.medicines = append(service.medicines, medicine)
		return true
	}
	return false
}
func (service *medicineService) Delete(id int) bool {
	if _, index := findMedicineByID(id, service.medicines); index != -1 {
		service.medicines = append(service.medicines[:index], service.medicines[index+1:]...)
		return true
	}
	return false
}
func (service *medicineService) Get(id int) (model.Medicine, bool) {
	if item, index := findMedicineByID(id, service.medicines); index != -1 {
		return item, true
	}
	return model.Medicine{}, false
}
func (service *medicineService) GetAll() []model.Medicine {
	return service.medicines
}
