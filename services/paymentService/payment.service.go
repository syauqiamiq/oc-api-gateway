package paymentService

import (
	"encoding/json"
	"fmt"
	"log"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
)

func (s *service) CheckoutOrder(input dto.CheckoutOrderInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("SERVICE-ERR-CO 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	data, err := helper.ApiRequest("POST", PAYMENT_SERVICE_URL, "/order", jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-CO 2: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetOrder(userId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/order?userId=%s", userId)
	data, err := helper.ApiRequest("GET", PAYMENT_SERVICE_URL, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GL 1: %v", err.Error())
		return data, err
	}

	return data, nil
}
