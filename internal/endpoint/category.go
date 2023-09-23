package endpoint

import (
	"context"
	"go_kit_inventory/internal/domain"
	"go_kit_inventory/internal/service"

	"github.com/go-kit/kit/endpoint"
)

type CategoryEndpoint struct {
	CreateEndpoint  endpoint.Endpoint
	ResultsEndpoint endpoint.Endpoint
	ResultEndpoint  endpoint.Endpoint
	DeleteEndpoint  endpoint.Endpoint
	UpdateEndpoint  endpoint.Endpoint
}

func NewCategoryEndpoints(s service.CategoryService) CategoryEndpoint {
	return CategoryEndpoint{
		CreateEndpoint:  MakeCreateCategoryEndpoint(s),
		ResultsEndpoint: MakeResultsCategoryEndpoint(s),
		ResultEndpoint:  MakeResultCategoryEndpoint(s),
		DeleteEndpoint:  MakeDeleteCategoryEndpoint(s),
		UpdateEndpoint:  MakeUpdateCategoryEndpoint(s),
	}
}

func MakeCreateCategoryEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.CreateCategoryRequest)
		category, err := s.Create(&req)
		return domain.CreateCategoryResponse{Category: category, Error: err}, nil
	}
}

func MakeResultsCategoryEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		categories, err := s.Results()
		return domain.ResultsCategoryResponse{Categories: *categories, Error: err}, nil
	}
}

func MakeResultCategoryEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.ResultCategoryRequest)
		category, err := s.Result(req.ID)
		return domain.ResultCategoryResponse{Category: category, Error: err}, nil
	}
}

func MakeDeleteCategoryEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.DeleteCategoryRequest)
		category, err := s.Delete(req.ID)
		return domain.DeleteCategoryResponse{Category: category, Error: err}, nil
	}
}

func MakeUpdateCategoryEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.UpdateCategoryRequest)
		category, err := s.Update(&req)
		return domain.UpdateCategoryResponse{Category: category, Error: err}, nil
	}
}
