package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	CourseServiceUrl       string
	UserServiceUrl         string
	MediaServiceUrl        string
	PaymentServiceUrl      string
	JwtTokenSecret         string
	JwtTokenExpired        string
	JwtRefreshTokenSecret  string
	JwtRefreshTokenExpired string
}

func GetEnv() Env {
	var (
		COURSE_SERVICE_URL        string
		USER_SERVICE_URL          string
		MEDIA_SERVICE_URL         string
		PAYMENT_SERVICE_URL       string
		JWT_TOKEN_SECRET          string
		JWT_REFRESH_TOKEN_SECRET  string
		JWT_TOKEN_EXPIRED         string
		JWT_REFRESH_TOKEN_EXPIRED string
	)

	if os.Getenv("ENV") == "production" {
		COURSE_SERVICE_URL = ReadSecretFile(os.Getenv("COURSE_SERVICE_URL"))
		USER_SERVICE_URL = ReadSecretFile(os.Getenv("USER_SERVICE_URL"))
		MEDIA_SERVICE_URL = ReadSecretFile(os.Getenv("MEDIA_SERVICE_URL"))
		PAYMENT_SERVICE_URL = ReadSecretFile(os.Getenv("PAYMENT_SERVICE_URL"))
		JWT_TOKEN_SECRET = ReadSecretFile(os.Getenv("JWT_TOKEN_SECRET"))
		JWT_TOKEN_EXPIRED = ReadSecretFile(os.Getenv("JWT_TOKEN_EXPIRED"))
		JWT_REFRESH_TOKEN_SECRET = ReadSecretFile(os.Getenv("JWT_REFRESH_TOKEN_SECRET"))
		JWT_REFRESH_TOKEN_EXPIRED = ReadSecretFile(os.Getenv("JWT_REFRESH_TOKEN_EXPIRED"))
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		COURSE_SERVICE_URL = os.Getenv("COURSE_SERVICE_URL")
		USER_SERVICE_URL = os.Getenv("USER_SERVICE_URL")
		MEDIA_SERVICE_URL = os.Getenv("MEDIA_SERVICE_URL")
		PAYMENT_SERVICE_URL = os.Getenv("PAYMENT_SERVICE_URL")
		JWT_TOKEN_SECRET = os.Getenv("JWT_TOKEN_SECRET")
		JWT_TOKEN_EXPIRED = os.Getenv("JWT_TOKEN_EXPIRED")
		JWT_REFRESH_TOKEN_SECRET = os.Getenv("JWT_REFRESH_TOKEN_SECRET")
		JWT_REFRESH_TOKEN_EXPIRED = os.Getenv("JWT_REFRESH_TOKEN_EXPIRED")

	}
	return Env{
		CourseServiceUrl:       COURSE_SERVICE_URL,
		UserServiceUrl:         USER_SERVICE_URL,
		MediaServiceUrl:        MEDIA_SERVICE_URL,
		PaymentServiceUrl:      PAYMENT_SERVICE_URL,
		JwtTokenSecret:         JWT_TOKEN_SECRET,
		JwtTokenExpired:        JWT_TOKEN_EXPIRED,
		JwtRefreshTokenSecret:  JWT_REFRESH_TOKEN_SECRET,
		JwtRefreshTokenExpired: JWT_REFRESH_TOKEN_EXPIRED,
	}
}
