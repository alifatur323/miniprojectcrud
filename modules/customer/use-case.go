package customer

import (
	"crmservice/entities"
	"crmservice/repositories"
	"time"
)

type UseCaseCustomer interface {
	CreateCustomer(customer CustomerParam) (entities.Customer, error)
	GetCustomerById(id uint) (entities.Customer, error)
	UpdateCustomer(customer CustomerParam, id uint) (any, error)
	DeleteCustomer(id uint) (any, error)
}
type useCaseCustomer struct {
	customerRepo repositories.CustomerInterfaceRepo
}

func (uc useCaseCustomer) CreateCustomer(customer CustomerParam) (entities.Customer, error) {
	var newCustomer *entities.Customer
	newCustomer = &entities.Customer{
		First_name: customer.First_name,
		Last_name:  customer.Last_name,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	_, err := uc.customerRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}
	return *newCustomer, nil
}
func (uc useCaseCustomer) GetCustomerById(id uint) (entities.Customer, error) {
	var customer entities.Customer
	customer, err := uc.customerRepo.GetCustomerById(id)
	return customer, err
}

func (uc useCaseCustomer) UpdateCustomer(customer CustomerParam, id uint) (any, error) {
	var editCustomer *entities.Customer
	editCustomer = &entities.Customer{
		Id:         id,
		First_name: customer.First_name,
		Last_name:  customer.Last_name,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	_, err := uc.customerRepo.UpdateCustomer(editCustomer)
	if err != nil {
		return *editCustomer, err
	}
	return *editCustomer, nil
}

func (uc useCaseCustomer) DeleteCustomer(id uint) (any, error) {
	_, err := uc.customerRepo.DeleteCustomer(id)
	return nil, err
}
