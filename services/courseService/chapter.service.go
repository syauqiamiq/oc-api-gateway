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

	data, err := helper.ApiRequest("POST", COURSE_SERVICE_URL, "/chapter", jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-CChapter 2: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetChapter() (dto.HttpResponse, error) {

	data, err := helper.ApiRequest("GET", COURSE_SERVICE_URL, "/chapter", nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GChapter 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetChapterByID(chapterId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/chapter/%s", chapterId)
	data, err := helper.ApiRequest("GET", COURSE_SERVICE_URL, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GChapterBI 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) DeleteChapterByID(chapterId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/chapter/%s", chapterId)
	data, err := helper.ApiRequest("DELETE", COURSE_SERVICE_URL, path, nil)
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
	data, err := helper.ApiRequest("PUT", COURSE_SERVICE_URL, path, jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-UChapterBI 2: %v", err.Error())
		return data, err
	}

	return data, nil
}
