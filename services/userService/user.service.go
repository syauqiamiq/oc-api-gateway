package userService

import (
	"encoding/json"
	"log"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var (
	_        = godotenv.Load(".env")
	BASE_URL = os.Getenv("USER_SERVICE_URL")
)

func (s *service) Register(payload dto.RegisterInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("SERVICE-R-ERR 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	registerData, err := helper.ApiRequest("POST", BASE_URL, "/user/register", jsonData)
	if err != nil {
		log.Printf("SERVICE-R-ERR 2: %v", err.Error())
		return registerData, err
	}

	return registerData, nil
}

type MyClaims struct {
	jwt.RegisteredClaims
	Name   string
	Email  string
	UserID string
}

var (
	JWT_TOKEN_SECRET         = os.Getenv("JWT_TOKEN_SECRET")
	JWT_REFRESH_TOKEN_SECRET = os.Getenv("JWT_REFRESH_TOKEN_SECRET")
)

func (s *service) Login(payload dto.LoginInputBody) (dto.HttpResponse, error) {

	loginBody, err := json.Marshal(payload)
	if err != nil {
		log.Printf("SERVICE-ERR-L 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	loginReponse, err := helper.ApiRequest("POST", BASE_URL, "/user/login", loginBody)
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

	saveRefreshTokenResponse, err := helper.ApiRequest("POST", BASE_URL, "/refresh-token", saveRefreshTokenBody)
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "OC-APIGATEWAY",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(refreshTokenExpireTime) * time.Second)),
	})

	tokenString, err = token.SignedString([]byte(JWT_TOKEN_SECRET))
	if err != nil {
		log.Printf("SERVICE-ERR-JWT 3: %v", err.Error())
		return "", "", err
	}

	refreshTokenString, err = refreshToken.SignedString([]byte(JWT_REFRESH_TOKEN_SECRET))
	if err != nil {
		log.Printf("SERVICE-ERR-JWT 4: %v", err.Error())
		return "", "", err
	}

	return tokenString, refreshTokenString, err
}
