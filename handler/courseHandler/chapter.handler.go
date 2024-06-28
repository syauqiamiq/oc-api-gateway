package courseHandler

import (
	"encoding/json"
	"net/http"
	"ocApiGateway/dto"
	"ocApiGateway/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateChapterHandler(c *gin.Context) {

	var input dto.ChapterInputBody
	err := c.ShouldBindJSON(&input)
	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := h.courseService.CreateChapter(input)
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

	var formattedData dto.ChapterResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) GetChapterHandler(c *gin.Context) {

	data, err := h.courseService.GetChapter()
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

	var formattedData []dto.ChapterResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) GetChapterByIDHandler(c *gin.Context) {

	chapterId := c.Param("id")
	data, err := h.courseService.GetChapterByID(chapterId)
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

	var formattedData dto.ChapterResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) DeleteChapterByIDHandler(c *gin.Context) {

	chapterId := c.Param("id")
	data, err := h.courseService.DeleteChapterByID(chapterId)
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

	var formattedData dto.ChapterResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) UpdateChapterByIDHandler(c *gin.Context) {

	chapterId := c.Param("id")

	var input dto.UpdateChapterInputBody
	err := c.ShouldBindJSON(&input)
	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := h.courseService.UpdateChapterByID(chapterId, input)
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

	var formattedData dto.ChapterResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}
