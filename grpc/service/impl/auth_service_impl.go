package impl

import (
	"go-account-service/config"
	"go-account-service/grpc/gen-proto/auth"
	"go-account-service/grpc/gen-proto/common"
	"go-account-service/model"
	"go-account-service/util"
	"net/http"
)

type AuthServiceImpl struct {
}

func (a *AuthServiceImpl) Login(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	var user model.User

	if result := config.DB.Where(&model.User{Email: req.Username}).First(&user); result.Error != nil {
		return &auth.LoginResponse{
			Success: false,
			Response: &auth.LoginResponse_Error{Error: &common.Error{
				Code:    common.ErrorCode_USERNAME_OR_PASSWORD_INVALID,
				Message: "Username or password invalid",
			}},
		}, nil
	}

	match := util.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &auth.LoginResponse{
			Success: false,
			Response: &auth.LoginResponse_Error{Error: &common.Error{
				Code:    http.StatusNotFound,
				Message: "User not found",
			}},
		}, nil
	}

	token, _ := util.JwtUtil.GenerateToken(user)

	return &auth.LoginResponse{
		Success: true,
		Response: &auth.LoginResponse_Data_{Data: &auth.LoginResponse_Data{
			AccessToken:  token,
			RefreshToken: token,
			User: &auth.AuthResponse{
				Email: user.Email,
			},
		}},
	}, nil
}

func (a *AuthServiceImpl) Register(req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	var user model.User
	if result := config.DB.Where(&model.User{Email: req.Email}).First(&user); result.Error == nil {
		return &auth.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = util.HashPassword(req.Password)

	config.DB.Create(&user)

	return &auth.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}
