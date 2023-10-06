package repository

import (
	"InjectGo-Workshop/model"
	"gorm.io/gorm"
	"log"
)

type ProductRepository interface {
	Migrate() error
}

type IProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &IProductRepository{
		DB: db,
	}
}

func (u *IProductRepository) Migrate() error {
	log.Print("[IProductRepository]...Migrate")
	return u.DB.AutoMigrate(&model.Product{})
}
