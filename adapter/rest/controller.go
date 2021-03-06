package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"assessment/adapter/presenter"
	"assessment/domain"
	"assessment/helpers"
	"assessment/usecases"

	"github.com/gorilla/mux"
)

type Controller struct {
	Service usecases.CarUseCase
	Logger  usecases.Logger
}

/*NewController : Creates A New REST Controller
Params:
	- usecases.CarUseCase
	- logger.Logger

Returns:
	- *Controller
*/
func NewController(service usecases.CarUseCase, logger usecases.Logger) *Controller {
	return &Controller{
		Service: service,
		Logger:  logger,
	}
}

/*Register : This Function Is Responsible For The Deserialization Of The JSON Input
And Successful Registration Of The Resulting Car Object
*/
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

	helpers.ReturnSuccess(r, w, http.StatusCreated, "Record Created Successfully")
}

/*GetCarsByColor : Gets Color Variable From The URI And Returns The Cars Whose Colors Match
The Query
*/
func (controller *Controller) GetCarsByColor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var resultCars []presenter.Car

	color := r.URL.Query().Get("color")

	cars, err := controller.Service.GetCarsByColor(color)

	if err != nil {
		helpers.ReturnFailure(r, w, http.StatusNotFound, "The Requested Resource Was Not Found")
		return
	}

	/*if len(cars) == 0 {
		helpers.ReturnFailure(r, w, http.StatusNotFound, "The Requested Resource Was Not Found")
		return
	}*/

	for _, car := range cars {
		resultCar := presenter.Car{
			Name:       car.Name,
			Type:       car.Type,
			Color:      car.Color,
			SpeedRange: car.SpeedRange,
			Features:   car.Features,
		}

		resultCars = append(resultCars, resultCar)
	}

	helpers.ReturnSuccess(r, w, http.StatusOK, resultCars)
}

/*GetCarsByType : Extracts The Type Query And Returns All Cars Whose Type Matches The
Extracted Query
*/
func (controller *Controller) GetCarsByType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var resultCars []presenter.Car

	carType := r.URL.Query().Get("type")

	cars, err := controller.Service.GetCarsByType(carType)

	if err != nil {
		helpers.ReturnFailure(r, w, http.StatusNotFound, "The Requested Resource Was Not Found")
		return
	}

	/*if len(cars) == 0 {
		helpers.ReturnFailure(r, w, http.StatusNotFound, "The Requested Resource Was Not Found")
		return
	}*/

	for _, car := range cars {
		resultCar := presenter.Car{
			Name:       car.Name,
			Type:       car.Type,
			Color:      car.Color,
			SpeedRange: car.SpeedRange,
			Features:   car.Features,
		}

		resultCars = append(resultCars, resultCar)
	}

	helpers.ReturnSuccess(r, w, http.StatusOK, resultCars)
}

/*ViewCarDetails : Fetches The Car Whose ID Matches The Specified Query

 */
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
		Features:   car.Features,
	}

	helpers.ReturnSuccess(r, w, http.StatusOK, result)
}
