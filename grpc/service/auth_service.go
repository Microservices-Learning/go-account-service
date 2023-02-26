package service

import (
	"go-account-service/grpc/gen-proto/auth"
	"go-account-service/grpc/service/impl"
)

type AuthService interface {
	Register(req *auth.RegisterRequest) (*auth.RegisterResponse, error)
	Login(req *auth.LoginRequest) (*auth.LoginResponse, error)
}

func NewAuthService() AuthService {
	return &impl.AuthServiceImpl{}
}
