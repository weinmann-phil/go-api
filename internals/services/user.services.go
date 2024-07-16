package services

import (
	"github.com/sirupsen/logrus"
	"github.com/weinmann-phil/gobank/_common/interfaces"
	"github.com/weinmann-phil/gobank/internals/repository"
)

type UserService interface {
	CreateUserAccount(userRequest *interfaces.UserRegistrationRequest) (*interfaces.UserData, error)
	GetUserAccount(email string) (*interfaces.UserData, error)
	GetAllUserAccounts() (*[]interfaces.UserData, error)
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

// Method to GET all Users
func (us *userService) GetAllUserAccounts() (*[]interfaces.UserData, error) {
	userData, err := us.userRepo.FetchAllUserAccounts()
	if err != nil {
		return nil, err
	}

	for k, v := range *userData {
		logrus.Infof("%d, %v", k, v)
	}

	output := []interfaces.UserData{}

	for _, data := range *userData {
		output = append(output, interfaces.UserData{
			ID:        data.ID,
			Username:  data.Username,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			UserRole:  data.UserRole,
			CreatedAt: data.CreatedAt,
		})
	}

	return &output, nil
}

// Method to GET a single User
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
