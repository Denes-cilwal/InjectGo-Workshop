package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint   `json:"age"`
}

func (u User) String() string {
	return "users"
}
