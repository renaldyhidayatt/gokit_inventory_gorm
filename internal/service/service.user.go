package service

import (
	"fmt"
	"go_kit_inventory/internal/domain/requests"
	"go_kit_inventory/internal/models"
	"go_kit_inventory/internal/repository"
	"go_kit_inventory/internal/schema"
	"go_kit_inventory/pkg/security"
)

type ServiceUser struct {
	Repository repository.UserRepository
	Hash       security.Hashing
	Token      security.TokenManager
}

func NewServiceUser(auth repository.UserRepository, hash security.Hashing, token security.TokenManager) *ServiceUser {
	return &ServiceUser{Repository: auth, Hash: hash, Token: token}
}

func (s *ServiceUser) Register(input *requests.RegisterRequest) (*models.ModelUser, error) {

	var schema requests.RegisterRequest

	passwordHash, err := s.Hash.HashPassword(input.Password)

	if err != nil {
		return nil, err
	}

	schema.FirstName = input.FirstName
	schema.LastName = input.LastName
	schema.Email = input.Email
	schema.Password = passwordHash
	schema.Role = input.Role

	res, err := s.Repository.Register(&schema)

	return res, err
}

func (s *ServiceUser) Login(input *requests.LoginRequest) (schema.Token, error) {
	var request requests.LoginRequest

	res, err := s.Repository.FindByEmail(request.Email)

	if err != nil {
		return schema.Token{}, fmt.Errorf("failed error %w", err)
	}

	err = s.Hash.ComparePassword(res.Password, input.Password)

	if err != nil {
		return schema.Token{}, fmt.Errorf("failed error :%w", err)
	}

	request.Email = input.Email
	request.Password = res.Password

	res, err = s.Repository.Login(&request)

	if err != nil {
		return schema.Token{}, err
	}

	return s.createJwt(res.ID)
}

func (s *ServiceUser) createJwt(id string) (schema.Token, error) {
	var (
		res schema.Token
		err error
	)

	res.Jwt, err = s.Token.NewJWT(id)

	if err != nil {
		return res, nil
	}

	return res, nil
}

func (s *ServiceUser) Results() (*[]models.ModelUser, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *ServiceUser) Result(id string) (*models.ModelUser, error) {

	res, err := s.Repository.Result(id)

	return res, err
}

func (s *ServiceUser) Delete(id string) (*models.ModelUser, error) {

	res, err := s.Repository.Delete(id)

	return res, err
}

func (s *ServiceUser) Update(input *requests.UpdateUserRequest) (*models.ModelUser, error) {
	var user requests.UpdateUserRequest

	user.ID = input.ID

	res, err := s.Repository.Update(&user)

	return res, err
}
