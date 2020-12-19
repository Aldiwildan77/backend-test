package repositories

import "backend-majoo-test/entities"

// UserRepository interface
type UserRepository interface {
	Create(user entities.User) (*entities.User, error)
	FindByID(id int) (*entities.User, error)
	FindAll() (*[]entities.User, error)
}
