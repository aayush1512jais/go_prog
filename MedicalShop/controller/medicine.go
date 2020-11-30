package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/aayush1512jais/go_prog/MedicalShop/apperrors"

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
	if err := json.NewDecoder(r.Body).Decode(&medicine); err == nil {
		if id, err := c.service.Add(medicine); id != -1 {
			Message := struct {
				Message     string `json:"message,omitempty"`
				Medicine_id int    `json:"medicine_id,omitempty"`
			}{Message: "Added Successfully", Medicine_id: id}
			//log.Fatal(Message)
			json.NewEncoder(w).Encode(Message)

			return
		} else {
			http.Error(w, err.Message+"\n error: "+err.Error.Error(), err.Code)
		}
	} else {
		err := apperrors.ErrorModel{
			Message: "Failed to parse body",
			Error:   err,
			Code:    302,
		}
		http.Error(w, err.Message+"\n error: "+err.Error.Error(), err.Code)
	}
	//fmt.Fprintf(w, "Unsuccessful")

}

func (c *medicineController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var medicine *model.Medicine
	if err := json.NewDecoder(r.Body).Decode(&medicine); err == nil {
		if status, err := c.service.Update(*medicine); status {
			Message := struct {
				Message string `json:"message,omitempty"`
			}{Message: "Updated Successfully"}
			json.NewEncoder(w).Encode(Message)
			return
		} else {
			http.Error(w, err.Message+"\n error: "+err.Error.Error(), err.Code)
		}
	} else {
		//log.Fatal(err)
		err := apperrors.ErrorModel{
			Message: "Failed to parse body",
			Error:   err,
			Code:    302,
		}
		http.Error(w, err.Message+"\n error: "+err.Error.Error(), err.Code)
	}

}

func (c *medicineController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if id, err := strconv.Atoi(params["id"]); err == nil {
		if status, err := c.service.Delete(id); status {
			Message := struct {
				Message string `json:"message,omitempty"`
			}{Message: "Deleted Successfully"}
			json.NewEncoder(w).Encode(Message)
			return
		} else {
			http.Error(w, err.Message+"\n error: "+err.Error.Error(), err.Code)
		}
	} else {
		log.Fatal(err)
		err := apperrors.ErrorModel{
			Message: "Failed to parse params",
			Error:   err,
			Code:    302,
		}
		http.Error(w, err.Message+"\n error: "+err.Error.Error(), err.Code)
	}
}

func (c *medicineController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if id, err := strconv.Atoi(params["id"]); err == nil {
		medicine, err := c.service.Get(id)
		if err == (apperrors.ErrorModel{}) {
			json.NewEncoder(w).Encode(medicine)
			return
		} else {
			http.Error(w, err.Message+"\n error: "+err.Error.Error(), err.Code)

		}
	} else {
		log.Fatal(err)
		err := apperrors.ErrorModel{
			Message: "Failed to parse params",
			Error:   err,
			Code:    302,
		}
		http.Error(w, err.Message+"\n error: "+err.Error.Error(), err.Code)
	}

}

func (c *medicineController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	medicines, err := c.service.GetAll()
	if err == (apperrors.ErrorModel{}) {
		json.NewEncoder(w).Encode(medicines)
		return
	} else {
		http.Error(w, err.Message+"\n error: "+err.Error.Error(), err.Code)

	}

}
