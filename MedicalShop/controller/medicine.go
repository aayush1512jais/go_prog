package controller

import (
	"encoding/json"
	"fmt"
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
	ServerStatus(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	service *service.Service
}

func NewMedicineController(service *service.Service) *Controller {
	return &Controller{
		service: service,
	}

}
func (c *Controller) ServerStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Server Running!!")
}

func (c *Controller) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var medicine model.Medicine
	if err := json.NewDecoder(r.Body).Decode(&medicine); err == nil {
		if id, err := c.service.Add(medicine); err == nil {
			Message := struct {
				Message     string `json:"message,omitempty"`
				Medicine_id int    `json:"medicine_id,omitempty"`
			}{Message: "Added Successfully", Medicine_id: id}
			//log.Fatal(Message)
			json.NewEncoder(w).Encode(Message)

			return
		} else {
			apperrors.ErrorResponse(err, w)
		}
	} else {
		// err := apperrors.ErrorModel{
		// 	Message: "Failed to parse body",
		// 	//Error:   err,
		// 	Code: 302,
		// }
		// http.Error(w, err.Message, err.Code)
		apperrors.ErrorResponse(apperrors.ErrBadRequest, w)
	}
	//fmt.Fprintf(w, "Unsuccessful")

}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var medicine *model.Medicine
	if err := json.NewDecoder(r.Body).Decode(&medicine); err == nil {
		if err := c.service.Update(*medicine); err == nil {
			Message := struct {
				Message string `json:"message,omitempty"`
			}{Message: "Updated Successfully"}
			json.NewEncoder(w).Encode(Message)
			return
		} else {
			apperrors.ErrorResponse(err, w)
		}
	} else {
		//log.Fatal(err)
		// err := apperrors.ErrorModel{
		// 	Message: "Failed to parse body",
		// 	//	Error:   err,
		// 	Code: 302,
		// }
		//http.Error(w, err.Message, err.Code)
		apperrors.ErrorResponse(apperrors.ErrBadRequest, w)
	}

}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if id, err := strconv.Atoi(params["id"]); err == nil {
		if err := c.service.Delete(id); err == nil {
			Message := struct {
				Message string `json:"message,omitempty"`
			}{Message: "Deleted Successfully"}
			json.NewEncoder(w).Encode(Message)

			return
		} else {
			apperrors.ErrorResponse(err, w)
		}
	} else {
		// log.Fatal(err)
		// err := apperrors.ErrorModel{
		// 	Message: "Failed to parse params",
		// 	//Error:   err,
		// 	Code: 302,
		// }
		//http.Error(w, err.Message, err.Code)
		apperrors.ErrorResponse(apperrors.ErrBadRequest, w)
	}
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if id, err := strconv.Atoi(params["id"]); err == nil {
		medicine, err := c.service.Get(id)
		if err == nil {
			json.NewEncoder(w).Encode(medicine)
			return
		} else {
			apperrors.ErrorResponse(err, w)

		}
	} else {
		log.Println(err)
		// err := apperrors.ErrorModel{
		// 	Message: "Failed to parse params",
		// 	Error:   err,
		// 	Code:    302,
		// }
		//http.Error(w, err.Message, err.Code)
		apperrors.ErrorResponse(apperrors.ErrBadRequest, w)
	}

}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	medicines, err := c.service.GetAll()
	if err == nil {
		json.NewEncoder(w).Encode(medicines)
		return
	} else {
		apperrors.ErrorResponse(err, w)

	}

}
