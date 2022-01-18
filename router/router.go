package router

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	const APIVERSION = "/v1"
	router := mux.NewRouter()
	versionRouter := router.PathPrefix(APIVERSION).Subrouter()
}
