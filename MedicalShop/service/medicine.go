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

type Service struct{}

var repository db.RepositoryInterface

func NewMedicineService(repo db.RepositoryInterface) *Service {
	repository = repo
	return &Service{}
}

func (*Service) Add(medicine model.Medicine) (int, error) {

	id, err := repository.AddMedicine(medicine)
	if err != nil {
		log.Println("Service Add", err)

		return -1, apperrors.ErrActionFailed

	}
	return id, nil
}

func (*Service) Update(medicine model.Medicine) error {
	if err := repository.UpdateMedicine(medicine); err != nil {
		log.Println("service Update", err)

		return apperrors.ErrActionFailed
	}
	return nil
}
func (*Service) Delete(id int) error {

	if err := repository.DeleteMedicine(id); err != nil {
		log.Println("service Delete ", err)

		return apperrors.ErrActionFailed
	}
	return nil
}

func (*Service) Get(id int) (model.Medicine, error) {

	item, err := repository.GetMedicine(id)
	if err != nil {
		log.Println("Service ", err)
		return model.Medicine{}, apperrors.ErrDataNotFound
	}
	return item, nil
}

func (*Service) GetAll() ([]model.Medicine, error) {

	medicines, err := repository.GetAllMedicine()
	if err == nil {
		return medicines, nil
	} else {
		log.Println("service GetAll ", err)
		return nil, apperrors.ErrActionFailed
	}
}
