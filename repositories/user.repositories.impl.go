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
