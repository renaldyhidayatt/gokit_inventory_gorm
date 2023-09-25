package transport

import (
	"context"
	"encoding/json"
	"go_kit_inventory/internal/domain/requests"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func DecodeCreateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeResultsProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodeResultProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")

	return requests.ResultProductRequest{ID: id}, nil
}

func DecodeUpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeProductResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func DecodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return requests.DeleteProductRequest{ID: id}, nil
}
