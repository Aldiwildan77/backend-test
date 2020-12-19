package services

import (
	"backend-majoo-test/entities"
	"backend-majoo-test/models"
	"backend-majoo-test/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type authServiceImpl struct {
	userRepo repositories.UserRepository
}

// NewAuthService to init auth service
func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authServiceImpl{userRepo}
}

func (ar *authServiceImpl) Login(request models.LoginRequest) (*models.LoginResponse, error) {
	user, err := ar.userRepo.FindUserCredentials(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}, nil
}

func generateToken(user *entities.User) (string, error) {
	expTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("rahasiayaahahaha"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
