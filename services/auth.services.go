package services

import "backend-majoo-test/models"

// AuthService for store auth behaviors
type AuthService interface {
	Login(request models.LoginRequest) (*models.LoginResponse, error)
}
