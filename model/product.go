package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name     string
	Category Category
}
