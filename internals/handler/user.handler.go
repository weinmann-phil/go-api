package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/weinmann-phil/gobank/_common/interfaces"
	"github.com/weinmann-phil/gobank/internals/services"
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

// func (h *UserHandler) ReadUser(c *gin.Context) {
// 	var userRequest interfaces.UserResponse
// 	if err := c.ShouldBindJSON(&userRequest); err != nil {
// 		c.JSON(http.StatusBadRequest, interfaces.ErrorMessage{
// 			Message: err.Error(),
// 			Status:  interfaces.StatusError,
// 			Code:    http.StatusBadRequest,
// 		})

// 		return
// 	}

// 	userData, err := h.userService.GetUserAccount()
// }

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

	c.JSON(http.StatusOK, interfaces.UserResponse{
		Message: "User created successfully",
		Status:  interfaces.StatusSuccess,
		Code:    http.StatusOK,
		Data:    *userData,
	})
}
