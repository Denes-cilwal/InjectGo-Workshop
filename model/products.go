package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	ProductCode string `json:"product_code"`
	UserID      uint   // Foreign key referencing User's primary key (ID)
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Define cascading behavior
}

func (p Product) String() string {
	return "products"
}
