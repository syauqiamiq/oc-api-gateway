package userService

import (
	"ocApiGateway/dto"
	"os"

	"github.com/joho/godotenv"
)

type UserService interface {
	Register(payload dto.RegisterInputBody) (dto.HttpResponse, error)
	Login(payload dto.LoginInputBody) (dto.HttpResponse, error)
	SaveRefreshToken(userId string, refreshToken string) (dto.HttpResponse, error)
	GenerateJWTToken(userData dto.User) (tokenString string, refreshTokenString string, err error)
	Logout(payload dto.LogoutBody) (dto.HttpResponse, error)
	UpdateProfile(userId string, payload dto.UpdateUserInputBody) (dto.HttpResponse, error)
	GetProfile(userId string) (dto.HttpResponse, error)
	ValidateRefreshToken(refreshToken string) (dto.HttpResponse, error)
	GenerateNewAccessToken(payload dto.RefreshTokenInputBody) (newAccessToken string, err error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

var (
	_        = godotenv.Load(".env")
	BASE_URL = os.Getenv("USER_SERVICE_URL")
)
