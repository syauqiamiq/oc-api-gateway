package courseService

import (
	"ocApiGateway/dto"
	"os"

	"github.com/joho/godotenv"
)

type CourseService interface {
	GetMentor() (dto.HttpResponse, error)
	GetMentorByID(mentorId string) (dto.HttpResponse, error)
	CreateMentor(input dto.MentorInputBody) (dto.HttpResponse, error)
	UpdateMentorByID(mentorId string, input dto.UpdateMentorInputBody) (dto.HttpResponse, error)
	DeleteMentorByID(mentorId string) (dto.HttpResponse, error)

	GetCourse() (dto.HttpResponse, error)
	GetCourseByID(courseId string) (dto.HttpResponse, error)
	CreateCourse(input dto.CourseInputBody) (dto.HttpResponse, error)
	UpdateCourseByID(courseId string, input dto.UpdateCourseInputBody) (dto.HttpResponse, error)
	DeleteCourseByID(courseId string) (dto.HttpResponse, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

var (
	_                  = godotenv.Load(".env")
	COURSE_SERVICE_URL = os.Getenv("COURSE_SERVICE_URL")
)
