package courseHandler

import (
	"encoding/json"
	"net/http"
	"ocApiGateway/dto"
	"ocApiGateway/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateMentorHandler(c *gin.Context) {

	var input dto.MentorInputBody
	err := c.ShouldBindJSON(&input)
	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := h.courseService.CreateMentor(input)
	if err != nil {

		response := helper.APIResponse(http.StatusServiceUnavailable, "Service Unavailable", nil)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if data.Status == "error" {
		response := helper.APIResponse(data.Code, data.Message, nil)
		c.JSON(data.Code, response)
		return
	}

	jsonData, err := json.Marshal(data.Data)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	var formattedData dto.MentorResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) GetMentorHandler(c *gin.Context) {

	data, err := h.courseService.GetMentor()
	if err != nil {

		response := helper.APIResponse(http.StatusServiceUnavailable, "Service Unavailable", nil)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if data.Status == "error" {
		response := helper.APIResponse(data.Code, data.Message, nil)
		c.JSON(data.Code, response)
		return
	}

	jsonData, err := json.Marshal(data.Data)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	var formattedData []dto.MentorResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) GetMentorByIDHandler(c *gin.Context) {

	mentorId := c.Param("id")
	data, err := h.courseService.GetMentorByID(mentorId)
	if err != nil {

		response := helper.APIResponse(http.StatusServiceUnavailable, "Service Unavailable", nil)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if data.Status == "error" {
		response := helper.APIResponse(data.Code, data.Message, nil)
		c.JSON(data.Code, response)
		return
	}

	jsonData, err := json.Marshal(data.Data)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	var formattedData dto.MentorResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) DeleteMentorByIDHandler(c *gin.Context) {

	mentorId := c.Param("id")
	data, err := h.courseService.DeleteMentorByID(mentorId)
	if err != nil {

		response := helper.APIResponse(http.StatusServiceUnavailable, "Service Unavailable", nil)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if data.Status == "error" {
		response := helper.APIResponse(data.Code, data.Message, nil)
		c.JSON(data.Code, response)
		return
	}

	jsonData, err := json.Marshal(data.Data)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	var formattedData dto.MentorResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) UpdateMentorByIDHandler(c *gin.Context) {

	mentorId := c.Param("id")

	var input dto.UpdateMentorInputBody
	err := c.ShouldBindJSON(&input)
	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := h.courseService.UpdateMentorByID(mentorId, input)
	if err != nil {

		response := helper.APIResponse(http.StatusServiceUnavailable, "Service Unavailable", nil)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if data.Status == "error" {
		response := helper.APIResponse(data.Code, data.Message, nil)
		c.JSON(data.Code, response)
		return
	}

	jsonData, err := json.Marshal(data.Data)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	var formattedData dto.MentorResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}
