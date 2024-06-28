package courseService

import (
	"encoding/json"
	"fmt"
	"log"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
)

func (s *service) CreateMentor(input dto.MentorInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("SERVICE-ERR-CM 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	data, err := helper.ApiRequest("POST", COURSE_SERVICE_URL, "/mentor", jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-CM 2: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetMentor() (dto.HttpResponse, error) {

	data, err := helper.ApiRequest("GET", COURSE_SERVICE_URL, "/mentor", nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GM 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) GetMentorByID(mentorId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/mentor/%s", mentorId)
	data, err := helper.ApiRequest("GET", COURSE_SERVICE_URL, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GMBI 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) DeleteMentorByID(mentorId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/mentor/%s", mentorId)
	data, err := helper.ApiRequest("DELETE", COURSE_SERVICE_URL, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-DMBI 1: %v", err.Error())
		return data, err
	}

	return data, nil
}

func (s *service) UpdateMentorByID(mentorId string, input dto.UpdateMentorInputBody) (dto.HttpResponse, error) {

	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("SERVICE-ERR-UMBI 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}
	path := fmt.Sprintf("/mentor/%s", mentorId)
	data, err := helper.ApiRequest("PUT", COURSE_SERVICE_URL, path, jsonData)
	if err != nil {
		log.Printf("SERVICE-ERR-UMBI 2: %v", err.Error())
		return data, err
	}

	return data, nil
}
