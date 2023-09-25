package service

import (
	"go_kit_inventory/internal/domain/requests"
	"go_kit_inventory/internal/models"
	"go_kit_inventory/internal/repository"
)

type ServiceCustomer struct {
	Repository repository.CustomerRepository
}

func NewServiceCustomer(customer repository.CustomerRepository) *ServiceCustomer {
	return &ServiceCustomer{Repository: customer}
}

func (s *ServiceCustomer) Create(input *requests.CreateCustomerRequest) (*models.ModelCustomer, error) {
	var customer requests.CreateCustomerRequest

	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	res, err := s.Repository.Create(&customer)

	return res, err
}

func (s *ServiceCustomer) Results() (*[]models.ModelCustomer, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *ServiceCustomer) Result(id string) (*models.ModelCustomer, error) {

	res, err := s.Repository.Result(id)

	return res, err
}

func (s *ServiceCustomer) Delete(id string) (*models.ModelCustomer, error) {

	res, err := s.Repository.Delete(id)

	return res, err
}

func (s *ServiceCustomer) Update(input *requests.UpdateCustomerRequest) (*models.ModelCustomer, error) {
	var customer requests.UpdateCustomerRequest

	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	res, err := s.Repository.Update(&customer)

	return res, err

}
