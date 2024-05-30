package userHandler

import (
	"net/http"
	"ocApiGateway/dto"
	"ocApiGateway/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) RefreshTokenHandler(c *gin.Context) {

	var input dto.RefreshTokenInputBody
	err := c.ShouldBindJSON(&input)

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	validateRefreshTokenResponse, err := h.userService.ValidateRefreshToken(input.RefreshToken)
	if err != nil {

		response := helper.APIResponse(http.StatusServiceUnavailable, "Service Unavailable", nil)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	if validateRefreshTokenResponse.Status == "error" {
		response := helper.APIResponse(validateRefreshTokenResponse.Code, validateRefreshTokenResponse.Message, nil)
		c.JSON(validateRefreshTokenResponse.Code, response)
		return
	}

	newAccessToken, err := h.userService.GenerateNewAccessToken(input)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	formattedResponse := dto.RefreshTokenResponse{
		AccessToken: newAccessToken,
	}
	response := helper.APIResponse(http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}
