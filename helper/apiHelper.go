package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ocApiGateway/dto"
)

func ApiRequest(method string, baseUrl string, apiRoute string, body []byte) (dto.HttpResponse, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", baseUrl, apiRoute), bytes.NewBuffer(body))
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
