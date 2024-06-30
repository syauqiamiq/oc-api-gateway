package userService

import (
	"encoding/json"
	"fmt"
	"log"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *service) Register(payload dto.RegisterInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("SERVICE-ERR-R 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	registerData, err := helper.ApiRequest("POST", s.env.UserServiceUrl, "/user/register", jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-R 2: %v", err.Error())
		return registerData, err
	}

	return registerData, nil
}

func (s *service) GetProfile(userId string) (dto.HttpResponse, error) {

	path := fmt.Sprintf("/user/%s", userId)
	updateData, err := helper.ApiRequest("GET", s.env.UserServiceUrl, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GP 1: %v", err.Error())
		return updateData, err
	}

	return updateData, nil
}

func (s *service) UpdateProfile(userId string, payload dto.UpdateUserInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("SERVICE-ERR-UP 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	path := fmt.Sprintf("/user/%s", userId)
	updateData, err := helper.ApiRequest("PUT", s.env.UserServiceUrl, path, jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-UP 2: %v", err.Error())
		return updateData, err
	}

	return updateData, nil
}

func (s *service) Logout(payload dto.LogoutBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("SERVICE-ERR-LG 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	logoutData, err := helper.ApiRequest("POST", s.env.UserServiceUrl, "/user/logout", jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-LG 2: %v", err.Error())
		return logoutData, err
	}

	return logoutData, nil
}

type MyClaims struct {
	jwt.RegisteredClaims
	Name   string
	Email  string
	UserID string
}

func (s *service) Login(payload dto.LoginInputBody) (dto.HttpResponse, error) {

	loginBody, err := json.Marshal(payload)
	if err != nil {
		log.Printf("SERVICE-ERR-L 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	loginReponse, err := helper.ApiRequest("POST", s.env.UserServiceUrl, "/user/login", loginBody)
	if err != nil {
		log.Printf("SERVICE-ERR-L 2: %v", err.Error())
		return loginReponse, err
	}

	return loginReponse, nil
}

func (s *service) SaveRefreshToken(userId string, refreshToken string) (dto.HttpResponse, error) {
	payloadSaveRefreshTokenBody := dto.SaveRefreshTokenInputBody{
		RefreshToken: refreshToken,
		UserID:       userId,
	}

	saveRefreshTokenBody, err := json.Marshal(payloadSaveRefreshTokenBody)
	if err != nil {
		log.Printf("SERVICE-ERR-SRT 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	saveRefreshTokenResponse, err := helper.ApiRequest("POST", s.env.UserServiceUrl, "/refresh-token", saveRefreshTokenBody)
	if err != nil {
		log.Printf("SERVICE-ERR-SRT 2: %v", err.Error())
		return saveRefreshTokenResponse, err
	}

	return saveRefreshTokenResponse, nil

}

func (s *service) GenerateJWTToken(userData dto.User) (tokenString string, refreshTokenString string, err error) {
	tokenExpireTime, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRED"))
	if err != nil {
		log.Printf("SERVICE-ERR-JWT 1: %v", err.Error())
		return "", "", err
	}

	refreshTokenExpireTime, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRED"))
	if err != nil {
		log.Printf("SERVICE-ERR-JWT 2: %v", err.Error())
		return "", "", err
	}

	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "OC-APIGATEWAY",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenExpireTime) * time.Second)),
		},
		Name:   userData.Name,
		Email:  userData.Email,
		UserID: userData.ID,
	}

	refreshTokenClaims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "OC-APIGATEWAY",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(refreshTokenExpireTime) * time.Second)),
		},
		Name:   userData.Name,
		Email:  userData.Email,
		UserID: userData.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenString, err = token.SignedString([]byte(s.env.JwtTokenSecret))
	if err != nil {
		log.Printf("SERVICE-ERR-JWT 3: %v", err.Error())
		return "", "", err
	}

	refreshTokenString, err = refreshToken.SignedString([]byte(s.env.JwtRefreshTokenSecret))
	if err != nil {
		log.Printf("SERVICE-ERR-JWT 4: %v", err.Error())
		return "", "", err
	}

	return tokenString, refreshTokenString, err
}
