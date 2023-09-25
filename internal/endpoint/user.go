package endpoint

import (
	"context"
	"go_kit_inventory/internal/domain/requests"
	"go_kit_inventory/internal/domain/responses"
	"go_kit_inventory/internal/models"
	"go_kit_inventory/internal/schema"
	"go_kit_inventory/internal/service"
	"net/http"

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

func UserEndpoints(s service.UserService) UserEndpoint {
	return UserEndpoint{
		RegisterEndpoint: makeRegisterEndpoint(s),
		LoginEndpoint:    makeLoginEndpoint(s),
		ResultsEndpoint:  makeResultsUserEndpoint(s),
		ResultEndpoint:   makeResultUserEndpoint(s),
		DeleteEndpoint:   makeDeleteUserEndpoint(s),
		UpdateEndpoint:   makeUpdateUserEndpoint(s),
	}
}

func makeRegisterEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.RegisterRequest)

		if err := req.Validate(); err != nil {
			return responses.RegisterUserResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		user, err := s.Register(&req)

		var response responses.RegisterUserResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to register user"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convertedUser := convertToUser(user)
			response.User = convertedUser
			response.Message = "User registered successfully"
			response.StatusCode = http.StatusCreated
		}

		return response, nil
	}
}

func makeLoginEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.LoginRequest)

		if err := req.Validate(); err != nil {
			return responses.LoginUserResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		token, err := s.Login(&req)

		var response responses.LoginUserResponse

		if err != nil {
			response.Error = err
			response.Message = "Login failed"
			response.StatusCode = http.StatusUnauthorized
		} else {
			response.Token = token
			response.Message = "Login successful"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeResultsUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.Results()

		var response responses.ResultsUserResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch users"
			response.StatusCode = http.StatusInternalServerError
		} else {

			var convertedUsers []schema.User
			for _, modelUser := range *users {
				convertedUser := convertToUser(&modelUser)
				convertedUsers = append(convertedUsers, *convertedUser)
			}
			response.Users = &convertedUsers
			response.Message = "Users fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeResultUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.ResultUserRequest)

		if err := req.Validate(); err != nil {
			return responses.ResultUserResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		user, err := s.Result(req.ID)

		var response responses.ResultUserResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to fetch user"
			response.StatusCode = http.StatusNotFound
		} else {
			convertedUser := convertToUser(user)
			response.User = convertedUser
			response.Message = "User fetched successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeDeleteUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.DeleteUserRequest)
		result, err := s.Delete(req.ID)

		var response responses.DeleteUserResponse

		if err := req.Validate(); err != nil {
			return responses.DeleteUserResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		if err != nil {
			response.Error = err
			response.Message = "Failed to delete user"
			response.StatusCode = http.StatusInternalServerError
		} else {
			convert := convertToUser(result)
			response.User = convert
			response.Message = "User deleted successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func makeUpdateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.UpdateUserRequest)

		if err := req.Validate(); err != nil {
			return responses.UpdateUserResponse{
				Error:      err,
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		user, err := s.Update(&req)

		var response responses.UpdateUserResponse

		if err != nil {
			response.Error = err
			response.Message = "Failed to update user"
			response.StatusCode = http.StatusInternalServerError
		} else {

			convertedUser := convertToUser(user)
			response.User = convertedUser
			response.Message = "User updated successfully"
			response.StatusCode = http.StatusOK
		}

		return response, nil
	}
}

func convertToUser(modelUser *models.ModelUser) *schema.User {
	return &schema.User{
		ID:        modelUser.ID,
		FirstName: modelUser.FirstName,
		LastName:  modelUser.LastName,
		Email:     modelUser.Email,
		Password:  modelUser.Password,
		Role:      modelUser.Role,
		CreatedAt: modelUser.CreatedAt,
		UpdatedAt: modelUser.UpdatedAt,
	}
}
