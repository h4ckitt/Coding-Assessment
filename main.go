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
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	logger := logger.NewLogger()
	helpers.InitializeLogger(logger)
	repo, err := db.NewPostgresHandler()
	if err != nil {
		log.Panicln(err)
	}
	usecase := usecases.NewService(repo)

	mode := flag.Bool("grpc", false, "grpc mode")
	flag.Parse()

	if *mode {
		lis, err := net.Listen("tcp", ":50051")

		if err != nil {
			log.Fatal(err)
		}

		server := grpc.NewServer()
		pb.RegisterCarServiceServer(server, rpc.NewGRPCController(usecase))

		if err := server.Serve(lis); err != nil {
			log.Fatal(err)
		}
	} else {

		controller := adapter.NewController(usecase, logger)

		fmt.Println("Starting Server ....")

		if err := http.ListenAndServe(":"+os.Getenv("PORT"), router.InitRouter(controller)); err != nil {
			log.Panicln(err)
		}
	}
}
