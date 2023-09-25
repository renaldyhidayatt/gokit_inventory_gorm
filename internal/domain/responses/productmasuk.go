package responses

import (
	"go_kit_inventory/internal/schema"
)

type CreateProductMasukResponse struct {
	ProductMasuk *schema.ProductMasuk `json:"product_masuk,omitempty"`
	Error        error                `json:"error,omitempty"`
	Message      string               `json:"message,omitempty"`
	StatusCode   int                  `json:"status_code,omitempty"`
}

type ResultsProductMasukResponse struct {
	ProductMasuk *[]schema.ProductMasuk `json:"product_masuk,omitempty"`
	Error        error                  `json:"error,omitempty"`
	Message      string                 `json:"message,omitempty"`
	StatusCode   int                    `json:"status_code,omitempty"`
}

type ResultProductMasukResponse struct {
	ProductMasuk *schema.ProductMasuk `json:"product_masuk,omitempty"`
	Error        error                `json:"error,omitempty"`
	Message      string               `json:"message,omitempty"`
	StatusCode   int                  `json:"status_code,omitempty"`
}

type DeleteProductMasukResponse struct {
	ProductMasuk *schema.ProductMasuk `json:"product_masuk,omitempty"`
	Error        error                `json:"error,omitempty"`
	Message      string               `json:"message,omitempty"`
	StatusCode   int                  `json:"status_code,omitempty"`
}

type UpdateProductMasukResponse struct {
	ProductMasuk *schema.ProductMasuk `json:"product_masuk,omitempty"`
	Error        error                `json:"error,omitempty"`
	Message      string               `json:"message,omitempty"`
	StatusCode   int                  `json:"status_code,omitempty"`
}
