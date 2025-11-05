package entity

import "gorm.io/gorm"

type Something struct{
	gorm.Model
	Title string `json:"title"`
	File string  `json:"file"`
}