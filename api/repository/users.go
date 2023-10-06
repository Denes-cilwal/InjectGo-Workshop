package repository

import (
	"InjectGo-Workshop/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	CreateUser(user *model.User) (User *model.User, err error)
	GetUsers() (users *[]model.User, err error)
	GetUserByID(userID uint) (*model.User, error)
	Migrate() error
}

type IUserRepository struct {
	DB *gorm.DB
}

// NewUserRepository -> returns new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &IUserRepository{
		DB: db,
	}
}

func (u *IUserRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&model.User{})
}

func (u *IUserRepository) CreateUser(user *model.User) (*model.User, error) {
	if err := u.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *IUserRepository) GetUsers() (*[]model.User, error) {
	var users []model.User
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *IUserRepository) GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	if err := u.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
