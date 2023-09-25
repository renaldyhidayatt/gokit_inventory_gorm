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

type CategoryEndpoint struct {
	CreateEndpoint  endpoint.Endpoint
	ResultsEndpoint endpoint.Endpoint
	ResultEndpoint  endpoint.Endpoint
	DeleteEndpoint  endpoint.Endpoint
	UpdateEndpoint  endpoint.Endpoint
}

func NewCategoryEndpoints(s service.CategoryService) CategoryEndpoint {
	return CategoryEndpoint{
		CreateEndpoint:  makeCreateCategoryEndpoint(s),
		ResultsEndpoint: makeResultsCategoryEndpoint(s),
		ResultEndpoint:  makeResultCategoryEndpoint(s),
		DeleteEndpoint:  makeDeleteCategoryEndpoint(s),
		UpdateEndpoint:  makeUpdateCategoryEndpoint(s),
	}
}

func makeCreateCategoryEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.CreateCategoryRequest)

		if err := req.Validate(); err != nil {
			return responses.CreateCategoryResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		category, err := s.Create(&req)

		var response responses.CreateCategoryResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to create category"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convert := convertToCategory(category)
			response.Category = convert
			response.Message = "Category created successfully"
			response.StatusCode = http.StatusCreated
		}

		return response, nil
	}
}

func makeResultsCategoryEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		categories, err := s.Results()

		var response responses.ResultsCategoryResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch categories"
			response.StatusCode = http.StatusInternalServerError
		} else {
			var convertedCategories []schema.Category
			for _, modelCategory := range *categories {
				convertedCategory := convertToCategory(&modelCategory)
				convertedCategories = append(convertedCategories, *convertedCategory)
			}
			response.Categories = &convertedCategories
			response.Message = "Categories fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeResultCategoryEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.ResultCategoryRequest)

		if err := req.Validate(); err != nil {
			return responses.ResultCategoryResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		category, err := s.Result(req.ID)

		var response responses.ResultCategoryResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch category"
			response.StatusCode = http.StatusNotFound
		} else {
			convert := convertToCategory(category)
			response.Category = convert
			response.Message = "Category fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeDeleteCategoryEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.DeleteCategoryRequest)

		if err := req.Validate(); err != nil {
			return responses.DeleteCategoryResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		category, err := s.Delete(req.ID)

		var response responses.DeleteCategoryResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to delete category"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convert := convertToCategory(category)
			response.Category = convert
			response.Message = "Category deleted successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeUpdateCategoryEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.UpdateCategoryRequest)

		if err := req.Validate(); err != nil {
			return responses.UpdateCategoryResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		category, err := s.Update(&req)

		var response responses.UpdateCategoryResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to update category"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convert := convertToCategory(category)
			response.Category = convert
			response.Message = "Category updated successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func convertToCategory(modelCategory *models.ModelCategory) *schema.Category {
	return &schema.Category{
		ID:        modelCategory.ID,
		Name:      modelCategory.Name,
		CreatedAt: modelCategory.CreatedAt,
		UpdatedAt: modelCategory.UpdatedAt,
	}
}
