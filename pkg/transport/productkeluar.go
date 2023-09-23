package transport

import (
	"context"
	"encoding/json"
	"go_kit_inventory/internal/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func DecodeCreateProductKeluarRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req domain.CreateProductKeluarRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeResultsProductKeluarRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodeResultsProductkeluarRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodeResultProductKeluarRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")

	return domain.ResultProductRequest{ID: id}, nil
}

func DecodeUpdateProductKeluarRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req domain.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeProductKeluarResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func DecodeDeleteProductKeluarRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return domain.DeleteProductRequest{ID: id}, nil
}
