package domain

import (
	"go_kit_inventory/internal/models"

	"github.com/go-playground/validator/v10"
)

type CreateProductMasukRequest struct {
	Name       string `json:"name" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	SupplierID string `json:"supplier_id" validate:"required"`
}

func (c *CreateProductMasukRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type UpdateProductMasukRequest struct {
	ID         string `json:"id" validate:"uuid"`
	Name       string `json:"name" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	SupplierID string `json:"supplier_id" validate:"required"`
}

func (c *UpdateProductMasukRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type CreateProductMasukResponse struct {
	ProductMasuk *models.ModelProductMasuk `json:"product_masuk,omitempty"`
	Error        error                     `json:"error,omitempty"`
}

type ResultsProductMasukResponse struct {
	ProductMasuk []models.ModelProductMasuk `json:"product_masuk,omitempty"`
	Error        error                      `json:"error,omitempty"`
}

type ResultProductMasukResponse struct {
	ProductMasuk *models.ModelProductMasuk `json:"product_masuk,omitempty"`
	Error        error                     `json:"error,omitempty"`
}

type ResultProductMasukRequest struct {
	ID string `json:"id"`
}

type DeleteProductMasukRequest struct {
	ID string `json:"id"`
}

type DeleteProductMasukResponse struct {
	ProductMasuk *models.ModelProductMasuk `json:"product_masuk,omitempty"`
	Error        error                     `json:"error,omitempty"`
}

type UpdateProductMasukResponse struct {
	ProductMasuk *models.ModelProductMasuk `json:"product_masuk,omitempty"`
	Error        error                     `json:"error,omitempty"`
}
