package adapter

import (
	"assessment/adapter/presenter"
	"assessment/helpers"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"assessment/domain"
	"assessment/usecases"

	"github.com/gorilla/mux"
)

type Controller struct {
	Service domain.CarUseCase
	Logger  usecases.Logger
}

func NewController(service domain.CarUseCase, logger usecases.Logger) *Controller {
	return &Controller{
		Service: service,
		Logger:  logger,
	}
}

func (controller *Controller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var car domain.Car

	err := json.NewDecoder(r.Body).Decode(&car)

	if err != nil {
		controller.Logger.LogError("%s", err)
		helpers.ReturnFailure(r, w, http.StatusUnprocessableEntity, "The Request Could Not Be Processed "+
			"At This Time, Please Try Again Later.")
		return
	}

	err = controller.Service.Register(car)

	if err != nil {
		controller.Logger.LogError("%s", err)
		helpers.ReturnFailure(r, w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ReturnSuccess(r, w, http.StatusCreated, "")
}

func (controller *Controller) GetCarsByColor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var resultCars []presenter.Car

	color := r.URL.Query().Get("color")

	cars, err := controller.Service.GetCarsByColor(color)

	if err != nil {
		helpers.ReturnFailure(r, w, http.StatusUnprocessableEntity, "The Request "+
			"Could Not Be Processed At This Time, Please Try Again Later")
		return
	}

	if len(cars) == 0 {
		helpers.ReturnFailure(r, w, http.StatusNotFound, "The Requested Resource Was Not Found")
		return
	}

	for _, car := range cars {
		resultCar := presenter.Car{
			Name:       car.Name,
			Type:       car.Type,
			Color:      car.Color,
			SpeedRange: car.SpeedRange,
		}

		resultCars = append(resultCars, resultCar)
	}

	helpers.ReturnSuccess(r, w, http.StatusOK, resultCars)
}

func (controller *Controller) ViewCarDetails(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	matched, err := regexp.Match("^[0-9]+$", []byte(id))

	if err != nil {
		controller.Logger.LogError("%s\n", err)
		helpers.ReturnFailure(r, w, http.StatusUnprocessableEntity, "The Request "+
			"Could Not Be Processed At This Time, Please Try Again Later")
		return
	}

	if !matched {
		controller.Logger.LogError("%s\n", errors.New("id Is Not A Number"))
		helpers.ReturnFailure(r, w, http.StatusBadRequest, "ID Was Not A Valid Number")
		return
	}

	car, err := controller.Service.ViewDetails(id)

	if err != nil {
		controller.Logger.LogError("%s\n", err)
		helpers.ReturnFailure(r, w, http.StatusNotFound, "Requested Resource Was Not Found")
		return
	}

	result := presenter.Car{
		Name:       car.Name,
		Type:       car.Type,
		Color:      car.Color,
		SpeedRange: car.SpeedRange,
	}

	helpers.ReturnSuccess(r, w, http.StatusOK, result)
}
