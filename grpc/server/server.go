package server

import (
	"go-account-service/config"
	"go-account-service/grpc/gen-proto/auth"
	"go-account-service/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
}

func NewServer() Server {
	return Server{}
}

func (server *Server) InitializeGrpcServer() {
	authService := service.NewAuthService()
	authServer := NewAuthServer(&authService)

	lis, err := net.Listen("tcp", config.Config.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, &authServer)

	reflection.Register(grpcServer)

	log.Printf("start gRPC server on %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
