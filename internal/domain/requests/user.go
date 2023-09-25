package requests

import "github.com/go-playground/validator/v10"

type ResultsUserRequest struct{}

type ResultUserRequest struct {
	ID string `json:"id" validate:"uuid"`
}

type DeleteUserRequest struct {
	ID string `json:"id" validate:"uuid"`
}

func (c *ResultUserRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

func (c *DeleteUserRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
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
