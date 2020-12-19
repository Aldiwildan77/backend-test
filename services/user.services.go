package services

import "backend-majoo-test/models"

// UserService for store user behaviors
type UserService interface {
	FindAllUser() (*[]models.ReadUserResponse, error)
}
