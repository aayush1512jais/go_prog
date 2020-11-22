package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aayush1512jais/go_prog/MedicalShop/model"

	"github.com/aayush1512jais/go_prog/MedicalShop/service"
	"github.com/gorilla/mux"
)

type MedicineController interface {
	Add(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type medicineController struct {
	service service.MedicineService
}

func New(service service.MedicineService) MedicineController {
	return &medicineController{
		service: service,
	}

}

func (c *medicineController) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var medicine model.Medicine
	json.NewDecoder(r.Body).Decode(&medicine)
	if c.service.Add(medicine) {
		fmt.Fprintf(w, "Added Successfully")
		return
	}
	fmt.Fprintf(w, "Unsuccessful")

}

func (c *medicineController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var medicine model.Medicine
	json.NewDecoder(r.Body).Decode(&medicine)
	if c.service.Update(medicine) {
		fmt.Fprintf(w, "Updated Successfully")
		return
	}
	fmt.Fprintf(w, "Unsuccessful")
}

func (c *medicineController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if c.service.Delete(id) {
		fmt.Fprintf(w, "Deleted Successfully")
		return
	}
	fmt.Fprintf(w, "Unsuccessful")
}

func (c *medicineController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if medicine, status := c.service.Get(id); status == true {
		json.NewEncoder(w).Encode(medicine)
		return
	}
	//TODO: handle error
}

func (c *medicineController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c.service.GetAll())
}
