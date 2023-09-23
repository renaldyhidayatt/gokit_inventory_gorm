package domain

import (
	"go_kit_inventory/internal/models"

	"github.com/go-playground/validator/v10"
)

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,lowercase" schema:"name"`
}

func (c *CreateCategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type UpdateCategoryRequest struct {
	ID   string `json:"id" validate:"required,uuid"`
	Name string `json:"name" validate:"required,lowercase" schema:"name"`
}

func (c *UpdateCategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type CreateCategoryResponse struct {
	Category *models.ModelCategory `json:"category,omitempty"`
	Error    error                 `json:"error,omitempty"`
}

type ResultsCategoryResponse struct {
	Categories []models.ModelCategory `json:"categories,omitempty"`
	Error      error                  `json:"error,omitempty"`
}

type ResultCategoryRequest struct {
	ID string `json:"id"`
}

type ResultCategoryResponse struct {
	Category *models.ModelCategory `json:"category,omitempty"`
	Error    error                 `json:"error,omitempty"`
}

type DeleteCategoryRequest struct {
	ID string `json:"id"`
}

type DeleteCategoryResponse struct {
	Category *models.ModelCategory `json:"category,omitempty"`
	Error    error                 `json:"error,omitempty"`
}

type UpdateCategoryResponse struct {
	Category *models.ModelCategory `json:"category,omitempty"`
	Error    error                 `json:"error,omitempty"`
}
