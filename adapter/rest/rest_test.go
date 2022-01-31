package rest_test

import (
	"assessment/adapter/rest"
	"assessment/adapter/rest/router"
	"assessment/domain"
	"assessment/helpers"
	"assessment/infrastructure/db/inmemoryteststore"
	"assessment/logger"
	"assessment/usecases"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

var rt *mux.Router

type RESTTestSuite struct {
	suite.Suite
	controller *rest.Controller
}

func (r *RESTTestSuite) SetupSuite() {
	repo := inmemoryteststore.NewMemoryStore()

	cars := []domain.Car{
		{
			Name:       "Toyota Tundra",
			Color:      "red",
			Type:       "suv",
			SpeedRange: 180,
			Features:   []string{"panorama", "surround-system"},
		},
		{
			Name:       "Tesla Model S Plaid",
			Color:      "blue",
			Type:       "sedan",
			SpeedRange: 240,
			Features:   []string{"auto-parking", "panorama", "surround-system"},
		},
		{
			Name:       "Ducati 99R",
			Color:      "green",
			Type:       "motor-bike",
			SpeedRange: 200,
			Features:   []string{"surround-system"},
		},
		{
			Name:       "Ford ES500",
			Color:      "blue",
			Type:       "van",
			SpeedRange: 80,
			Features:   []string{"surround-system", "sunroof", "surround-system"},
		},
	}

	for _, car := range cars {
		err := repo.Store(car)

		require.NoError(r.T(), err)
	}
	service := usecases.NewService(repo)
	logService := logger.NewTestLogger()

	r.controller = rest.NewController(service, logService)
	rt = router.InitRouter(r.controller)
	helpers.InitializeLogger(logService)
}

func TestRESTSuite(t *testing.T) {
	suite.Run(t, new(RESTTestSuite))
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	rt.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected: %v, Got: %v\n", expected, actual)
	}
}

func constructRequest(t *testing.T, car domain.Car) *http.Request {
	jsonBody, err := json.Marshal(car)

	if err != nil {
		t.Fatal("An Error Occurred While Marshalling Car Struct")
	}
	return httptest.NewRequest("POST", "/v1/cars", bytes.NewBuffer(jsonBody))

}

func (r *RESTTestSuite) TestRegister() {
	car := domain.Car{
		Name:       "Toyota Tundra",
		Color:      "Blue",
		Type:       "SUV",
		SpeedRange: 240,
		Features:   []string{"sunroof", "panorama"},
	}

	response := executeRequest(constructRequest(r.T(), car))

	checkResponseCode(r.T(), response.Code, http.StatusCreated)

	car.Color = "brown"

	response = executeRequest(constructRequest(r.T(), car))

	checkResponseCode(r.T(), response.Code, http.StatusBadRequest)

}

func (r *RESTTestSuite) TestGetCarsByColor() {
	req := httptest.NewRequest("GET", "/v1/cars?color=blue", nil)

	response := executeRequest(req)

	checkResponseCode(r.T(), http.StatusOK, response.Code)

	req = httptest.NewRequest("GET", "/v1/cars?color=hazelnut", nil)

	response = executeRequest(req)

	checkResponseCode(r.T(), http.StatusNotFound, response.Code)
}

func (r *RESTTestSuite) TestGetCarsByType() {

	req := httptest.NewRequest("GET", "/v1/cars?type=sedan", nil)

	response := executeRequest(req)

	checkResponseCode(r.T(), http.StatusOK, response.Code)

	req = httptest.NewRequest("GET", "/v1/cars?type=footwagon", nil)

	response = executeRequest(req)

	checkResponseCode(r.T(), http.StatusNotFound, response.Code)
}

func (r *RESTTestSuite) TestViewDetails() {

	req := httptest.NewRequest("GET", "/v1/cars/2", nil)

	response := executeRequest(req)

	checkResponseCode(r.T(), http.StatusOK, response.Code)

	req = httptest.NewRequest("GET", "/v1/cars/28", nil)

	response = executeRequest(req)

	checkResponseCode(r.T(), http.StatusNotFound, response.Code)

	req = httptest.NewRequest("GET", "/v1/cars/area99", nil)

	response = executeRequest(req)

	checkResponseCode(r.T(), http.StatusBadRequest, response.Code)
}
