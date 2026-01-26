package services

import (
	"errors"
	"ginLearning/05_Auth/models"
	"ginLearning/05_Auth/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(db *gorm.DB) *AuthService {
	db.AutoMigrate(&models.User{})
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) LoginService(email *string, password *string) (*models.User, error) {
	if email == nil {
		return nil, errors.New("Email can't be null")
	}
	if password == nil {
		return nil, errors.New("Password can't be null")
	}

	var user models.User

	if err := a.db.Where("email = ?", *email).First(&user).Error; err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(*password, user.Password) {
		return nil, errors.New("Invalid credentials")
	}

	return &user, nil
}

func (a *AuthService) RegisterService(email *string, password *string, name *string) (*models.User, error) {
	if email == nil {
		return nil, errors.New("Email can't be null")
	}
	if password == nil {
		return nil, errors.New("Password can't be null")
	}
	if name == nil {
		return nil, errors.New("Name can't be null")
	}

	hashedPassword, err := utils.HashPassword(*password)
	if err != nil {
		return nil, err
	}

	var user models.User

	user.Name = *name
	user.Email = *email
	user.Password = hashedPassword

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
