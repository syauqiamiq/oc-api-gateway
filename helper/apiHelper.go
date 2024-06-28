package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"ocApiGateway/dto"
)

func ApiRequest(method string, baseUrl string, apiRoute string, body []byte) (dto.HttpResponse, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", baseUrl, apiRoute), bytes.NewBuffer(body))
	if err != nil {
		return dto.HttpResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	logEntry, _ := json.Marshal(map[string]string{
		"URL":     baseUrl + apiRoute,
		"METHOD":  method,
		"BODY":    string(body),
		"Message": "Sending HTTP request",
	})
	log.Println(string(logEntry))

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
