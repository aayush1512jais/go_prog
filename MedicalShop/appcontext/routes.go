package appcontext

import (
	"github.com/aayush1512jais/go_prog/MedicalShop/controller"
	"github.com/gorilla/mux"
)

func SetupRoutes(medicineController *controller.Controller) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/MedicalShop/status", medicineController.ServerStatus).Methods("GET")
	router.HandleFunc("/MedicalShop/getAll/{id}", medicineController.Get).Methods("GET")
	router.HandleFunc("/MedicalShop/getAll", medicineController.GetAll).Methods("GET")
	router.HandleFunc("/MedicalShop/add", medicineController.Add).Methods("POST")
	router.HandleFunc("/MedicalShop/update", medicineController.Update).Methods("PUT")
	router.HandleFunc("/MedicalShop/delete/{id}", medicineController.Delete).Methods("DELETE")
	return router
}
