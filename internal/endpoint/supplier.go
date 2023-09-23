package endpoint

import (
	"context"
	"errors"
	"go_kit_inventory/internal/domain"
	"go_kit_inventory/internal/service"

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
		req, ok := request.(domain.CreateSupplierRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		supplier, err := s.Create(&req)
		if err != nil {
			return nil, err
		}

		return domain.CreateSupplierResponse{Supplier: supplier, Error: nil}, nil
	}
}

func makeResultsSupplierEndpoint(s service.SupplierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		suppliers, err := s.Results()
		if err != nil {
			return nil, err
		}

		return domain.ResultsSupplierResponse{Suppliers: *suppliers, Error: nil}, nil
	}
}

func makeResultSupplierEndpoint(s service.SupplierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.ResultSupplierRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		supplier, err := s.Result(req.ID)
		if err != nil {
			return nil, err
		}

		return domain.ResultSupplierResponse{Supplier: supplier, Error: nil}, nil
	}
}

func makeDeleteSupplierEndpoint(s service.SupplierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.DeleteSupplierRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		supplier, err := s.Delete(req.ID)
		if err != nil {
			return nil, err
		}

		return domain.DeleteSupplierResponse{Supplier: supplier, Error: nil}, nil
	}
}

func makeUpdateSupplierEndpoint(s service.SupplierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.UpdateSupplierRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		supplier, err := s.Update(&req)
		if err != nil {
			return nil, err
		}

		return domain.UpdateSupplierResponse{Supplier: supplier, Error: nil}, nil
	}
}
