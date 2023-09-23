package endpoint

import (
	"context"
	"go_kit_inventory/internal/domain"
	"go_kit_inventory/internal/service"

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
		req := request.(domain.CreateCustomerRequest)
		customer, err := s.Create(&req)
		return domain.CreateCustomerResponse{Customer: customer, Error: err}, nil
	}
}

func makeResultsCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		customers, err := s.Results()
		return domain.ResultsCustomerResponse{Customers: customers, Error: err}, nil
	}
}

func makeResultCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.ResultCustomerRequest)
		customer, err := s.Result(req.ID)
		return domain.ResultCustomerResponse{Customer: customer, Error: err}, nil
	}
}

func makeDeleteCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.DeleteCustomerRequest)
		customer, err := s.Delete(req.ID)
		return domain.DeleteCustomerResponse{Customer: customer, Error: err}, nil
	}
}

func makeUpdateCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.UpdateCustomerRequest)
		customer, err := s.Update(&req)
		return domain.UpdateCustomerResponse{Customer: customer, Error: err}, nil
	}
}
