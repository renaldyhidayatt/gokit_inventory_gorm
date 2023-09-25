package service

import (
	"go_kit_inventory/internal/domain/requests"
	"go_kit_inventory/internal/models"
	"go_kit_inventory/internal/schema"
)

type UserService interface {
	Register(input *requests.RegisterRequest) (*models.ModelUser, error)
	Login(input *requests.LoginRequest) (schema.Token, error)
	Results() (*[]models.ModelUser, error)
	Result(id string) (*models.ModelUser, error)
	Delete(id string) (*models.ModelUser, error)
	Update(input *requests.UpdateUserRequest) (*models.ModelUser, error)
}

type CategoryService interface {
	Create(input *requests.CreateCategoryRequest) (*models.ModelCategory, error)
	Results() (*[]models.ModelCategory, error)
	Result(id string) (*models.ModelCategory, error)
	Delete(id string) (*models.ModelCategory, error)
	Update(input *requests.UpdateCategoryRequest) (*models.ModelCategory, error)
}

type CustomerService interface {
	Create(input *requests.CreateCustomerRequest) (*models.ModelCustomer, error)
	Results() (*[]models.ModelCustomer, error)
	Result(id string) (*models.ModelCustomer, error)
	Delete(id string) (*models.ModelCustomer, error)
	Update(input *requests.UpdateCustomerRequest) (*models.ModelCustomer, error)
}

type ProductService interface {
	Create(input *requests.CreateProductRequest) (*models.ModelProduct, error)
	Delete(id string) (*models.ModelProduct, error)
	Result(id string) (*models.ModelProduct, error)
	Update(input *requests.UpdateProductRequest) (*models.ModelProduct, error)
	Results() (*[]models.ModelProduct, error)
}

type ProductKeluarService interface {
	Create(input *requests.CreateProductKeluarRequest) (*models.ModelProductKeluar, error)
	Result(id string) (*models.ModelProductKeluar, error)
	Results() (*[]models.ModelProductKeluar, error)
	Delete(id string) (*models.ModelProductKeluar, error)
	Update(input *requests.UpdateProductKeluarRequest) (*models.ModelProductKeluar, error)
}

type ProductMasukService interface {
	Create(input *requests.CreateProductMasukRequest) (*models.ModelProductMasuk, error)
	Result(id string) (*models.ModelProductMasuk, error)
	Results() (*[]models.ModelProductMasuk, error)
	Delete(id string) (*models.ModelProductMasuk, error)
	Update(input *requests.UpdateProductMasukRequest) (*models.ModelProductMasuk, error)
}

type SaleService interface {
	Create(input *requests.CreateSaleRequest) (*models.ModelSale, error)
	Result(id string) (*models.ModelSale, error)
	Results() (*[]models.ModelSale, error)
	Delete(id string) (*models.ModelSale, error)
	Update(input *requests.UpdateSaleRequest) (*models.ModelSale, error)
}

type SupplierService interface {
	Create(input *requests.CreateSupplierRequest) (*models.ModelSupplier, error)
	Result(id string) (*models.ModelSupplier, error)
	Results() (*[]models.ModelSupplier, error)
	Delete(id string) (*models.ModelSupplier, error)
	Update(input *requests.UpdateSupplierRequest) (*models.ModelSupplier, error)
}
