package domain

import (
	"go_kit_inventory/internal/models"

	"github.com/go-playground/validator/v10"
)

type CreateProductKeluarRequest struct {
	ID         string `json:"id" validate:"uuid"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
}

func (c *CreateProductKeluarRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type UpdateProductKeluarRequest struct {
	ID         string `json:"id" validate:"uuid"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
}

func (c *UpdateProductKeluarRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type CreateProductKeluarResponse struct {
	ProductKeluar *models.ModelProductKeluar `json:"product_keluar,omitempty"`
	Error         error                      `json:"error,omitempty"`
}

type ResultsProductKeluarResponse struct {
	ProductKeluar []models.ModelProductKeluar `json:"product_keluar,omitempty"`
	Error         error                       `json:"error,omitempty"`
}

type ResultProductKeluarResponse struct {
	ProductKeluar *models.ModelProductKeluar `json:"product_keluar,omitempty"`
	Error         error                      `json:"error,omitempty"`
}

type ResultProductKeluarRequest struct {
	ID string `json:"id"`
}

type DeleteProductKeluarRequest struct {
	ID string `json:"id"`
}

type DeleteProductKeluarResponse struct {
	ProductKeluar *models.ModelProductKeluar `json:"product_keluar,omitempty"`
	Error         error                      `json:"error,omitempty"`
}

// UpdateProductKeluarResponse represents the response for updating a product keluar.
type UpdateProductKeluarResponse struct {
	ProductKeluar *models.ModelProductKeluar `json:"product_keluar,omitempty"`
	Error         error                      `json:"error,omitempty"`
}
