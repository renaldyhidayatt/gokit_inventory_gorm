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

type ProductMasukEndpoints struct {
	CreateProductMasukEndpoint  endpoint.Endpoint
	ResultsProductMasukEndpoint endpoint.Endpoint
	ResultProductMasukEndpoint  endpoint.Endpoint
	DeleteProductMasukEndpoint  endpoint.Endpoint
	UpdateProductMasukEndpoint  endpoint.Endpoint
}

func NewProductMasukEndpoints(s service.ProductMasukService) ProductMasukEndpoints {
	return ProductMasukEndpoints{
		CreateProductMasukEndpoint:  makeCreateProductMasukEndpoint(s),
		ResultsProductMasukEndpoint: makeResultsProductMasukEndpoint(s),
		ResultProductMasukEndpoint:  makeResultProductMasukEndpoint(s),
		DeleteProductMasukEndpoint:  makeDeleteProductMasukEndpoint(s),
		UpdateProductMasukEndpoint:  makeUpdateProductMasukEndpoint(s),
	}
}

func makeCreateProductMasukEndpoint(s service.ProductMasukService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.CreateProductMasukRequest)

		if err := req.Validate(); err != nil {
			return responses.CreateProductMasukResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Create(&req)

		var response responses.CreateProductMasukResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to create product"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedProduct := convertToProductMasuk(product)
			response.ProductMasuk = convertedProduct
			response.Message = "Product created successfully"
			response.StatusCode = http.StatusCreated
		}

		return response, nil
	}
}

func makeResultsProductMasukEndpoint(s service.ProductMasukService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.Results()

		var response responses.ResultsProductMasukResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch products"
			response.StatusCode = http.StatusInternalServerError
		} else {
			var convertedProducts []schema.ProductMasuk
			for _, modelProduct := range *products {
				convertedProduct := convertToProductMasuk(&modelProduct)
				convertedProducts = append(convertedProducts, *convertedProduct)
			}
			response.ProductMasuk = &convertedProducts
			response.Message = "Products fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeResultProductMasukEndpoint(s service.ProductMasukService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.ResultProductMasukRequest)

		if err := req.Validate(); err != nil {
			return responses.ResultProductMasukResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Result(req.ID)

		var response responses.ResultProductMasukResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch product"
			response.StatusCode = http.StatusNotFound
		} else {
			convertedProduct := convertToProductMasuk(product)
			response.ProductMasuk = convertedProduct
			response.Message = "Product fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeDeleteProductMasukEndpoint(s service.ProductMasukService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.DeleteProductMasukRequest)

		if err := req.Validate(); err != nil {
			return responses.DeleteProductMasukResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Delete(req.ID)

		var response responses.DeleteProductMasukResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to delete product"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedProduct := convertToProductMasuk(product)
			response.ProductMasuk = convertedProduct
			response.Message = "Product deleted successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeUpdateProductMasukEndpoint(s service.ProductMasukService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.UpdateProductMasukRequest)

		if err := req.Validate(); err != nil {
			return responses.UpdateProductMasukResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		product, err := s.Update(&req)

		var response responses.UpdateProductMasukResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to update product"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedProduct := convertToProductMasuk(product)
			response.ProductMasuk = convertedProduct
			response.Message = "Product updated successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func convertToProductMasuk(modelProductMasuk *models.ModelProductMasuk) *schema.ProductMasuk {
	return &schema.ProductMasuk{
		ID:        modelProductMasuk.ID,
		Qty:       modelProductMasuk.Qty,
		Product:   *convertToProduct(&modelProductMasuk.Product),
		Supplier:  *convertToSupplier(&modelProductMasuk.Supplier),
		CreatedAt: modelProductMasuk.CreatedAt,
		UpdatedAt: modelProductMasuk.UpdatedAt,
	}
}
