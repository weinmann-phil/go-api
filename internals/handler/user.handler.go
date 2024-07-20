package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/weinmann-phil/go-api/_common/interfaces"
	"github.com/weinmann-phil/go-api/internals/services"
)

type UserHandler struct {
	userService services.UserService
	validate    *validator.Validate
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validate:    validator.New(),
	}
}

// Get all users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	userData, err := h.userService.GetAllUserAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	c.JSON(http.StatusOK, interfaces.UserResponse{
		Message: "Users retrieved successfully",
		Status:  interfaces.StatusSuccess,
		Code:    http.StatusOK,
		Data:    *userData,
	})
}

// Get a single User by email
func (h *UserHandler) GetSingleUser(c *gin.Context) {
	userMail := c.Param("email")

	if userMail == "" {
		c.JSON(http.StatusBadRequest, interfaces.ErrorMessage{
			Message: "A required field is empty: email",
			Status:  interfaces.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	userData, err := h.userService.GetUserAccount(userMail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	output := []interfaces.UserData{}
	output = append(output, *userData)

	c.JSON(http.StatusOK, interfaces.UserResponse{
		Message: "User retrieved successfully",
		Status:  interfaces.StatusSuccess,
		Code:    http.StatusOK,
		Data:    output,
	})
}

// Create a User
func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest interfaces.UserRegistrationRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfaces.ErrorMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	if err := h.validate.Struct(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfaces.ErrorMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	userData, err := h.userService.CreateUserAccount(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	output := []interfaces.UserData{}
	output = append(output, *userData)

	c.JSON(http.StatusOK, interfaces.UserResponse{
		Message: "User created successfully",
		Status:  interfaces.StatusSuccess,
		Code:    http.StatusOK,
		Data:    output,
	})
}

// Update a User
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var userRequest interfaces.UserRegistrationRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfaces.ErrorMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	if err := h.validate.Struct(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfaces.ErrorMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	userData, err := h.userService.UpdateUserAccount(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	output := []interfaces.UserData{}
	output = append(output, *userData)

	c.JSON(http.StatusOK, interfaces.UserResponse{
		Message: "User updated successfully",
		Status:  interfaces.StatusSuccess,
		Code:    http.StatusOK,
		Data:    output,
	})
}

// Delete a User
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userMail := c.Param("email")

	if userMail == "" {
		c.JSON(http.StatusBadRequest, interfaces.ErrorMessage{
			Message: "A required field is empty: email",
			Status:  interfaces.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	userData, err := h.userService.DeleteUserAccount(userMail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorMessage{
			Message: err.Error(),
			Status:  interfaces.StatusError,
			Code:    http.StatusInternalServerError,
		})

		return
	}

	output := []interfaces.UserData{}
	output = append(output, *userData)

	c.JSON(http.StatusOK, interfaces.UserResponse{
		Message: "User deleted successfully",
		Status:  interfaces.StatusSuccess,
		Code:    http.StatusOK,
		Data:    output,
	})
}
