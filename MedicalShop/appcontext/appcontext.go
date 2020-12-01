package appcontext

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aayush1512jais/go_prog/MedicalShop/apperrors"

	"github.com/aayush1512jais/go_prog/MedicalShop/controller"
	"github.com/aayush1512jais/go_prog/MedicalShop/db"
	"github.com/aayush1512jais/go_prog/MedicalShop/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type AppContext struct {
	DatabaseURI string
	DB          *sql.DB
	Router      *mux.Router
}

func LoadApp() apperrors.ErrorModel { //(*AppContext, error) {

	//var context *AppContext
	// var myEnv map[string]string
	// var enverr error
	myEnv, enverr := godotenv.Read()
	if enverr != nil {
		error := apperrors.ErrorModel{
			Message: "Cannot read environment variables",
			Error:   enverr,
			Code:    502,
		}
		log.Fatal(error)
		return error
	}

	dbURI := myEnv["POSTGRES_DATABASE_URI"]

	database, connectionerr := db.ConnectToPostgres(dbURI)
	if connectionerr != (apperrors.ErrorModel{}) {

		log.Fatal(connectionerr)
		return connectionerr
	}
	//fmt.Println("database connection %v", status)

	//var err apperrors.ErrorHandler = apperrors.NewErrorHandler()
	repository := db.NewRepository(database)
	medicineService := service.NewMedicineService(*repository) //, err)
	medicineController := controller.NewMedicineController(*medicineService)

	router := SetupRoutes(medicineController)

	log.Fatal(http.ListenAndServe(":8000", router))

	//context = &AppContext{dbURI, database, router}
	return apperrors.ErrorModel{}

}
