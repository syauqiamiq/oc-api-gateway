package userHandler

import (
	"ocApiGateway/services/userService"
)

type handler struct {
	userService userService.UserService
}

func NewHandler(userService userService.UserService) *handler {
	return &handler{
		userService: userService,
	}

}
