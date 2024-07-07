package services

import (
	"github.com/weinmann-phil/gobank/_common/interfaces"
	"github.com/weinmann-phil/gobank/internals/repository"
)

type UserService interface {
	CreateUserAccount(userRequest *interfaces.UserRegistrationRequest) (*interfaces.UserData, error)
	GetUserAccount(email string) (*interfaces.UserData, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Method to create User
func (us *userService) CreateUserAccount(userRequest *interfaces.UserRegistrationRequest) (*interfaces.UserData, error) {
	userData, err := us.userRepo.CreateUserAccount(userRequest)
	if err != nil {
		return nil, err
	}

	return &interfaces.UserData{
		ID:        userData.ID,
		Username:  userData.Username,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		UserRole:  userData.UserRole,
		CreatedAt: userData.CreatedAt,
	}, nil
}

// Method to GET User
func (us *userService) GetUserAccount(email string) (*interfaces.UserData, error) {
	userData, err := us.userRepo.FetchUserDetails(email)
	if err != nil {
		return nil, err
	}

	return &interfaces.UserData{
		ID:        userData.ID,
		Username:  userData.Username,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		UserRole:  userData.UserRole,
		CreatedAt: userData.CreatedAt,
	}, nil
}
