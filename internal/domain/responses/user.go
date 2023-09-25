package responses

import (
	"go_kit_inventory/internal/schema"
)

type RegisterUserResponse struct {
	User       *schema.User `json:"user,omitempty"`
	Message    string       `json:"message,omitempty"`
	StatusCode int          `json:"status_code,omitempty"`
	Error      error        `json:"error,omitempty"`
}

type LoginUserResponse struct {
	Token      schema.Token `json:"token,omitempty"`
	Message    string       `json:"message,omitempty"`
	StatusCode int          `json:"status_code,omitempty"`
	Error      error        `json:"error,omitempty"`
}

type ResultsUserResponse struct {
	Users      *[]schema.User `json:"users,omitempty"`
	Message    string         `json:"message,omitempty"`
	StatusCode int            `json:"status_code,omitempty"`
	Error      error          `json:"error,omitempty"`
}

type ResultUserResponse struct {
	User       *schema.User `json:"user,omitempty"`
	Message    string       `json:"message,omitempty"`
	StatusCode int          `json:"status_code,omitempty"`
	Error      error        `json:"error,omitempty"`
}

type UpdateUserResponse struct {
	User       *schema.User `json:"user,omitempty"`
	Message    string       `json:"message,omitempty"`
	StatusCode int          `json:"status_code,omitempty"`
	Error      error        `json:"error,omitempty"`
}

type DeleteUserResponse struct {
	User       *schema.User `json:"user,omitempty"`
	Message    string       `json:"message,omitempty"`
	StatusCode int          `json:"status_code,omitempty"`
	Error      error        `json:"error,omitempty"`
}
