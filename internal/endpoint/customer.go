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

type CustomerEndpoint struct {
	CreateEndpoint  endpoint.Endpoint
	ResultsEndpoint endpoint.Endpoint
	ResultEndpoint  endpoint.Endpoint
	DeleteEndpoint  endpoint.Endpoint
	UpdateEndpoint  endpoint.Endpoint
}

func NewCustomerEndpoints(s service.CustomerService) CategoryEndpoint {
	return CategoryEndpoint{
		CreateEndpoint:  makeCreateCustomerEndpoint(s),
		ResultsEndpoint: makeResultsCustomerEndpoint(s),
		ResultEndpoint:  makeResultCustomerEndpoint(s),
		DeleteEndpoint:  makeDeleteCustomerEndpoint(s),
		UpdateEndpoint:  makeUpdateCustomerEndpoint(s),
	}
}

func makeCreateCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.CreateCustomerRequest)

		if err := req.Validate(); err != nil {
			return responses.CreateCustomerResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		customer, err := s.Create(&req)

		var response responses.CreateCustomerResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to create customer"
			response.StatusCode = http.StatusInternalServerError
		} else {

			convertedCustomer := convertToCustomer(customer)
			response.Customer = convertedCustomer
			response.Message = "Customer created successfully"
			response.StatusCode = http.StatusCreated
		}

		return response, nil
	}
}

func makeResultsCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		customers, err := s.Results()

		var response responses.ResultsCustomerResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch customers"
			response.StatusCode = http.StatusInternalServerError
		} else {

			var convertedCustomers []schema.Customer
			for _, modelCustomer := range *customers {
				convertedCustomer := convertToCustomer(&modelCustomer)
				convertedCustomers = append(convertedCustomers, *convertedCustomer)
			}
			response.Customers = &convertedCustomers
			response.Message = "Customers fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeResultCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.ResultCustomerRequest)

		if err := req.Validate(); err != nil {
			return responses.ResultCustomerResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		customer, err := s.Result(req.ID)

		var response responses.ResultCustomerResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch customer"
			response.StatusCode = http.StatusNotFound
		} else {

			convertedCustomer := convertToCustomer(customer)
			response.Customer = convertedCustomer
			response.Message = "Customer fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeDeleteCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.DeleteCustomerRequest)

		if err := req.Validate(); err != nil {
			return responses.DeleteCustomerResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		customer, err := s.Delete(req.ID)

		var response responses.DeleteCustomerResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to delete customer"
			response.StatusCode = http.StatusInternalServerError
		} else {

			convertedCustomer := convertToCustomer(customer)
			response.Customer = convertedCustomer
			response.Message = "Customer deleted successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeUpdateCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.UpdateCustomerRequest)

		if err := req.Validate(); err != nil {
			return responses.UpdateCustomerResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		customer, err := s.Update(&req)

		var response responses.UpdateCustomerResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to update customer"
			response.StatusCode = http.StatusInternalServerError
		} else {

			convertedCustomer := convertToCustomer(customer)
			response.Customer = convertedCustomer
			response.Message = "Customer updated successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}
func convertToCustomer(modelCustomer *models.ModelCustomer) *schema.Customer {
	return &schema.Customer{
		ID:        modelCustomer.ID,
		Name:      modelCustomer.Name,
		Alamat:    modelCustomer.Alamat,
		Email:     modelCustomer.Email,
		CreatedAt: modelCustomer.CreatedAt,
		UpdatedAt: modelCustomer.UpdatedAt,
	}
}
