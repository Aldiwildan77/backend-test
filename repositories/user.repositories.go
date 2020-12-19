package repositories

import "backend-majoo-test/entities"

// UserRepository interface
type UserRepository interface {
	FindAll() (*[]entities.User, error)
}
