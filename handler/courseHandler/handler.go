package courseHandler

import (
	"ocApiGateway/services/courseService"
)

type handler struct {
	courseService courseService.CourseService
}

func NewHandler(courseService courseService.CourseService) *handler {
	return &handler{
		courseService: courseService,
	}

}
