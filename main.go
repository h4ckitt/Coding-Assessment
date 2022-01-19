package main

import (
	adapter "assessment/adapter/http"
	"assessment/helpers"
	"assessment/infrastructure"
	"assessment/infrastructure/http/router"
	db "assessment/repository/postgres"
	"assessment/usecases"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := infrastructure.NewLogger()
	helpers.InitializeLogger(logger)
	repo, err := db.NewPostgresHandler()
	if err != nil {
		log.Panicln(err)
	}
	usecase := usecases.NewService(repo, logger)

	controller := adapter.NewController(usecase, logger)

	fmt.Println("Starting Server ....")

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router.InitRouter(controller)); err != nil {
		log.Panicln(err)
	}
}
