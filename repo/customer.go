package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/ptloc1992/boilerplate-microservice/model"
)

type CustomerRepo interface {
	Create(model.Customer) (*model.Customer, error)
	Find(uint) (*model.Customer, error)
}

type CustomerRepoImpl struct {
	DB *gorm.DB
}

func (customerRepo CustomerRepoImpl) Create(customer model.Customer) (*model.Customer, error) {
	err := customerRepo.DB.Create(&customer).Error
	return &customer, err
}

// Find a customer
func (customerRepo CustomerRepoImpl) Find(id uint) (*model.Customer, error) {
	customer := &model.Customer{}
	err := customerRepo.DB.Find(customer, id).Error
	return customer, err
}
