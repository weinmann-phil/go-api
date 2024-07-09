package repository

import (
	"github.com/weinmann-phil/gobank/_common/interfaces"
	"github.com/weinmann-phil/gobank/internals/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserAccount(userRequest *interfaces.UserRegistrationRequest) (*model.User, error)
	FetchAllUserAccounts() ([]*model.User, error)
	FetchUserDetails(userEmail string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Function to save a user
func (r *userRepository) CreateUserAccount(userRequest *interfaces.UserRegistrationRequest) (*model.User, error) {
	user := &model.User{
		Email:     userRequest.Email,
		Username:  userRequest.Username,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		UserRole:  model.UserRole,
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Method to return all users
func (r *userRepository) FetchAllUserAccounts() ([]*model.User, error) {
	users := []*model.User{}

	if err := r.db.Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Method to return user by email
func (r *userRepository) FetchUserDetails(email string) (*model.User, error) {
	user := &model.User{}

	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
