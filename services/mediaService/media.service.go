package mediaService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"ocApiGateway/dto"
	"os"

	"github.com/joho/godotenv"
)

var (
	_        = godotenv.Load(".env")
	BASE_URL = os.Getenv("MEDIA_SERVICE_URL")
)

func (s *service) GetAllMedia() (dto.HttpResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/media", BASE_URL), nil)
	if err != nil {
		return dto.HttpResponse{}, err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return dto.HttpResponse{}, err
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return dto.HttpResponse{}, err
	}

	var formattedResponse dto.HttpResponse

	err = json.Unmarshal(response, &formattedResponse)
	if err != nil {
		return dto.HttpResponse{}, err
	}

	return formattedResponse, nil
}

func (s *service) DeleteMediaByID(id string) (dto.HttpResponse, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/media/%s", BASE_URL, id), nil)
	if err != nil {
		log.Printf("ERR 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("ERR 2: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("ERR 3: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	fmt.Println(string(response))
	var formattedResponse dto.HttpResponse

	err = json.Unmarshal(response, &formattedResponse)
	if err != nil {
		log.Printf("ERR 4: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	return formattedResponse, nil
}

func (s *service) UploadMediaImage(data dto.UploadMediaBody) (dto.HttpResponse, error) {

	payload := dto.UploadMediaBody{
		Image: data.Image,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("ERR 1: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/media", BASE_URL), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("ERR 2: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("ERR 3: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("ERR 4: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	var formattedResponse dto.HttpResponse

	err = json.Unmarshal(response, &formattedResponse)
	if err != nil {
		log.Printf("ERR 5: %v", err.Error())
		return dto.HttpResponse{}, err
	}

	return formattedResponse, nil
}
