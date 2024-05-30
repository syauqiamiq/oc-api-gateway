package mediaHandler

import "ocApiGateway/services/mediaService"

type handler struct {
	mediaService mediaService.MediaService
}

func NewHandler(mediaService mediaService.MediaService) *handler {
	return &handler{
		mediaService: mediaService,
	}

}
