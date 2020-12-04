package service

import (
	"log"

	"github.com/aayush1512jais/go_prog/MedicalShop/apperrors"

	"github.com/aayush1512jais/go_prog/MedicalShop/db"
	"github.com/aayush1512jais/go_prog/MedicalShop/model"
)

type MedicineService interface {
	Add(medicine model.Medicine) error
	Update(medicine model.Medicine) error
	Delete(id int) error
	Get(id int) (model.Medicine, error)
	GetAll() ([]model.Medicine, error)
}

type Service struct {
	// medicines  []model.Medicine
	// repository *db.Repository
	//err        apperrors.ErrorHandler
}

var repository db.RepositoryInterface

func NewMedicineService(repo db.RepositoryInterface) *Service {
	repository = repo
	return &Service{}

	//err:        err,
	//	medicines: nil,
	//}
}

// func findMedicineByID(id int, medicines []model.Medicine) (model.Medicine, int) {
// 	for index, item := range medicines {
// 		if item.MedicineID == id {
// 			return item, index
// 		}
// 	}
// 	return model.Medicine{}, -1
// }

func (*Service) Add(medicine model.Medicine) error {
	//	medicines = append(medicines, medicine)
	err := repository.AddMedicine(medicine)
	if err != nil {
		log.Println("Service Add", err)

		return apperrors.ErrActionFailed

	}
	return nil
}

func (*Service) Update(medicine model.Medicine) error {
	// if _, index := findMedicineByID(medicine.MedicineID, medicines); index != -1 {
	// 	medicines = append(medicines[:index], medicines[index+1:]...)
	// 	medicines = append(medicines, medicine)
	//	return true
	//}
	//medicines = nil
	if err := repository.UpdateMedicine(medicine); err != nil {
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
func (*Service) Delete(id int) error {
	// if _, index := findMedicineByID(id, medicines); index != -1 {
	// 	medicines = append(medicines[:index], medicines[index+1:]...)
	// 	return true
	// }
	// return false
	//medicines = nil
	if err := repository.DeleteMedicine(id); err != nil {
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
func (*Service) Get(id int) (model.Medicine, error) {
	// if item, index := findMedicineByID(id, medicines); index != -1 {
	// 	return item, nil
	// } else {

	item, err := repository.GetMedicine(id)
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
	//	}

}
func (*Service) GetAll() ([]model.Medicine, error) {
	//medicines = nil
	rows, err := repository.GetAllMedicine()
	if err == nil {
		var meds []model.Medicine
		for rows.Next() {

			var med model.Medicine
			err = rows.Scan(&med.MedicineID, &med.Name, &med.Company, &med.IsExpired)
			if err != nil {
				log.Println("service GetAll ", err)

				return nil, apperrors.ErrDatabaseRecord
			}

			meds = append(meds, med)
		}

		return meds, nil
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
