package main

import (
	"assessment/infrastructure"
	"assessment/infrastructure/router"
	"assessment/interfaces"
	db "assessment/repository/postgres"
	"assessment/usecases"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := infrastructure.NewLogger()

	repo, err := db.NewPostgresHandler()

	if err != nil {
		log.Panicln(err)
	}

	usecase := usecases.NewService(repo, logger)

	controller := interfaces.NewController(usecase, logger)

	fmt.Println("Starting Server ....")

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router.InitRouter(controller)); err != nil {
		log.Panicln(err)
	}
}

/*func main() {
	fmt.Println("vim-go")
}*/
