package domain

import (
	"go_kit_inventory/internal/models"

	"github.com/go-playground/validator/v10"
)

type CreateProductRequest struct {
	Name       string `json:"name" validate:"required,lowercase"`
	Image      string `json:"image" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
}

func (c *CreateProductRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type UpdateProductRequest struct {
	ID         string `json:"id" validate:"uuid"`
	Name       string `json:"name" validate:"required,lowercase"`
	Image      string `json:"image" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
}

func (c *UpdateProductRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type CreateProductResponse struct {
	Product *models.ModelProduct `json:"product,omitempty"`
	Error   error                `json:"error,omitempty"`
}

type ResultsProductResponse struct {
	Products *[]models.ModelProduct `json:"products,omitempty"`
	Error    error                  `json:"error,omitempty"`
}

type ResultProductRequest struct {
	ID string `json:"id"`
}

type ResultProductResponse struct {
	Product *models.ModelProduct `json:"product,omitempty"`
	Error   error                `json:"error,omitempty"`
}

type DeleteProductRequest struct {
	ID string `json:"id"`
}

type DeleteProductResponse struct {
	Product *models.ModelProduct `json:"product,omitempty"`
	Error   error                `json:"error,omitempty"`
}

type UpdateProductResponse struct {
	Product *models.ModelProduct `json:"product,omitempty"`
	Error   error                `json:"error,omitempty"`
}
