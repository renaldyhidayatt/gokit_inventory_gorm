package transport

import (
	"context"
	"encoding/json"
	"go_kit_inventory/internal/domain/requests"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func RegisterRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func RegisterResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func LoginRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func LoginResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}
func ResultsRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return requests.ResultsUserRequest{}, nil
}

func ResultsResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func ResultRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return requests.ResultUserRequest{ID: id}, nil
}

func ResultResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, response)
}

func DeleteRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return requests.DeleteUserRequest{ID: id}, nil
}

func UpdateRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.UpdateUserRequest
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
