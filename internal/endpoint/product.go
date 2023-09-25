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

type ProductEndpoints struct {
	CreateEndpoint  endpoint.Endpoint
	ResultsEndpoint endpoint.Endpoint
	ResultEndpoint  endpoint.Endpoint
	DeleteEndpoint  endpoint.Endpoint
	UpdateEndpoint  endpoint.Endpoint
}

func NewProductEndpoints(s service.ProductService) ProductEndpoints {
	return ProductEndpoints{
		CreateEndpoint:  makeCreateProductEndpoint(s),
		ResultsEndpoint: makeResultsProductEndpoint(s),
		ResultEndpoint:  makeResultProductEndpoint(s),
		DeleteEndpoint:  makeDeleteProductEndpoint(s),
		UpdateEndpoint:  makeUpdateProductEndpoint(s),
	}
}
func makeCreateProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.CreateProductRequest)

		if err := req.Validate(); err != nil {
			return responses.CreateProductResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Create(&req)

		var response responses.CreateProductResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to create product"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedProduct := convertToProduct(product)
			response.Product = convertedProduct
			response.Message = "Product created successfully"
			response.StatusCode = http.StatusCreated
		}

		return response, nil
	}
}

func makeResultsProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.Results()

		var response responses.ResultsProductResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch products"
			response.StatusCode = http.StatusInternalServerError
		} else {
			var convertedProducts []schema.Product
			for _, modelProduct := range *products {
				convertedProduct := convertToProduct(&modelProduct)
				convertedProducts = append(convertedProducts, *convertedProduct)
			}
			response.Products = &convertedProducts
			response.Message = "Products fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeResultProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.ResultProductRequest)

		if err := req.Validate(); err != nil {
			return responses.ResultProductResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Result(req.ID)

		var response responses.ResultProductResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch product"
			response.StatusCode = http.StatusNotFound
		} else {
			convertedProduct := convertToProduct(product)
			response.Product = convertedProduct
			response.Message = "Product fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeDeleteProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.DeleteProductRequest)

		if err := req.Validate(); err != nil {
			return responses.DeleteProductResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Delete(req.ID)

		var response responses.DeleteProductResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to delete product"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedProduct := convertToProduct(product)
			response.Product = convertedProduct
			response.Message = "Product deleted successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeUpdateProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.UpdateProductRequest)

		if err := req.Validate(); err != nil {
			return responses.UpdateProductResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Update(&req)

		var response responses.UpdateProductResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to update product"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedProduct := convertToProduct(product)
			response.Product = convertedProduct
			response.Message = "Product updated successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func convertToProduct(modelProduct *models.ModelProduct) *schema.Product {
	return &schema.Product{
		ID:        modelProduct.ID,
		Name:      modelProduct.Name,
		Image:     modelProduct.Image,
		Qty:       modelProduct.Qty,
		Category:  schema.Category(modelProduct.Category),
		CreatedAt: modelProduct.CreatedAt,
		UpdatedAt: modelProduct.UpdatedAt,
	}
}
