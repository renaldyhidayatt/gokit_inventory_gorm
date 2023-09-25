package service

import (
	"go_kit_inventory/internal/domain/requests"
	"go_kit_inventory/internal/models"
	"go_kit_inventory/internal/repository"
)

type ServiceProductKeluar struct {
	Repository repository.ProductKeluarRepository
}

func NewServiceProductKeluar(productkeluar repository.ProductKeluarRepository) *ServiceProductKeluar {
	return &ServiceProductKeluar{Repository: productkeluar}
}

func (s *ServiceProductKeluar) Create(input *requests.CreateProductKeluarRequest) (*models.ModelProductKeluar, error) {
	var productkeluar requests.CreateProductKeluarRequest

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	res, err := s.Repository.Create(&productkeluar)

	return res, err
}

func (s *ServiceProductKeluar) Results() (*[]models.ModelProductKeluar, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *ServiceProductKeluar) Result(id string) (*models.ModelProductKeluar, error) {

	res, err := s.Repository.Result(id)

	return res, err
}

func (s *ServiceProductKeluar) Delete(id string) (*models.ModelProductKeluar, error) {

	res, err := s.Repository.Delete(id)

	return res, err
}

func (s *ServiceProductKeluar) Update(input *requests.UpdateProductKeluarRequest) (*models.ModelProductKeluar, error) {
	var productkeluar requests.UpdateProductKeluarRequest

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	res, err := s.Repository.Update(&productkeluar)

	return res, err
}
