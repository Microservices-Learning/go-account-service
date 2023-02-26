package common

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email    string
	Password string
}

type RegisterResponse struct {
	Status int64
	Error  string
}

type JwtClaims struct {
	jwt.StandardClaims
	Id    uuid.UUID
	Email string
}
