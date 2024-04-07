package handler

import (
	"net/http"
	"testBW/helper"
	"testBW/users"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(uhuy users.Service) *userHandler {
	return &userHandler{uhuy}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ApiErrorResponse(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusUnprocessableEntity, helper.ApiResponse("Failed register", http.StatusUnprocessableEntity, "Failed", errorMessage))
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Failed register", http.StatusBadRequest, "Failed", nil))
		return
	}

	newFormat := users.UserFormatter(newUser)

	response := helper.ApiResponse("Account Registered", http.StatusOK, "Success", newFormat)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input users.LoginIniput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ApiErrorResponse(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusUnprocessableEntity, helper.ApiResponse("Failed email", http.StatusUnprocessableEntity, "Failed", errorMessage))
		return
	}

	user, err := h.userService.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Failed email", http.StatusBadRequest, "Failed", nil))
		return
	}

	newFormat := users.UserFormatter(user)

	response := helper.ApiResponse("Account found", http.StatusOK, "Success", newFormat)

	c.JSON(http.StatusOK, response)
}
