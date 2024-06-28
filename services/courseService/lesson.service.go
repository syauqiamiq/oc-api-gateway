package courseService

import (
	"encoding/json"
	"fmt"
	"log"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
)

func (s *service) CreateLesson(input dto.LessonInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("SERVICE-ERR-CL 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	data, err := helper.ApiRequest("POST", COURSE_SERVICE_URL, "/lesson", jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-CL 2: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetLesson() (dto.HttpResponse, error) {

	data, err := helper.ApiRequest("GET", COURSE_SERVICE_URL, "/lesson", nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GL 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetLessonByID(lessonId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/lesson/%s", lessonId)
	data, err := helper.ApiRequest("GET", COURSE_SERVICE_URL, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GLBI 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) DeleteLessonByID(lessonId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/lesson/%s", lessonId)
	data, err := helper.ApiRequest("DELETE", COURSE_SERVICE_URL, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-DLBI 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) UpdateLessonByID(lessonId string, input dto.UpdateLessonInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("SERVICE-ERR-ULBI 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}
	path := fmt.Sprintf("/lesson/%s", lessonId)
	data, err := helper.ApiRequest("PUT", COURSE_SERVICE_URL, path, jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-ULBI 2: %v", err.Error())
		return data, err
	}

	return data, nil
}
