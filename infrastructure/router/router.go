package router

import (
	"assessment/interfaces"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter(controller *interfaces.Controller) *mux.Router {
	const APIVERSION = "/v1"
	router := mux.NewRouter()
	versionRouter := router.PathPrefix(APIVERSION).Subrouter()
	pathSubrouter := versionRouter.PathPrefix("/cars").Subrouter()

	pathSubrouter.HandleFunc("", controller.Register).Methods(http.MethodPost)
	pathSubrouter.HandleFunc("/{id}", controller.ViewCarDetails).Methods(http.MethodGet)
	pathSubrouter.HandleFunc("", controller.GetCarsByColor).Queries("color", "{color}").Methods(http.MethodGet)

	return router

}
