package appcontext

import (
	"log"

	"github.com/aayush1512jais/go_prog/MedicalShop/controller"
	"github.com/gorilla/mux"
)

func SetupRoutes(medicineController controller.MedicineController) *mux.Router {
	router := mux.NewRouter()
	log.Println("MedicineController ", &medicineController)
	log.Println("Router ", &router)
	router.HandleFunc("/MedicalShop/getAll/{id}", medicineController.Get).Methods("GET")
	router.HandleFunc("/MedicalShop/getAll", medicineController.GetAll).Methods("GET")
	router.HandleFunc("/MedicalShop/add", medicineController.Add).Methods("POST")
	router.HandleFunc("/MedicalShop/update", medicineController.Update).Methods("PUT")
	router.HandleFunc("/MedicalShop/delete/{id}", medicineController.Delete).Methods("DELETE")
	return router
}
