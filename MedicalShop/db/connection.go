package db

import (
	"database/sql"
	"log"

	"github.com/aayush1512jais/go_prog/MedicalShop/apperrors"
	_ "github.com/lib/pq"
)

func ConnectToPostgres(databaseURI string) (*sql.DB, apperrors.ErrorModel) {

	connStr := databaseURI //"postgres://postgres:aayush15@localhost/medicalShop?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	//log.Println("DB ", &db)
	erro := db.Ping()
	if erro != nil {
		log.Fatal(erro)
		error := apperrors.ErrorModel{
			Message: "Database connection dead",
			Error:   erro,
			Code:    502,
		}
		return nil, error
		//return true
	}
	if err != nil {
		log.Fatal(err)
		error := apperrors.ErrorModel{
			Message: "Cannot connect to postgresql",
			Error:   err,
			Code:    502,
		}
		return nil, error
		//return true
	}
	return db, apperrors.ErrorModel{}
	//return false

}
