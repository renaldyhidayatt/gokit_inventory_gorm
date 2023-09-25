package responses

import (
	"go_kit_inventory/internal/schema"
)

type CreateProductKeluarResponse struct {
	ProductKeluar *schema.ProductKeluar `json:"product_keluar,omitempty"`
	Message       string                `json:"message,omitempty"`
	StatusCode    int                   `json:"status_code,omitempty"`
	Error         error                 `json:"error,omitempty"`
}

type ResultsProductKeluarResponse struct {
	ProductKeluar *[]schema.ProductKeluar `json:"product_keluar,omitempty"`
	Message       string                  `json:"message,omitempty"`
	StatusCode    int                     `json:"status_code,omitempty"`
	Error         error                   `json:"error,omitempty"`
}

type ResultProductKeluarResponse struct {
	ProductKeluar *schema.ProductKeluar `json:"product_keluar,omitempty"`
	Message       string                `json:"message,omitempty"`
	StatusCode    int                   `json:"status_code,omitempty"`
	Error         error                 `json:"error,omitempty"`
}

type DeleteProductKeluarResponse struct {
	ProductKeluar *schema.ProductKeluar `json:"product_keluar,omitempty"`
	Message       string                `json:"message,omitempty"`
	StatusCode    int                   `json:"status_code,omitempty"`
	Error         error                 `json:"error,omitempty"`
}

type UpdateProductKeluarResponse struct {
	ProductKeluar *schema.ProductKeluar `json:"product_keluar,omitempty"`
	Message       string                `json:"message,omitempty"`
	StatusCode    int                   `json:"status_code,omitempty"`
	Error         error                 `json:"error,omitempty"`
}
