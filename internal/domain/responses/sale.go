package responses

import (
	"go_kit_inventory/internal/schema"
)

type CreateSaleResponse struct {
	Sale       *schema.Sale `json:"sale,omitempty"`
	Error      error        `json:"error,omitempty"`
	Message    string       `json:"message,omitempty"`
	StatusCode int          `json:"status_code,omitempty"`
}

type ResultsSaleResponse struct {
	Sales      *[]schema.Sale `json:"sales,omitempty"`
	Error      error          `json:"error,omitempty"`
	Message    string         `json:"message,omitempty"`
	StatusCode int            `json:"status_code,omitempty"`
}

type ResultSaleResponse struct {
	Sale       *schema.Sale `json:"sale,omitempty"`
	Error      error        `json:"error,omitempty"`
	Message    string       `json:"message,omitempty"`
	StatusCode int          `json:"status_code,omitempty"`
}

type DeleteSaleResponse struct {
	Sale       *schema.Sale `json:"sale,omitempty"`
	Error      error        `json:"error,omitempty"`
	Message    string       `json:"message,omitempty"`
	StatusCode int          `json:"status_code,omitempty"`
}

type UpdateSaleResponse struct {
	Sale       *schema.Sale `json:"sale,omitempty"`
	Error      error        `json:"error,omitempty"`
	Message    string       `json:"message,omitempty"`
	StatusCode int          `json:"status_code,omitempty"`
}
