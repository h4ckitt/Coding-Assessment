package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"assessment/domain"
	"assessment/usecases"

	"github.com/gorilla/mux"
)

type Controller struct {
	Service *usecases.Service
	Logger  usecases.Logger
}

func NewController(service *usecases.Service, logger usecases.Logger) *Controller {
	var returnValue Controller

	returnValue.Service = service
	returnValue.Logger = logger

	return &returnValue
}

func (controller *Controller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var car domain.Schema

	err := json.NewDecoder(r.Body).Decode(&car)

	fmt.Println(car)

	if err != nil {
		controller.Logger.LogError("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		//	json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("Modebiyi")

	err = controller.Service.Register(car)

	if err != nil {
		controller.Logger.LogError("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		//	json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	controller.Logger.LogAccess(r, http.StatusCreated)
	//json.NewEncoder(w).Encode(id)
}

func (controller *Controller) GetCarsByColor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	color := r.URL.Query().Get("color")

	cars, err := controller.Service.GetCarsByColor(color)

	if err != nil {
		controller.Logger.LogError("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(cars) == 0 {
		w.WriteHeader(http.StatusNotFound)
		controller.Logger.LogAccess(r, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cars)
	controller.Logger.LogAccess(r, http.StatusOK)
}

func (controller *Controller) ViewCarDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Am Hia Ohh")
	id := mux.Vars(r)["id"]

	matched, err := regexp.Match("^[0-9]+$", []byte(id))

	if err != nil {
		controller.Logger.LogError("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !matched {
		controller.Logger.LogError("%s\n", errors.New("id Is Not A Number"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	car, err := controller.Service.ViewDetails(id)

	if err != nil {
		controller.Logger.LogError("%s\n", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(car)
	controller.Logger.LogAccess(r, http.StatusOK)
}
