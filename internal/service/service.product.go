package service

import (
	"go_kit_inventory/internal/domain"
	"go_kit_inventory/internal/models"
	"go_kit_inventory/internal/repository"
)

type ServiceProduct struct {
	Repository repository.ProductRepository
}

func NewServiceProduct(product repository.ProductRepository) *ServiceProduct {
	return &ServiceProduct{Repository: product}
}

func (s *ServiceProduct) Create(input *domain.CreateProductRequest) (*models.ModelProduct, error) {
	var product domain.CreateProductRequest

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	res, err := s.Repository.Create(&product)

	return res, err
}

func (s *ServiceProduct) Results() (*[]models.ModelProduct, error) {
	res, err := s.Repository.Results()
	return res, err
}

func (s *ServiceProduct) Result(id string) (*models.ModelProduct, error) {

	res, err := s.Repository.Result(id)

	return res, err
}

func (s *ServiceProduct) Delete(id string) (*models.ModelProduct, error) {

	res, err := s.Repository.Delete(id)

	return res, err
}

func (s *ServiceProduct) Update(input *domain.UpdateProductRequest) (*models.ModelProduct, error) {
	var product domain.UpdateProductRequest

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	res, err := s.Repository.Update(&product)

	return res, err
}
