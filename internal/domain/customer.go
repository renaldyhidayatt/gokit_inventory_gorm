package domain

import (
	"go_kit_inventory/internal/models"

	"github.com/go-playground/validator/v10"
)

type CreateCustomerRequest struct {
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Email   string `json:"email" validate:"required,email"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

type CreateCustomerResponse struct {
	Customer *models.ModelCustomer `json:"customer,omitempty"`
	Error    error                 `json:"error,omitempty"`
}

func (c *CreateCustomerRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type UpdateCustomerRequest struct {
	ID      string `json:"id" validate:"required,uuid"`
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Email   string `json:"email" validate:"required,email"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

type UpdateCustomerResponse struct {
	Customer *models.ModelCustomer `json:"customer,omitempty"`
	Error    error                 `json:"error,omitempty"`
}

func (c *UpdateCustomerRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type ResultsCustomerResponse struct {
	Customers *[]models.ModelCustomer `json:"customers,omitempty"`
	Error     error                   `json:"error,omitempty"`
}

type ResultCustomerRequest struct {
	ID string `json:"id"`
}

type ResultCustomerResponse struct {
	Customer *models.ModelCustomer `json:"customer,omitempty"`
	Error    error                 `json:"error,omitempty"`
}

type DeleteCustomerRequest struct {
	ID string `json:"id"`
}

type DeleteCustomerResponse struct {
	Customer *models.ModelCustomer `json:"customer,omitempty"`
	Error    error                 `json:"error,omitempty"`
}
