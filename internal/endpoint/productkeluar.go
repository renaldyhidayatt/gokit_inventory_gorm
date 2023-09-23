package endpoint

import (
	"context"
	"go_kit_inventory/internal/domain"
	"go_kit_inventory/internal/service"

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
		req := request.(domain.CreateProductKeluarRequest)
		productkeluar, err := s.Create(&req)
		return domain.CreateProductKeluarResponse{ProductKeluar: productkeluar, Error: err}, nil
	}
}

func makeResultsProductKeluarEndpoint(s service.ProductKeluarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		productskeluar, err := s.Results()
		return domain.ResultsProductKeluarResponse{ProductKeluar: *productskeluar, Error: err}, nil
	}
}

func makeResultProductKeluarEndpoint(s service.ProductKeluarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.ResultProductKeluarRequest)
		productkeluar, err := s.Result(req.ID)
		return domain.ResultProductKeluarResponse{ProductKeluar: productkeluar, Error: err}, nil
	}
}

func makeDeleteProductKeluarEndpoint(s service.ProductKeluarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.DeleteProductKeluarRequest)
		productkeluar, err := s.Delete(req.ID)
		return domain.DeleteProductKeluarResponse{ProductKeluar: productkeluar, Error: err}, nil
	}
}

func makeUpdateProductKeluarEndpoint(s service.ProductKeluarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.UpdateProductKeluarRequest)
		productkeluar, err := s.Update(&req)
		return domain.UpdateProductKeluarResponse{ProductKeluar: productkeluar, Error: err}, nil
	}
}
