package endpoint

import (
	"context"
	"go_kit_inventory/internal/domain"
	"go_kit_inventory/internal/service"

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
		CreateEndpoint:  makeCreateEndpoint(s),
		ResultsEndpoint: makeResultsEndpoint(s),
		ResultEndpoint:  makeResultEndpoint(s),
		DeleteEndpoint:  makeDeleteEndpoint(s),
		UpdateEndpoint:  makeUpdateEndpoint(s),
	}
}

func makeCreateEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.CreateProductRequest)
		product, err := s.Create(&req)
		return domain.CreateProductResponse{Product: product, Error: err}, nil
	}
}

func makeResultsEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.Results()
		return domain.ResultsProductResponse{Products: products, Error: err}, nil
	}
}

func makeResultEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.ResultProductRequest)
		product, err := s.Result(req.ID)
		return domain.ResultProductResponse{Product: product, Error: err}, nil
	}
}

func makeDeleteEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.DeleteProductRequest)
		product, err := s.Delete(req.ID)
		return domain.DeleteProductResponse{Product: product, Error: err}, nil
	}
}

func makeUpdateEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.UpdateProductRequest)
		product, err := s.Update(&req)
		return domain.UpdateProductResponse{Product: product, Error: err}, nil
	}
}
