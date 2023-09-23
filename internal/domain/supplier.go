package domain

import (
	"go_kit_inventory/internal/models"

	"github.com/go-playground/validator/v10"
)

type CreateSupplierRequest struct {
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Email   string `json:"email" validate:"required,email"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

func (c *CreateSupplierRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type UpdateSupplierRequest struct {
	ID      string `json:"id" validate:"required,uuid"`
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Email   string `json:"email" validate:"required,email"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

func (c *UpdateSupplierRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type CreateSupplierResponse struct {
	Supplier *models.ModelSupplier `json:"supplier,omitempty"`
	Error    error                 `json:"error,omitempty"`
}

type ResultsSupplierResponse struct {
	Suppliers []models.ModelSupplier `json:"suppliers,omitempty"`
	Error     error                  `json:"error,omitempty"`
}

type ResultSupplierResponse struct {
	Supplier *models.ModelSupplier `json:"supplier,omitempty"`
	Error    error                 `json:"error,omitempty"`
}

type ResultSupplierRequest struct {
	ID string `json:"id"`
}

type DeleteSupplierRequest struct {
	ID string `json:"id"`
}

type DeleteSupplierResponse struct {
	Supplier *models.ModelSupplier `json:"supplier,omitempty"`
	Error    error                 `json:"error,omitempty"`
}

type UpdateSupplierResponse struct {
	Supplier *models.ModelSupplier `json:"supplier,omitempty"`
	Error    error                 `json:"error,omitempty"`
}
