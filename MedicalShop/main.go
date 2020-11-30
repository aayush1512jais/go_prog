package main

import (
	"log"

	"github.com/aayush1512jais/go_prog/MedicalShop/appcontext"
	"github.com/aayush1512jais/go_prog/MedicalShop/apperrors"
)

func main() {

	err := appcontext.LoadApp()
	if err != (apperrors.ErrorModel{}) {
		log.Fatal(err)
	}
	//fmt.Println(*context)
	// var connect db.DBinterface = db.NewConnectionURI(*context)
	// status := connect.ConnectToPostgres()
	// fmt.Println(status)

	//router := mux.NewRouter()

	// log.Fatal(http.ListenAndServe(":8000", router))
}
