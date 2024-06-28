package courseHandler

import (
	"encoding/json"
	"net/http"
	"ocApiGateway/dto"
	"ocApiGateway/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCourseHandler(c *gin.Context) {

	var input dto.CourseInputBody
	err := c.ShouldBindJSON(&input)
	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := h.courseService.CreateCourse(input)
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

	var formattedData dto.CourseResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) GetCourseHandler(c *gin.Context) {

	data, err := h.courseService.GetCourse()
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

	var formattedData []dto.CourseResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) GetCourseByIDHandler(c *gin.Context) {

	courseId := c.Param("id")
	data, err := h.courseService.GetCourseByID(courseId)
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

	var formattedData dto.CourseResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) DeleteCourseByIDHandler(c *gin.Context) {

	courseId := c.Param("id")
	data, err := h.courseService.DeleteCourseByID(courseId)
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

	var formattedData dto.CourseResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}

func (h *handler) UpdateCourseByIDHandler(c *gin.Context) {

	courseId := c.Param("id")

	var input dto.UpdateCourseInputBody
	err := c.ShouldBindJSON(&input)
	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := h.courseService.UpdateCourseByID(courseId, input)
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

	var formattedData dto.CourseResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}
