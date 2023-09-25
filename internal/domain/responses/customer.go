package responses

import (
	"go_kit_inventory/internal/schema"
)

type CreateCustomerResponse struct {
	Customer   *schema.Customer `json:"customer,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}

type UpdateCustomerResponse struct {
	Customer   *schema.Customer `json:"customer,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}

type ResultsCustomerResponse struct {
	Customers  *[]schema.Customer `json:"customers,omitempty"`
	Error      error              `json:"error,omitempty"`
	Message    string             `json:"message,omitempty"`
	StatusCode int                `json:"status_code,omitempty"`
}

type ResultCustomerResponse struct {
	Customer   *schema.Customer `json:"customer,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}

type DeleteCustomerResponse struct {
	Customer   *schema.Customer `json:"customer,omitempty"`
	Error      error            `json:"error,omitempty"`
	Message    string           `json:"message,omitempty"`
	StatusCode int              `json:"status_code,omitempty"`
}
