package courseService

import (
	"fmt"
	"log"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
)

func (s *service) GetMyCourse(userId string) (dto.HttpResponse, error) {
	path := fmt.Sprintf("/my-course?user_id=%s", userId)
	data, err := helper.ApiRequest("GET", COURSE_SERVICE_URL, path, nil)
	if err != nil {
		log.Printf("SERVICE-ERR-GMC 1: %v", err.Error())
		return data, err
	}

	return data, nil
}
