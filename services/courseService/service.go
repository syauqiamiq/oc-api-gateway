package courseService

import (
	"ocApiGateway/dto"
	"ocApiGateway/helper"
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

	GetChapter() (dto.HttpResponse, error)
	GetChapterByID(chapterId string) (dto.HttpResponse, error)
	CreateChapter(input dto.ChapterInputBody) (dto.HttpResponse, error)
	UpdateChapterByID(chapterId string, input dto.UpdateChapterInputBody) (dto.HttpResponse, error)
	DeleteChapterByID(chapterId string) (dto.HttpResponse, error)

	GetLesson() (dto.HttpResponse, error)
	GetLessonByID(lessonId string) (dto.HttpResponse, error)
	CreateLesson(input dto.LessonInputBody) (dto.HttpResponse, error)
	UpdateLessonByID(lessonId string, input dto.UpdateLessonInputBody) (dto.HttpResponse, error)
	DeleteLessonByID(lessonId string) (dto.HttpResponse, error)

	GetMyCourse(userId string) (dto.HttpResponse, error)
}

type service struct {
	env helper.Env
}

func NewService(env helper.Env) *service {
	return &service{
		env: env,
	}
}
