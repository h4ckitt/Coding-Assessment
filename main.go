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

	switch mode {
	case "GRPC":
		port := os.Getenv("GRPC_PORT")

		if port == "" {
			log.Fatal(err)
		}

		lis, err := net.Listen("tcp", ":"+port)

		if err != nil {
			log.Fatal(err)
		}

		server := grpc.NewServer()
		pb.RegisterCarServiceServer(server, rpc.NewGRPCController(usecase))

		log.Println("Starting GRPC Server")
		if err := server.Serve(lis); err != nil {
			log.Fatal(err)
		}

	case "REST":
		port := os.Getenv("HTTP_PORT")

		if port == "" {
			log.Fatalln("Http Port Not Set")
		}

		controller := adapter.NewController(usecase, logService)

		fmt.Println("Starting Server ....")

		if err := http.ListenAndServe(":"+os.Getenv("PORT"), router.InitRouter(controller)); err != nil {
			log.Panicln(err)
		}

	default:
		log.Fatalf("Unknown Mode : %s\n", mode)
	}
}
