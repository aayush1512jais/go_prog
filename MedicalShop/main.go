package main

import (
	"log"
	"net/http"

	"github.com/aayush1512jais/go_prog/MedicalShop/controller"
	"github.com/aayush1512jais/go_prog/MedicalShop/service"
	"github.com/gorilla/mux"
)

var (
	medicineService    service.MedicineService       = service.New()
	medicineController controller.MedicineController = controller.New(medicineService)
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/get/{id}", medicineController.Get).Methods("GET")
	router.HandleFunc("/getAll", medicineController.GetAll).Methods("GET")
	router.HandleFunc("/add", medicineController.Add).Methods("POST")
	router.HandleFunc("/update", medicineController.Update).Methods("PUT")
	router.HandleFunc("/delete/{id}", medicineController.Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
