package helpers

import (
	"assessment/usecases"
	"encoding/json"
	"net/http"
)

var logService usecases.Logger

type success struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

type failure struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}

func InitializeLogger(log usecases.Logger) {
	logService = log
}

/*ReturnFailure : Responsible For Gracefully Handling A Failed REST Operation
Params:
	- *http.Request
	- http.ResponseWriter
	- code
	- err
*/
func ReturnFailure(r *http.Request, w http.ResponseWriter, code int, err string) {
	response := failure{
		Status: "failure",
		Code:   code,
		Error:  err,
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&response)

	logService.LogAccess(r, code)
}

/*ReturnSuccess : Responsible For Gracefully Handling A Successful REST Operation
Params:
	- *http.Request
	- http.ResponseWriter
	- code
	- data
*/
func ReturnSuccess(r *http.Request, w http.ResponseWriter, code int, data interface{}) {
	response := success{
		Status: "success",
		Code:   code,
		Data:   data,
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&response)
	logService.LogAccess(r, code)
}
