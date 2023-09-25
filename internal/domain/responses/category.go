package responses

import (
	"go_kit_inventory/internal/schema"
)

type ResultCategoryResponse struct {
	Category   *schema.Category `json:"category,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}

type DeleteCategoryResponse struct {
	Category   *schema.Category `json:"category,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}

type UpdateCategoryResponse struct {
	Category   *schema.Category `json:"category,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}

type CreateCategoryResponse struct {
	Category   *schema.Category `json:"category,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}

type ResultsCategoryResponse struct {
	Categories *[]schema.Category `json:"categories,omitempty"`
	Error      error              `json:"error,omitempty"`
	Message    string             `json:"message,omitempty"`
	StatusCode int                `json:"status_code,omitempty"`
}
