package repositories

import "backend-majoo-test/entities"

// UserRepository interface
type UserRepository interface {
	FindByID(id int) (*entities.User, error)
	FindAll() (*[]entities.User, error)
}
