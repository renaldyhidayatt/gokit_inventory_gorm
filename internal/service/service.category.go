package service

import (
	"go_kit_inventory/internal/domain/requests"
	"go_kit_inventory/internal/models"
	"go_kit_inventory/internal/repository"
)

type ServiceCategory struct {
	Repository repository.CategoryRepository
}

func NewServiceCategory(category repository.CategoryRepository) *ServiceCategory {
	return &ServiceCategory{Repository: category}
}

func (s *ServiceCategory) Create(input *requests.CreateCategoryRequest) (*models.ModelCategory, error) {
	var category requests.CreateCategoryRequest

	category.Name = input.Name

	res, err := s.Repository.Create(&category)
	return res, err
}

func (s *ServiceCategory) Results() (*[]models.ModelCategory, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *ServiceCategory) Result(id string) (*models.ModelCategory, error) {

	res, err := s.Repository.Result(id)

	return res, err
}

func (s *ServiceCategory) Delete(id string) (*models.ModelCategory, error) {

	res, err := s.Repository.Delete(id)

	return res, err
}

func (s *ServiceCategory) Update(input *requests.UpdateCategoryRequest) (*models.ModelCategory, error) {
	var category requests.UpdateCategoryRequest

	category.ID = input.ID
	category.Name = input.Name

	res, err := s.Repository.Update(&category)

	return res, err
}
