package main

import (
	rpc "assessment/adapter/grpc"
	pb "assessment/adapter/grpc/grpc_proto"
	adapter "assessment/adapter/rest"
	"assessment/adapter/rest/router"
	"assessment/helpers"
	db "assessment/infrastructure/db/postgres"
	"assessment/logger"
	"assessment/usecases"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	logService := logger.NewLogger()     //Initialize Logger
	helpers.InitializeLogger(logService) //Use Initialized Logger To Initialize The Helper Package
	repo, err := db.NewPostgresHandler() //Create A New (Postgres) Db Handler
	if err != nil {
		log.Panicln(err)
	}
	usecase := usecases.NewService(repo) //Create A New Service From The Previously Initialized Db Handler

	mode := os.Getenv("OP_MODE")

	if mode == "GRPC" {
		lis, err := net.Listen("tcp", ":50051")

		if err != nil {
			log.Fatal(err)
		}

		server := grpc.NewServer()
		pb.RegisterCarServiceServer(server, rpc.NewGRPCController(usecase))

		if err := server.Serve(lis); err != nil {
			log.Fatal(err)
		}
	} else if mode == "REST" {

		controller := adapter.NewController(usecase, logService)

		fmt.Println("Starting Server ....")

		if err := http.ListenAndServe(":"+os.Getenv("PORT"), router.InitRouter(controller)); err != nil {
			log.Panicln(err)
		}
	} else {
		log.Fatalf("Unknown Mode : %s\n", mode)
	}
}
