package courseHandler

import (
	"encoding/json"
	"net/http"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
	"ocApiGateway/middleware"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetMyCourseHandler(c *gin.Context) {
	accessData := middleware.GetSessionAccessData(c)

	data, err := h.courseService.GetMyCourse(accessData.UserID)
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

	var formattedData []dto.MyCourseResponse
	err = json.Unmarshal(jsonData, &formattedData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, formattedData)
	c.JSON(data.Code, response)
}
