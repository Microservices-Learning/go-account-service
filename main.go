package main

import (
	"go-account-service/config"
	"go-account-service/grpc/server"
	"log"
)

func main() {
	start()
}

func start() {
	err := config.Config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	dbErr := config.ConnectDB()
	if dbErr != nil {
		log.Fatal("Cannot connect db:", dbErr)
	}

	server := server.NewServer()
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
	server.InitializeGrpcServer()
}
