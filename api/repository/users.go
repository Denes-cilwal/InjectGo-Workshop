package repository

import (
	"InjectGo-Workshop/model"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	CreateUser(user *model.User) (User *model.User, err error)
	GetUsers() (users *[]model.User, err error)
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
