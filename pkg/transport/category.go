package transport

import (
	"context"
	"encoding/json"
	"go_kit_inventory/internal/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func DecodeCreateCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeResultsCategoryRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil // No specific decoding needed for this request
}

func DecodeResultCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return domain.ResultCategoryRequest{ID: id}, nil
}

func DecodeDeleteCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return domain.DeleteCategoryRequest{ID: id}, nil
}

func DecodeUpdateCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeCategoryResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}
