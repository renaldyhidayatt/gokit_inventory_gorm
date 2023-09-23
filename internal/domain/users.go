package domain

import (
	"go_kit_inventory/internal/models"

	"github.com/go-playground/validator/v10"
)

type ResultsUserRequest struct{}

type ResultUserRequest struct {
	ID string
}

type DeleteUserRequest struct {
	ID string
}

type UpdateUserRequest struct {
	ID        string `json:"id" validate:"uuid"`
	FirstName string `json:"first_name" validate:"required,lowercase"`
	LastName  string `json:"last_name" validate:"required,lowercase"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
	Role      string `json:"role" validate:"required,lowercase"`
}

func (c *UpdateUserRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type RegisterResponse struct {
	User *models.ModelUser `json:"user,omitempty"`
}

type LoginResponse struct {
	Token Token `json:"token,omitempty"`
}

type ResultsResponse struct {
	Users *[]models.ModelUser `json:"users,omitempty"`
}

type ResultResponse struct {
	User *models.ModelUser `json:"user,omitempty"`
}

type UpdateResponse struct {
	User *models.ModelUser `json:"user,omitempty"`
}

type DeleteUserResponse struct {
	User *models.ModelUser `json:"user,omitempty"`
}
