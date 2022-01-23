package main

import (
	adapter "assessment/adapter/http"
	"assessment/helpers"
	"assessment/infrastructure"
	rpc "assessment/infrastructure/grpc"
	pb "assessment/infrastructure/grpc/grpc_proto"
	"assessment/infrastructure/rest/router"
	db "assessment/repository/postgres"
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
	logger := infrastructure.NewLogger()
	helpers.InitializeLogger(logger)
	repo, err := db.NewPostgresHandler()
	if err != nil {
		log.Panicln(err)
	}
	usecase := usecases.NewService(repo, logger)

	mode := flag.Bool("grpc", false, "grpc mode")
	flag.Parse()

	if *mode {
		lis, err := net.Listen("tcp", ":50051")

		if err != nil {
			log.Fatal(err)
		}

		server := grpc.NewServer()
		pb.RegisterCarServiceServer(server, rpc.NewGRPCController(usecase, logger))

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
