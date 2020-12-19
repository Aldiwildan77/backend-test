package services

import "backend-majoo-test/models"

// UserService for store user behaviors
type UserService interface {
	CreateNewUser(request models.CreateUserRequest) (*models.CreateUserResponse, error)
	FindUserByID(id int) (*models.ReadUserResponse, error)
	FindAllUser() (*[]models.ReadUserResponse, error)
	UpdateUser(id int, request models.UpdateUserRequest) (*models.UpdateUserResponse, error)
	DeleteUserByID(id int) error
}
