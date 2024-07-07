package interfaces

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/weinmann-phil/gobank/internals/model"
)

type ResponsesStatus string

const (
	StatusSuccess ResponsesStatus = "success"
	StatusError   ResponsesStatus = "error"
)

type UserRegistrationRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"user" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type UserData struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json: "user"`
	FirstName string     `json: "firstName"`
	LastName  string     `json: "lastName"`
	UserRole  model.Role `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
}

type ErrorMessage struct {
	Message string          `json:"message"`
	Code    int             `json:"code"`
	Status  ResponsesStatus `json:"status"`
}

type UserResponse struct {
	Message string          `json:"message"`
	Code    int             `json:"code"`
	Status  ResponsesStatus `json:"status"`
	Data    UserData        `json:"data"`
}

type RouteDefinition struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}
