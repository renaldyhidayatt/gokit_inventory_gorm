package endpoint

import (
	"context"
	"go_kit_inventory/internal/domain"
	"go_kit_inventory/internal/service"

	"github.com/go-kit/kit/endpoint"
)

type UserEndpoint struct {
	RegisterEndpoint endpoint.Endpoint
	LoginEndpoint    endpoint.Endpoint
	ResultsEndpoint  endpoint.Endpoint
	ResultEndpoint   endpoint.Endpoint
	DeleteEndpoint   endpoint.Endpoint
	UpdateEndpoint   endpoint.Endpoint
}

func MakeUserEndpoints(s service.UserService) UserEndpoint {
	return UserEndpoint{
		RegisterEndpoint: MakeRegisterEndpoint(s),
		LoginEndpoint:    MakeLoginEndpoint(s),
		ResultsEndpoint:  MakeResultsEndpoint(s),
		ResultEndpoint:   MakeResultEndpoint(s),
		DeleteEndpoint:   MakeDeleteEndpoint(s),
		UpdateEndpoint:   MakeUpdateEndpoint(s),
	}
}

func MakeRegisterEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.RegisterInput)
		user, err := s.Register(&req)
		if err != nil {
			return nil, err
		}
		return domain.RegisterResponse{User: user}, nil
	}
}

func MakeLoginEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.LoginInput)
		token, err := s.Login(&req)
		if err != nil {
			return nil, err
		}
		return domain.LoginResponse{Token: token}, nil
	}
}

func MakeResultsEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.Results()
		if err != nil {
			return nil, err
		}
		return domain.ResultsResponse{Users: users}, nil
	}
}

func MakeResultEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.ResultUserRequest)
		user, err := s.Result(req.ID)
		if err != nil {
			return nil, err
		}
		return domain.ResultResponse{User: user}, nil
	}
}

func MakeDeleteEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.DeleteUserRequest)
		result, err := s.Delete(req.ID)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}

func MakeUpdateEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.UpdateUserRequest)
		user, err := s.Update(&req)
		if err != nil {
			return nil, err
		}
		return domain.UpdateResponse{User: user}, nil
	}
}
