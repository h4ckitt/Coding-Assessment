package interfaces

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"assessment/domain"
	"assessment/usecases"

	"github.com/gorilla/mux"
)

type Controller struct {
	Service usecases.Service
	Logger  usecases.Logger
}

func (controller *Controller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var car domain.Schema

	err := json.NewDecoder(r.Body).Decode(&car)

	if err != nil {
		controller.Logger.LogError("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		//	json.NewEncoder(w).Encode(err)
		return
	}

	id, err := controller.Service.Register(car)

	if err != nil {
		controller.Logger.LogError("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		//	json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	controller.Logger.LogAccess("%s - %s -> %s\n", r.RemoteAddr, r.Method, r.URL)
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

	json.NewEncoder(w).Encode(cars)
	controller.Logger.LogAccess("%s - %s -> %s\n", r.RemoteAddr, r.Method, r.URL)
}

func (controller *Controller) ViewCarDetails(w http.ResponseWriter, r *http.Request) {
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
}
