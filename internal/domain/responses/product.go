package responses

import (
	"go_kit_inventory/internal/schema"
)

type CreateProductResponse struct {
	Product    *schema.Product `json:"product,omitempty"`
	Message    string          `json:"message,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
	Error      error           `json:"error,omitempty"`
}

type ResultsProductResponse struct {
	Products   *[]schema.Product `json:"products,omitempty"`
	Message    string            `json:"message,omitempty"`
	StatusCode int               `json:"status_code,omitempty"`
	Error      error             `json:"error,omitempty"`
}

type ResultProductResponse struct {
	Product    *schema.Product `json:"product,omitempty"`
	Message    string          `json:"message,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
	Error      error           `json:"error,omitempty"`
}

type DeleteProductResponse struct {
	Product    *schema.Product `json:"product,omitempty"`
	Message    string          `json:"message,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
	Error      error           `json:"error,omitempty"`
}

type UpdateProductResponse struct {
	Product    *schema.Product `json:"product,omitempty"`
	Message    string          `json:"message,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
	Error      error           `json:"error,omitempty"`
}
