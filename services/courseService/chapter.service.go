package courseService

import (
	"encoding/json"
	"fmt"
	"log"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
)

func (s *service) CreateChapter(input dto.ChapterInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("SERVICE-ERR-CChapter 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	data, err := helper.ApiRequest("POST", s.env.CourseServiceUrl, "/chapter", jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-CChapter 2: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetChapter() (dto.HttpResponse, error) {

	data, err := helper.ApiRequest("GET", s.env.CourseServiceUrl, "/chapter", nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GChapter 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetChapterByID(chapterId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/chapter/%s", chapterId)
	data, err := helper.ApiRequest("GET", s.env.CourseServiceUrl, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GChapterBI 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) DeleteChapterByID(chapterId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/chapter/%s", chapterId)
	data, err := helper.ApiRequest("DELETE", s.env.CourseServiceUrl, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-DChapterBI 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) UpdateChapterByID(chapterId string, input dto.UpdateChapterInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("SERVICE-ERR-UChapterBI 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}
	path := fmt.Sprintf("/chapter/%s", chapterId)
	data, err := helper.ApiRequest("PUT", s.env.CourseServiceUrl, path, jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-UChapterBI 2: %v", err.Error())
		return data, err
	}

	return data, nil
}
