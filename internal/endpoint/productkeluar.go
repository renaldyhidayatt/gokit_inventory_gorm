package endpoint

import (
	"context"
	"go_kit_inventory/internal/domain/requests"
	"go_kit_inventory/internal/domain/responses"
	"go_kit_inventory/internal/models"
	"go_kit_inventory/internal/schema"
	"go_kit_inventory/internal/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type ProductKeluarEndpoints struct {
	CreateProductKeluarEndpoint  endpoint.Endpoint
	ResultsProductKeluarEndpoint endpoint.Endpoint
	ResultProductKeluarEndpoint  endpoint.Endpoint
	DeleteProductKeluarEndpoint  endpoint.Endpoint
	UpdateProductKeluarEndpoint  endpoint.Endpoint
}

func NewProductKeluarEndpoints(s service.ProductKeluarService) ProductKeluarEndpoints {
	return ProductKeluarEndpoints{
		CreateProductKeluarEndpoint:  makeCreateProductKeluarEndpoint(s),
		ResultsProductKeluarEndpoint: makeResultsProductKeluarEndpoint(s),
		ResultProductKeluarEndpoint:  makeResultProductKeluarEndpoint(s),
		DeleteProductKeluarEndpoint:  makeDeleteProductKeluarEndpoint(s),
		UpdateProductKeluarEndpoint:  makeUpdateProductKeluarEndpoint(s),
	}
}

func makeCreateProductKeluarEndpoint(s service.ProductKeluarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.CreateProductKeluarRequest)

		if err := req.Validate(); err != nil {
			return responses.CreateProductKeluarResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Create(&req)

		var response responses.CreateProductKeluarResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to create product"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedProduct := convertToProductKeluar(product)
			response.ProductKeluar = convertedProduct
			response.Message = "Product created successfully"
			response.StatusCode = http.StatusCreated
		}

		return response, nil
	}
}

func makeResultsProductKeluarEndpoint(s service.ProductKeluarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.Results()

		var response responses.ResultsProductKeluarResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch products"
			response.StatusCode = http.StatusInternalServerError
		} else {
			var convertedProducts []schema.ProductKeluar
			for _, modelProduct := range *products {
				convertedProduct := convertToProductKeluar(&modelProduct)
				convertedProducts = append(convertedProducts, *convertedProduct)
			}
			response.ProductKeluar = &convertedProducts
			response.Message = "Products fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeResultProductKeluarEndpoint(s service.ProductKeluarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.ResultProductKeluarRequest)

		if err := req.Validate(); err != nil {
			return responses.ResultProductKeluarResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Result(req.ID)

		var response responses.ResultProductKeluarResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch product"
			response.StatusCode = http.StatusNotFound
		} else {
			convertedProduct := convertToProductKeluar(product)
			response.ProductKeluar = convertedProduct
			response.Message = "Product fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeDeleteProductKeluarEndpoint(s service.ProductKeluarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.DeleteProductKeluarRequest)

		if err := req.Validate(); err != nil {
			return responses.DeleteProductKeluarResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Delete(req.ID)

		var response responses.DeleteProductKeluarResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to delete product"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedProduct := convertToProductKeluar(product)
			response.ProductKeluar = convertedProduct
			response.Message = "Product deleted successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeUpdateProductKeluarEndpoint(s service.ProductKeluarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.UpdateProductKeluarRequest)

		if err := req.Validate(); err != nil {
			return responses.UpdateProductKeluarResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Update(&req)

		var response responses.UpdateProductKeluarResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to update product"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedProduct := convertToProductKeluar(product)
			response.ProductKeluar = convertedProduct
			response.Message = "Product updated successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func convertToProductKeluar(modelProductKeluar *models.ModelProductKeluar) *schema.ProductKeluar {
	return &schema.ProductKeluar{
		ID:        modelProductKeluar.ID,
		Qty:       modelProductKeluar.Qty,
		Product:   *convertToProduct(&modelProductKeluar.Product),
		Category:  *convertToCategory(&modelProductKeluar.Category),
		CreatedAt: modelProductKeluar.CreatedAt,
		UpdatedAt: modelProductKeluar.UpdatedAt,
	}
}
