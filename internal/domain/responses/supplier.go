package responses

import (
	"go_kit_inventory/internal/schema"
)

type CreateSupplierResponse struct {
	Supplier   *schema.Supplier `json:"supplier,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}

type ResultsSupplierResponse struct {
	Suppliers  *[]schema.Supplier `json:"suppliers,omitempty"`
	Error      error              `json:"error,omitempty"`
	Message    string             `json:"message,omitempty"`
	StatusCode int                `json:"status_code,omitempty"`
}

type ResultSupplierResponse struct {
	Supplier   *schema.Supplier `json:"supplier,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}

type DeleteSupplierResponse struct {
	Supplier   *schema.Supplier `json:"supplier,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}

type UpdateSupplierResponse struct {
	Supplier   *schema.Supplier `json:"supplier,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}
