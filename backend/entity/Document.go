package entity

import "gorm.io/gorm"

type Document struct{
	gorm.Model
	Title string `json:"title"`
	File string  `json:"file"`
}