package paymentService

import (
	"ocApiGateway/dto"
	"ocApiGateway/helper"
)

type PaymentService interface {
	CheckoutOrder(input dto.CheckoutOrderInputBody) (dto.HttpResponse, error)
	GetOrder(userId string) (dto.HttpResponse, error)
}

type service struct {
	env helper.Env
}

func NewService(env helper.Env) *service {
	return &service{
		env: env,
	}
}
