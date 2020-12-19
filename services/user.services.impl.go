package services

import (
	"backend-majoo-test/models"
	"backend-majoo-test/repositories"
)

type userServiceImpl struct {
	UserRepo repositories.UserRepository
}

// NewUserService to init user service
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userServiceImpl{userRepo}
}

func (ur *userServiceImpl) FindUserByID(id int) (*models.ReadUserResponse, error) {
	result, err := ur.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := models.ReadUserResponse{
		ID:       result.ID,
		FullName: result.FullName,
		Username: result.Username,
		Photo:    result.Photo,
	}

	return &response, nil
}

func (ur *userServiceImpl) FindAllUser() (*[]models.ReadUserResponse, error) {
	result, err := ur.UserRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.ReadUserResponse
	for _, val := range *result {
		response := models.ReadUserResponse{
			ID:       val.ID,
			FullName: val.FullName,
			Username: val.Username,
			Photo:    val.Photo,
		}

		responses = append(responses, response)
	}

	return &responses, nil
}
