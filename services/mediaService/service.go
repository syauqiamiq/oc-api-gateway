package mediaService

import (
	"ocApiGateway/dto"
	"ocApiGateway/helper"
)

type MediaService interface {
	GetAllMedia() (dto.HttpResponse, error)
	DeleteMediaByID(id string) (dto.HttpResponse, error)
	UploadMediaImage(data dto.UploadMediaBody) (dto.HttpResponse, error)
}

type service struct {
	env helper.Env
}

func NewService(env helper.Env) *service {
	return &service{
		env: env,
	}
}
