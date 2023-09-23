package transport

import (
	"context"
	"encoding/json"
	"go_kit_inventory/internal/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func DecodeCreateSupplierRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req domain.CreateSupplierRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeResultsSupplierRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodeResultSupplierRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return domain.ResultSupplierRequest{ID: id}, nil
}

func DecodeUpdateSupplierRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req domain.UpdateSupplierRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeSupplierResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func DecodeDeleteSupplierRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return domain.DeleteSupplierRequest{ID: id}, nil
}
