package endpoint

import (
	"context"
	"go_kit_inventory/internal/domain"
	"go_kit_inventory/internal/service"

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
		req := request.(domain.CreateProductMasukRequest)
		productMasuk, err := s.Create(&req)
		if err != nil {
			return nil, err
		}
		return domain.CreateProductMasukResponse{ProductMasuk: productMasuk, Error: nil}, nil
	}
}

func makeResultsProductMasukEndpoint(s service.ProductMasukService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		productMasuk, err := s.Results()
		if err != nil {
			return nil, err
		}
		return domain.ResultsProductMasukResponse{ProductMasuk: *productMasuk, Error: nil}, nil
	}
}

func makeResultProductMasukEndpoint(s service.ProductMasukService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.ResultProductMasukRequest)
		productMasuk, err := s.Result(req.ID)
		if err != nil {
			return nil, err
		}
		return domain.ResultProductMasukResponse{ProductMasuk: productMasuk, Error: nil}, nil
	}
}

func makeDeleteProductMasukEndpoint(s service.ProductMasukService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.DeleteProductMasukRequest)
		productMasuk, err := s.Delete(req.ID)
		if err != nil {
			return nil, err
		}
		return domain.DeleteProductMasukResponse{ProductMasuk: productMasuk, Error: nil}, nil
	}
}

func makeUpdateProductMasukEndpoint(s service.ProductMasukService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.UpdateProductMasukRequest)
		productMasuk, err := s.Update(&req)
		if err != nil {
			return nil, err
		}
		return domain.UpdateProductMasukResponse{ProductMasuk: productMasuk, Error: nil}, nil
	}
}
