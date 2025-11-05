package entity

import "gorm.io/gorm"

type Student struct{
	gorm.Model
	FirstName string `json:"first_name"`
	LastName string  `json:"last_name"`
	Age string `json:"age"`
	Email string `json:"email"`
}