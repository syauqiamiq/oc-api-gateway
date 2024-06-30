package userService

import (
	"errors"
	"fmt"
	"log"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *service) ValidateRefreshToken(refreshToken string) (dto.HttpResponse, error) {

	refreshTokenResponse, err := helper.ApiRequest("GET", s.env.UserServiceUrl, fmt.Sprintf("/refresh-token?refresh_token=%s", refreshToken), nil)
	if err != nil {
		log.Printf("SERVICE-ERR-VRT 1: %v", err.Error())
		return refreshTokenResponse, err
	}
	return refreshTokenResponse, nil
}

func (s *service) GenerateNewAccessToken(payload dto.RefreshTokenInputBody) (newAccessToken string, err error) {

	token, err := jwt.ParseWithClaims(payload.RefreshToken, &dto.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.env.JwtRefreshTokenSecret), nil
	})
	if err != nil {
		log.Printf("SERVICE-ERR-GNAT 1: %v", errors.New("invalid refresh token"))
		return "", errors.New("invalid refresh token")
	}

	if claims, ok := token.Claims.(*dto.MyClaims); ok && token.Valid {
		// Check if the token is expired
		if claims.ExpiresAt.Time.Before(time.Now()) {
			log.Printf("SERVICE-ERR-GNAT 2: %v", errors.New("refresh token expired"))
			return "", errors.New("refresh token expired")
		}

		if payload.Email != claims.Email {
			log.Printf("SERVICE-ERR-GNAT 3: %v", errors.New("refresh token not match with provided email"))
			return "", errors.New("refresh token not match with provided email")
		}

		// Generate new access token
		tokenExpireTime, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRED"))
		if err != nil {
			log.Printf("SERVICE-ERR-GNAT 4: %v", err.Error())
			return "", err
		}

		accessTokenclaims := MyClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "OC-APIGATEWAY",
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenExpireTime) * time.Second)),
			},
			Name:   claims.Name,
			Email:  claims.Email,
			UserID: claims.UserID,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenclaims)

		tokenString, err := token.SignedString([]byte(s.env.JwtTokenSecret))
		if err != nil {
			log.Printf("SERVICE-ERR-GNAT 5: %v", err.Error())
			return "", err
		}

		return tokenString, nil

	} else {
		log.Printf("SERVICE-ERR-GNAT 6: %v", errors.New("something wrong with refresh token"))
		return "", errors.New("something wrong with refresh token")
	}
}
