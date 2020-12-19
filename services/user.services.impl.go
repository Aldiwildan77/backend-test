package services

import (
	"backend-majoo-test/entities"
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

func (ur *userServiceImpl) CreateNewUser(request models.CreateUserRequest) (*models.CreateUserResponse, error) {
	user := entities.User{
		Username: request.Username,
		Password: request.Password,
		FullName: request.FullName,
	}

	result, err := ur.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	response := models.CreateUserResponse{
		ID:       result.ID,
		FullName: result.FullName,
		Username: result.Username,
	}

	return &response, nil
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

func (ur *userServiceImpl) UpdateUser(id int, request models.UpdateUserRequest) (*models.UpdateUserResponse, error) {
	user, err := ur.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	userData := prepareUpdate(user, request)

	result, err := ur.UserRepo.Update(id, userData)
	if err != nil {
		return nil, err
	}

	response := models.UpdateUserResponse{
		ID:       result.ID,
		Username: result.Username,
		FullName: result.FullName,
		Photo:    result.Photo,
	}

	return &response, nil
}

func (ur *userServiceImpl) DeleteUserByID(id int) error {
	user, err := ur.UserRepo.FindByID(id)
	if err != nil {
		return err
	}

	return ur.UserRepo.Delete(int(user.ID))
}

func (ur *userServiceImpl) UploadPhoto(id int, location string) (*models.UpdateUserResponse, error) {
	if err := ur.UserRepo.UpdatePhoto(id, location); err != nil {
		return nil, err
	}

	result, err := ur.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := models.UpdateUserResponse{
		ID:       result.ID,
		FullName: result.FullName,
		Username: result.Username,
		Photo:    result.Photo,
	}

	return &response, nil
}

func prepareUpdate(user *entities.User, request models.UpdateUserRequest) entities.User {
	if request.FullName != "" {
		user.FullName = request.FullName
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	return *user
}
