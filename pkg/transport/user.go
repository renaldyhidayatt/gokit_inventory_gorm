package transport

import (
	"context"
	"encoding/json"
	"go_kit_inventory/internal/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func RegisterRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func RegisterResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func LoginRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func LoginResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}
func ResultsRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return domain.ResultsUserRequest{}, nil
}

func ResultsResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func ResultRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return domain.ResultUserRequest{ID: id}, nil
}

func ResultResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func DeleteRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return domain.DeleteUserRequest{ID: id}, nil
}

func UpdateRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.UpdateSupplierRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func UpdateResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func DeleteResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}
