package paymentHandler

import "ocApiGateway/services/paymentService"

type handler struct {
	paymentService paymentService.PaymentService
}

func NewHandler(paymentService paymentService.PaymentService) *handler {
	return &handler{
		paymentService: paymentService,
	}

}
