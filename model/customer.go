package model

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name    string
	Phone   string
	Address string
}
