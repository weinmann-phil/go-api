package repository

import (
	"github.com/weinmann-phil/go-api/_common/interfaces"
	"github.com/weinmann-phil/go-api/internals/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserAccount(userRequest *interfaces.UserRegistrationRequest) (*model.User, error)
	UpdateUserAccount(userRequest *interfaces.UserRegistrationRequest) (*model.User, error)
	FetchAllUserAccounts() (*[]model.User, error)
	FetchUserDetails(userEmail string) (*model.User, error)
	DeleteUserAccount(userEmail string) (*model.User, error)
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
func (r *userRepository) FetchAllUserAccounts() (*[]model.User, error) {
	users := &[]model.User{}

	if err := r.db.Find(&users).Error; err != nil {
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

// Method to update user by email
func (r *userRepository) UpdateUserAccount(userRequest *interfaces.UserRegistrationRequest) (*model.User, error) {
	user, _ := r.FetchUserDetails(userRequest.Email)

	user.Username = userRequest.Username
	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName

	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Method to delete user record
func (r *userRepository) DeleteUserAccount(email string) (*model.User, error) {
	user, err := r.FetchUserDetails(email)
	if err != nil {
		return nil, err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
