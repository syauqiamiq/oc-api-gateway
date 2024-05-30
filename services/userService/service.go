package userService

import "ocApiGateway/dto"

type UserService interface {
	Register(payload dto.RegisterInputBody) (dto.HttpResponse, error)
	Login(payload dto.LoginInputBody) (dto.HttpResponse, error)
	SaveRefreshToken(userId string, refreshToken string) (dto.HttpResponse, error)
	GenerateJWTToken(userData dto.User) (tokenString string, refreshTokenString string, err error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}
