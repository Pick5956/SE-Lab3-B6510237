package entity

import "gorm.io/gorm"

type Template struct{
	gorm.Model
	FieldName string `json:"title"`
	FieldType string  `json:"file"`
	BoxPositionx string `json:"box_positionx"`
}