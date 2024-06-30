package courseService

import (
	"encoding/json"
	"fmt"
	"log"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
)

func (s *service) CreateCourse(input dto.CourseInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("SERVICE-ERR-CC 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	data, err := helper.ApiRequest("POST", s.env.CourseServiceUrl, "/course", jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-CC 2: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetCourse() (dto.HttpResponse, error) {

	data, err := helper.ApiRequest("GET", s.env.CourseServiceUrl, "/course", nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GC 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetCourseByID(courseId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/course/%s", courseId)
	data, err := helper.ApiRequest("GET", s.env.CourseServiceUrl, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GCBI 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) DeleteCourseByID(courseId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/course/%s", courseId)
	data, err := helper.ApiRequest("DELETE", s.env.CourseServiceUrl, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-DCBI 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) UpdateCourseByID(courseId string, input dto.UpdateCourseInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("SERVICE-ERR-UCBI 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}
	path := fmt.Sprintf("/course/%s", courseId)
	data, err := helper.ApiRequest("PUT", s.env.CourseServiceUrl, path, jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-UCBI 2: %v", err.Error())
		return data, err
	}

	return data, nil
}
