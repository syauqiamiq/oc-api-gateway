package mediaService

import "ocApiGateway/dto"

type MediaService interface {
	GetAllMedia() (dto.HttpResponse, error)
	DeleteMediaByID(id string) (dto.HttpResponse, error)
	UploadMediaImage(data dto.UploadMediaBody) (dto.HttpResponse, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}
