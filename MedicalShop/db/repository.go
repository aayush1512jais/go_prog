package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aayush1512jais/go_prog/MedicalShop/model"
)

const (
	MedicineTableName    = "medicines"
	MedicineIdCol        = "medicine_id"
	MedicineNameCol      = "name"
	MedicineCompanyCol   = "company"
	MedicineIsExpiredCol = "is_expired"
)

type Repository struct{ db *sql.DB }

type RepositoryInterface interface {
	AddMedicine(medicine model.Medicine) error
	GetMedicine(id int) (model.Medicine, error)
	GetAllMedicine() (*sql.Rows, error)
	UpdateMedicine(newMedicine model.Medicine) error
	DeleteMedicine(id int) error
}

func NewRepository(database *sql.DB) RepositoryInterface {
	return &Repository{
		db: database,
	}
}
func (repo *Repository) AddMedicine(medicine model.Medicine) error {
	var id int
	err := repo.db.QueryRow(`INSERT INTO medicines(name, company,is_expired)
	VALUES($1,$2,$3) RETURNING medicine_id;`, medicine.Name, medicine.Company, medicine.IsExpired).Scan(&id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (repo *Repository) GetMedicine(id int) (model.Medicine, error) {
	row := repo.db.QueryRow(
		fmt.Sprintf(
			"SELECT * FROM %s WHERE %s=$1;",
			MedicineTableName,
			MedicineIdCol,
		),
		id,
	)
	var med model.Medicine
	if err := row.Scan(&med.MedicineID, &med.Name, &med.Company, &med.IsExpired); err != nil {
		//	log.Fatal("Repository ", err)
		return model.Medicine{}, err
	}
	return med, nil
}

func (repo *Repository) GetAllMedicine() (*sql.Rows, error) {
	//rows,err :=
	return repo.db.Query(
		fmt.Sprintf(
			"SELECT * FROM %s",
			MedicineTableName,
		),
	)
}

func (repo *Repository) UpdateMedicine(newMedicine model.Medicine) error {
	_, err := repo.db.Exec(
		fmt.Sprintf(
			"UPDATE %s SET %s=$1,%s=$2,%s=$3 WHERE %s=$4;",
			MedicineTableName,
			MedicineNameCol,
			MedicineCompanyCol,
			MedicineIsExpiredCol,
			MedicineIdCol,
		),
		newMedicine.Name,
		newMedicine.Company,
		newMedicine.IsExpired,
		newMedicine.MedicineID,
	)
	return err
}

func (repo *Repository) DeleteMedicine(id int) error {
	_, err := repo.db.Exec(
		fmt.Sprintf(
			"DELETE FROM %s WHERE %s=$1;",
			MedicineTableName,
			MedicineIdCol,
		),
		id,
	)
	return err
}
