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

type SupplierEndpoint struct {
	CreateEndpoint  endpoint.Endpoint
	ResultsEndpoint endpoint.Endpoint
	ResultEndpoint  endpoint.Endpoint
	DeleteEndpoint  endpoint.Endpoint
	UpdateEndpoint  endpoint.Endpoint
}

func NewSupplierEndpoints(s service.SupplierService) SupplierEndpoint {
	return SupplierEndpoint{
		CreateEndpoint:  makeCreateSupplierEndpoint(s),
		ResultsEndpoint: makeResultsSupplierEndpoint(s),
		ResultEndpoint:  makeResultSupplierEndpoint(s),
		DeleteEndpoint:  makeDeleteSupplierEndpoint(s),
		UpdateEndpoint:  makeUpdateSupplierEndpoint(s),
	}
}

func makeCreateSupplierEndpoint(s service.SupplierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.CreateSupplierRequest)

		if err := req.Validate(); err != nil {
			return responses.CreateSupplierResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		supplier, err := s.Create(&req)

		var response responses.CreateSupplierResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to create supplier"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedSupplier := convertToSupplier(supplier)
			response.Supplier = convertedSupplier
			response.Message = "Supplier created successfully"
			response.StatusCode = http.StatusCreated
		}

		return response, nil
	}
}

func makeResultsSupplierEndpoint(s service.SupplierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		supplier, err := s.Results()

		var response responses.ResultsSupplierResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch supplier"
			response.StatusCode = http.StatusInternalServerError
		} else {
			var convertedSuppliers []schema.Supplier
			for _, modelsupplier := range *supplier {
				convertedSupplier := convertToSupplier(&modelsupplier)
				convertedSuppliers = append(convertedSuppliers, *convertedSupplier)
			}
			response.Suppliers = &convertedSuppliers
			response.Message = "Products fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeResultSupplierEndpoint(s service.SupplierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.ResultSupplierRequest)

		if err := req.Validate(); err != nil {
			return responses.ResultSupplierResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		supplier, err := s.Result(req.ID)

		var response responses.ResultSupplierResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch supplier"
			response.StatusCode = http.StatusNotFound
		} else {
			convertedSupplier := convertToSupplier(supplier)
			response.Supplier = convertedSupplier
			response.Message = "Supplier fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeDeleteSupplierEndpoint(s service.SupplierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.DeleteSupplierRequest)
		supplier, err := s.Delete(req.ID)

		if err := req.Validate(); err != nil {
			return responses.DeleteSupplierResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		var response responses.DeleteSupplierResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to delete supplier"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedProduct := convertToSupplier(supplier)
			response.Supplier = convertedProduct
			response.Message = "Supplier deleted successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeUpdateSupplierEndpoint(s service.SupplierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.UpdateSupplierRequest)

		if err := req.Validate(); err != nil {
			return responses.UpdateSupplierResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		supplier, err := s.Update(&req)

		var response responses.UpdateSupplierResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to update supplier"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedSupplier := convertToSupplier(supplier)
			response.Supplier = convertedSupplier
			response.Message = "Product updated successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func convertToSupplier(modelSupplier *models.ModelSupplier) *schema.Supplier {
	return &schema.Supplier{
		ID:        modelSupplier.ID,
		Name:      modelSupplier.Name,
		Alamat:    modelSupplier.Alamat,
		Email:     modelSupplier.Email,
		Telepon:   modelSupplier.Telepon,
		CreatedAt: modelSupplier.CreatedAt,
		UpdatedAt: modelSupplier.UpdatedAt,
	}
}
