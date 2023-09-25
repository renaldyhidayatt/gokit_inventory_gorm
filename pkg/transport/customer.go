package transport

import (
	"context"
	"encoding/json"
	"go_kit_inventory/internal/domain/requests"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func DecodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeResultsCustomerRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodeResultCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return requests.ResultCustomerRequest{ID: id}, nil
}

func DecodeDeleteCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return requests.DeleteCustomerRequest{ID: id}, nil
}

func DecodeUpdateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.UpdateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeCustomerResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}
