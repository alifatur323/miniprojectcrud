package repositories

import (
	"crmservice/entities"
	"gorm.io/gorm"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{
		db: dbCrud,
	}
}

type CustomerInterfaceRepo interface {
	CreateCustomer(customer *entities.Customer) (*entities.Customer, error)
	GetCustomerById(id uint) (entities.Customer, error)
	UpdateCustomer(customer *entities.Customer) (any, error)
	DeleteCustomer(id uint) (any, error)
}

func (repo Customer) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Create(customer).Error
	return customer, err
}

func (repo Customer) GetCustomerById(id uint) (entities.Customer, error) {
	var customer entities.Customer
	repo.db.First(&customer, "id = ? ", id)
	return customer, nil
}

func (repo Customer) UpdateCustomer(customer *entities.Customer) (any, error) {
	err := repo.db.Save(customer).Error
	return nil, err
}

func (repo Customer) DeleteCustomer(id uint) (any, error) {
	err := repo.db.Model(&entities.Customer{}).
		Where("id = ?", id).
		Delete(&entities.Customer{}).
		Error
	return nil, err
}
