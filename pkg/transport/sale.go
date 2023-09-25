package transport

import (
	"context"
	"encoding/json"
	"go_kit_inventory/internal/domain/requests"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func DecodeCreateSaleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.CreateSaleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeResultsSaleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodeResultSaleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return requests.ResultSaleRequest{ID: id}, nil
}

func DecodeUpdateSaleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requests.UpdateSaleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeSaleResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func DecodeDeleteSaleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return requests.DeleteSaleRequest{ID: id}, nil
}
