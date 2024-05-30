package userHandler

import (
	"encoding/json"
	"net/http"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
	"ocApiGateway/middleware"

	"github.com/gin-gonic/gin"
)

func (h *handler) RegisterHandler(c *gin.Context) {

	var input dto.RegisterInputBody
	err := c.ShouldBindJSON(&input)

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data, err := h.userService.Register(input)
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

	jsonUserData, err := json.Marshal(data.Data)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	var userData dto.User
	err = json.Unmarshal(jsonUserData, &userData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, userData)
	c.JSON(data.Code, response)
}

func (h *handler) UpdateUserHandler(c *gin.Context) {
	accessData := middleware.GetSessionAccessData(c)
	var input dto.UpdateUserInputBody
	err := c.ShouldBindJSON(&input)

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data, err := h.userService.UpdateProfile(accessData.UserID, input)
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

	jsonUserData, err := json.Marshal(data.Data)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	var userData dto.User
	err = json.Unmarshal(jsonUserData, &userData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, userData)
	c.JSON(data.Code, response)
}

func (h *handler) GetProfileHandler(c *gin.Context) {
	accessData := middleware.GetSessionAccessData(c)
	data, err := h.userService.GetProfile(accessData.UserID)
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

	jsonUserData, err := json.Marshal(data.Data)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	var userData dto.User
	err = json.Unmarshal(jsonUserData, &userData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	response := helper.APIResponse(data.Code, data.Message, userData)
	c.JSON(data.Code, response)
}

func (h *handler) LogoutHandler(c *gin.Context) {

	accessData := middleware.GetSessionAccessData(c)

	input := dto.LogoutBody{
		UserID: accessData.UserID,
	}

	data, err := h.userService.Logout(input)
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

	jsonUserData, err := json.Marshal(data.Data)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	var userData dto.User
	err = json.Unmarshal(jsonUserData, &userData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}
	response := helper.APIResponse(data.Code, data.Message, userData)
	c.JSON(data.Code, response)
}

func (h *handler) LoginHandler(c *gin.Context) {

	var input dto.LoginInputBody
	err := c.ShouldBindJSON(&input)

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse(http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginResponse, err := h.userService.Login(input)
	if err != nil {

		response := helper.APIResponse(http.StatusServiceUnavailable, "Service Unavailable", nil)
		c.JSON(response.Code, response)
		return
	}
	if loginResponse.Status == "error" {
		response := helper.APIResponse(loginResponse.Code, loginResponse.Message, nil)
		c.JSON(loginResponse.Code, response)
		return
	}

	jsonUserData, err := json.Marshal(loginResponse.Data)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	var userData dto.User
	err = json.Unmarshal(jsonUserData, &userData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	token, refreshToken, err := h.userService.GenerateJWTToken(userData)
	if err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, err.Error(), nil)
		c.JSON(response.Code, response)
		return
	}

	saveRefreshTokenRespnse, err := h.userService.SaveRefreshToken(userData.ID, refreshToken)
	if err != nil {
		response := helper.APIResponse(http.StatusServiceUnavailable, "Service Unavailable", nil)
		c.JSON(response.Code, response)
		return
	}

	if saveRefreshTokenRespnse.Status == "error" {
		response := helper.APIResponse(saveRefreshTokenRespnse.Code, saveRefreshTokenRespnse.Message, nil)
		c.JSON(saveRefreshTokenRespnse.Code, response)
		return
	}

	formattedResponse := dto.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	response := helper.APIResponse(http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}
