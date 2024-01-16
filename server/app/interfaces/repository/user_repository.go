package repository

import (
	"go-restapi-template/app/domain/entity"
	"go-restapi-template/db/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() (*model.Users, error)
	GetByID(user entity.User) (*model.User, error)
	Create(user entity.User) (*model.User, error)
	Update(user entity.User) (*model.User, error)
	Delete(user entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetAll() (*model.Users, error) {
	var usersData model.Users
	if err := ur.db.Order("id asc").Find(&usersData).Error; err != nil {
		return nil, err
	}
	return &usersData, nil
} 
	
func (ur *userRepository) GetByID(user entity.User) (*model.User, error) {
	var userData model.User
	if err := ur.db.First(&userData, *user.ID).Error; err != nil {
		return nil, err
	}

	return &userData, nil
}

func (ur *userRepository) Create(user entity.User) (*model.User, error) {
	userData := model.User{
		FirstName: *user.FirstName,
		LastName: *user.LastName,
	}

	if err := ur.db.Create(&userData).Error; err != nil {
		return nil, err
	}

	return &userData, nil
}

func (ur *userRepository) Update(user entity.User) (*model.User, error) {
	var userData model.User

	if user.FirstName != nil {
		userData.FirstName = *user.FirstName
	}

	if user.LastName != nil {
		userData.LastName = *user.LastName
	}

	err := ur.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", *user.ID).Updates(&userData).Error; err != nil {
			return err
		}

		tx.First(&userData, *user.ID)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &userData, nil
}

func (ur *userRepository) Delete(user entity.User) error {
	var userData model.User
	err := ur.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&userData, user.ID).Error; err != nil {
			return err
		}
			
		if err := ur.db.Delete(&model.User{}, user.ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}