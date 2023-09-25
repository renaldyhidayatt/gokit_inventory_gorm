package transport

import (
	"context"
	"encoding/json"
	"errors"
	"go_kit_inventory/internal/domain/requests"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func DecodeCreateProductMasukRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.CreateProductMasukRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeResultsProductMasukRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodeResultProductMasukRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, errors.New("missing 'id' parameter")
	}
	return requests.ResultProductMasukRequest{ID: id}, nil
}

func DecodeDeleteProductMasukRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, errors.New("missing 'id' parameter")
	}
	return requests.DeleteProductMasukRequest{ID: id}, nil
}

func DecodeUpdateProductMasukRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.UpdateProductMasukRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeProductMasukResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}
