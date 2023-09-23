package endpoint

import (
	"context"
	"errors"
	"go_kit_inventory/internal/domain"
	"go_kit_inventory/internal/service"

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
		req, ok := request.(domain.CreateSaleRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		sale, err := s.Create(&req)
		if err != nil {
			return nil, err
		}

		return domain.CreateSaleResponse{Sale: sale, Error: nil}, nil
	}
}

func makeResultsSaleEndpoint(s service.SaleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		sales, err := s.Results()
		if err != nil {
			return nil, err
		}

		return domain.ResultsSaleResponse{Sales: *sales, Error: nil}, nil
	}
}

func makeResultSaleEndpoint(s service.SaleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.ResultSaleRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		sale, err := s.Result(req.ID)
		if err != nil {
			return nil, err
		}

		return domain.ResultSaleResponse{Sale: sale, Error: nil}, nil
	}
}

func makeDeleteSaleEndpoint(s service.SaleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.DeleteSaleRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		sale, err := s.Delete(req.ID)
		if err != nil {
			return nil, err
		}

		return domain.DeleteSaleResponse{Sale: sale, Error: nil}, nil
	}
}

func makeUpdateSaleEndpoint(s service.SaleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.UpdateSaleRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		sale, err := s.Update(&req)
		if err != nil {
			return nil, err
		}

		return domain.UpdateSaleResponse{Sale: sale, Error: nil}, nil
	}
}
