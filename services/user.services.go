package services

import "backend-majoo-test/models"

// UserService for store user behaviors
type UserService interface {
	FindUserByID(id int) (*models.ReadUserResponse, error)
	FindAllUser() (*[]models.ReadUserResponse, error)
}
