package helpers

import (
	"assessment/usecases"
	"encoding/json"
	"net/http"
)

var logger usecases.Logger

func InitializeLogger(log usecases.Logger) {
	logger = log
}

func ReturnFailure(r *http.Request, w http.ResponseWriter, code int, err string) {
	response := failure{
		Status: "failure",
		Code:   code,
		Error:  err,
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&response)
	logger.LogAccess(r, code)
}

func ReturnSuccess(r *http.Request, w http.ResponseWriter, code int, data interface{}) {
	response := success{
		Status: "success",
		Code:   code,
		Data:   data,
	}
	
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&response)
	logger.LogAccess(r, code)
}
