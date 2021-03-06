package repositories

import (
	"backend-majoo-test/entities"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepository to init user repository
func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepositoryImpl{DB}
}

func (r *userRepositoryImpl) Create(user entities.User) (*entities.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) FindByID(id int) (*entities.User, error) {
	var user entities.User

	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) FindAll() (*[]entities.User, error) {
	var users []entities.User

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *userRepositoryImpl) FindUserCredentials(username string, password string) (*entities.User, error) {
	var user entities.User

	if err := r.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) Update(id int, userRequest entities.User) (*entities.User, error) {
	var user entities.User

	if err := r.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}

	user.Password = userRequest.Password
	user.FullName = userRequest.FullName

	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) Delete(id int) error {
	if err := r.DB.Delete(&entities.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryImpl) UpdatePhoto(id int, location string) error {
	if err := r.DB.Model(&entities.User{}).Where("id", id).Update("photo", location).Error; err != nil {
		return err
	}

	return nil
}
