package service

import (
	"log"

	"github.com/aayush1512jais/go_prog/MedicalShop/apperrors"

	"github.com/aayush1512jais/go_prog/MedicalShop/db"
	"github.com/aayush1512jais/go_prog/MedicalShop/model"
)

type MedicineService interface {
	Add(medicine model.Medicine) (int, apperrors.ErrorModel)
	Update(medicine model.Medicine) (bool, apperrors.ErrorModel)
	Delete(id int) (bool, apperrors.ErrorModel)
	Get(id int) (model.Medicine, apperrors.ErrorModel)
	GetAll() ([]model.Medicine, apperrors.ErrorModel)
}

type Service struct {
	medicines  []model.Medicine
	repository *db.Repository
	//err        apperrors.ErrorHandler
}

func NewMedicineService(repo *db.Repository) *Service {
	return &Service{
		repository: repo,
		//err:        err,
		medicines: nil,
	}
}

func findMedicineByID(id int, medicines []model.Medicine) (model.Medicine, int) {
	for index, item := range medicines {
		if item.MedicineID == id {
			return item, index
		}
	}
	return model.Medicine{}, -1
}

func (service *Service) Add(medicine model.Medicine) (int, error) {
	service.medicines = append(service.medicines, medicine)
	id, err := service.repository.AddMedicine(medicine)
	if err != nil {
		log.Println("Service Add", err)
		// error := apperrors.ErrorModel{
		// 	Message: "Failed to Add requsted medicine",
		// 	//Error:   err,
		// 	Code: 502,
		// }

		return -1, apperrors.ErrActionFailed

	}
	return id, nil
}

func (service *Service) Update(medicine model.Medicine) error {
	// if _, index := findMedicineByID(medicine.MedicineID, service.medicines); index != -1 {
	// 	service.medicines = append(service.medicines[:index], service.medicines[index+1:]...)
	// 	service.medicines = append(service.medicines, medicine)
	//	return true
	//}
	service.medicines = nil
	if err := service.repository.UpdateMedicine(medicine); err != nil {
		log.Println("service Update", err)
		// error := apperrors.ErrorModel{
		// 	Message: "Failed to Update requsted medicine",
		// 	//	Error:   err,
		// 	Code: 502,
		// }

		return apperrors.ErrActionFailed
	}
	return nil
}
func (service *Service) Delete(id int) error {
	// if _, index := findMedicineByID(id, service.medicines); index != -1 {
	// 	service.medicines = append(service.medicines[:index], service.medicines[index+1:]...)
	// 	return true
	// }
	// return false
	service.medicines = nil
	if err := service.repository.DeleteMedicine(id); err != nil {
		log.Println("service Delete ", err)
		// error := apperrors.ErrorModel{
		// 	Message: "Failed to Delete requsted medicine",
		// 	//Error:   err,
		// 	Code: 302,
		// }

		return apperrors.ErrActionFailed
	}
	return nil
}
func (service *Service) Get(id int) (model.Medicine, error) {
	if item, index := findMedicineByID(id, service.medicines); index != -1 {
		return item, nil
	} else {

		item, err := service.repository.GetMedicine(id)
		if err != nil {
			log.Println("Service ", err)
			// error := apperrors.ErrorModel{
			// 	Message: "Cannot Find requsted medicine",
			// 	Error:   err,
			// 	Code:    502,
			// }

			return model.Medicine{}, apperrors.ErrDataNotFound

		}

		return item, nil
	}

}
func (service *Service) GetAll() ([]model.Medicine, error) {
	service.medicines = nil
	rows, err := service.repository.GetAllMedicine()
	if err == nil {
		for rows.Next() {

			var med model.Medicine
			err = rows.Scan(&med.MedicineID, &med.Name, &med.Company, &med.IsExpired)
			if err != nil {
				log.Println("service GetAll ", err)

				return nil, apperrors.ErrDatabaseRecord
			}

			service.medicines = append(service.medicines, med)
		}

		return service.medicines, nil
	} else {
		log.Println("service GetAll ", err)
		// error := apperrors.ErrorModel{
		// 	Message: "Failed to get all medicines",
		// 	//Error:   err,
		// 	Code: 302,
		// }

		return nil, apperrors.ErrActionFailed
	}

}
