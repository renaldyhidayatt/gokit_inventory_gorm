package requests

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

func (l *LoginRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}

type RegisterRequest struct {
	FirstName string `json:"firstname" validate:"required,lowercase"`
	LastName  string `json:"lastname" validate:"required,lowercase"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
	Role      string `json:"role" validate:"required,lowercase"`
}

func (c *RegisterRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
