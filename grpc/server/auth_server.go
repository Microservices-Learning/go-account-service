package server

import (
	"context"
	"go-account-service/grpc/gen-proto/auth"
	"go-account-service/grpc/service"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	authService service.AuthService
}

func NewAuthServer(authService *service.AuthService) AuthServer {
	return AuthServer{authService: *authService}
}

func (s *AuthServer) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	return s.authService.Login(request)
}

func (s *AuthServer) Validate(ctx context.Context, request *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	panic("d")
}

func (s *AuthServer) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return s.authService.Register(request)
}
