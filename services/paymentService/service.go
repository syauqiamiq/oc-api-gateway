package paymentService

import (
	"ocApiGateway/dto"
	"os"

	"github.com/joho/godotenv"
)

type PaymentService interface {
	CheckoutOrder(input dto.CheckoutOrderInputBody) (dto.HttpResponse, error)
	GetOrder(userId string) (dto.HttpResponse, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

var (
	_                   = godotenv.Load(".env")
	PAYMENT_SERVICE_URL = os.Getenv("PAYMENT_SERVICE_URL")
)
