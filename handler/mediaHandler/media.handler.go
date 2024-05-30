package mediaHandler

import (
	"net/http"
	"ocApiGateway/dto"
	"ocApiGateway/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAllMediaHandler(c *gin.Context) {

	data, err := h.mediaService.GetAllMedia()
	if err != nil {

		response := helper.APIResponse(http.StatusServiceUnavailable, "Service Error", nil)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if data.Status == "error" {
		response := helper.APIResponse(data.Code, data.Message, nil)
		c.JSON(data.Code, response)
		return
	}
	response := helper.APIResponse(data.Code, data.Message, data.Data)
	c.JSON(data.Code, response)
}

func (h *handler) DeleteMediaByIdHandler(c *gin.Context) {

	id := c.Param("id")

	data, err := h.mediaService.DeleteMediaByID(id)
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
	response := helper.APIResponse(data.Code, data.Message, data.Data)
	c.JSON(data.Code, response)
}

func (h *handler) UploadImageHandler(c *gin.Context) {

	var input dto.UploadMediaBody
	err := c.ShouldBindJSON(&input)

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data, err := h.mediaService.UploadMediaImage(input)
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

	response := helper.APIResponse(data.Code, data.Message, data.Data)
	c.JSON(data.Code, response)
}
