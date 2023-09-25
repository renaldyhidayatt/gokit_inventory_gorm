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

type SaleEndpoint struct {
	CreateEndpoint  endpoint.Endpoint
	ResultsEndpoint endpoint.Endpoint
	ResultEndpoint  endpoint.Endpoint
	DeleteEndpoint  endpoint.Endpoint
	UpdateEndpoint  endpoint.Endpoint
}

func NewSaleEndpoints(s service.SaleService) SaleEndpoint {
	return SaleEndpoint{
		CreateEndpoint:  makeCreateSaleEndpoint(s),
		ResultsEndpoint: makeResultsSaleEndpoint(s),
		ResultEndpoint:  makeResultSaleEndpoint(s),
		DeleteEndpoint:  makeDeleteSaleEndpoint(s),
		UpdateEndpoint:  makeUpdateSaleEndpoint(s),
	}
}

func makeCreateSaleEndpoint(s service.SaleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.CreateSaleRequest)

		if err := req.Validate(); err != nil {
			return responses.CreateSaleResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		sale, err := s.Create(&req)

		var response responses.CreateSaleResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to create sale"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedSale := convertToSale(sale)
			response.Sale = convertedSale
			response.Message = "Sale created successfully"
			response.StatusCode = http.StatusCreated
		}

		return response, nil
	}
}

func makeResultsSaleEndpoint(s service.SaleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		sale, err := s.Results()

		var response responses.ResultsSaleResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch sale"
			response.StatusCode = http.StatusInternalServerError
		} else {
			var convertedSales []schema.Sale
			for _, modelsupplier := range *sale {
				convertedSale := convertToSale(&modelsupplier)
				convertedSales = append(convertedSales, *convertedSale)
			}
			response.Sales = &convertedSales
			response.Message = "Sales fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeResultSaleEndpoint(s service.SaleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.ResultSaleRequest)

		if err := req.Validate(); err != nil {
			return responses.ResultSaleResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		sale, err := s.Result(req.ID)

		var response responses.ResultSaleResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch sale"
			response.StatusCode = http.StatusNotFound
		} else {
			convertedSale := convertToSale(sale)
			response.Sale = convertedSale
			response.Message = "Sale fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeDeleteSaleEndpoint(s service.SaleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.DeleteSaleRequest)

		if err := req.Validate(); err != nil {
			return responses.DeleteSaleResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		sale, err := s.Delete(req.ID)

		var response responses.DeleteSaleResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to delete sale"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedSale := convertToSale(sale)
			response.Sale = convertedSale
			response.Message = "Sale deleted successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeUpdateSaleEndpoint(s service.SaleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.UpdateSaleRequest)

		if err := req.Validate(); err != nil {
			return responses.UpdateSaleResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		sale, err := s.Update(&req)

		var response responses.UpdateSaleResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to update sale"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedSale := convertToSale(sale)
			response.Sale = convertedSale
			response.Message = "Sale updated successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func convertToSale(modelSale *models.ModelSale) *schema.Sale {
	return &schema.Sale{
		ID:        modelSale.ID,
		Name:      modelSale.Name,
		Alamat:    modelSale.Alamat,
		Email:     modelSale.Email,
		Telepon:   modelSale.Telepon,
		CreatedAt: modelSale.CreatedAt,
		UpdatedAt: modelSale.UpdatedAt,
	}
}
