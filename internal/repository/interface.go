package repository

import (
	"go_kit_inventory/internal/domain/requests"
	"go_kit_inventory/internal/models"
)

type UserRepository interface {
	Register(input *requests.RegisterRequest) (*models.ModelUser, error)
	Login(input *requests.LoginRequest) (*models.ModelUser, error)
	Results() (*[]models.ModelUser, error)
	Result(id string) (*models.ModelUser, error)
	Delete(id string) (*models.ModelUser, error)
	Update(input *requests.UpdateUserRequest) (*models.ModelUser, error)
	FindByEmail(email string) (*models.ModelUser, error)
}

type CategoryRepository interface {
	Create(input *requests.CreateCategoryRequest) (*models.ModelCategory, error)
	Results() (*[]models.ModelCategory, error)
	Result(id string) (*models.ModelCategory, error)
	Delete(id string) (*models.ModelCategory, error)
	Update(input *requests.UpdateCategoryRequest) (*models.ModelCategory, error)
}

type CustomerRepository interface {
	Create(input *requests.CreateCustomerRequest) (*models.ModelCustomer, error)
	Results() (*[]models.ModelCustomer, error)
	Result(id string) (*models.ModelCustomer, error)
	Delete(od string) (*models.ModelCustomer, error)
	Update(input *requests.UpdateCustomerRequest) (*models.ModelCustomer, error)
}

type ProductRepository interface {
	Create(input *requests.CreateProductRequest) (*models.ModelProduct, error)
	Delete(id string) (*models.ModelProduct, error)
	Result(id string) (*models.ModelProduct, error)
	Update(input *requests.UpdateProductRequest) (*models.ModelProduct, error)
	Results() (*[]models.ModelProduct, error)
}

type ProductKeluarRepository interface {
	Create(input *requests.CreateProductKeluarRequest) (*models.ModelProductKeluar, error)
	Result(id string) (*models.ModelProductKeluar, error)
	Results() (*[]models.ModelProductKeluar, error)
	Delete(id string) (*models.ModelProductKeluar, error)
	Update(input *requests.UpdateProductKeluarRequest) (*models.ModelProductKeluar, error)
}

type ProductMasukRepository interface {
	Create(input *requests.CreateProductMasukRequest) (*models.ModelProductMasuk, error)
	Result(id string) (*models.ModelProductMasuk, error)
	Results() (*[]models.ModelProductMasuk, error)
	Delete(id string) (*models.ModelProductMasuk, error)
	Update(input *requests.UpdateProductMasukRequest) (*models.ModelProductMasuk, error)
}

type SaleRepository interface {
	Create(input *requests.CreateSaleRequest) (*models.ModelSale, error)
	Result(id string) (*models.ModelSale, error)
	Results() (*[]models.ModelSale, error)
	Delete(id string) (*models.ModelSale, error)
	Update(input *requests.UpdateSaleRequest) (*models.ModelSale, error)
}

type SupplierRepository interface {
	Create(input *requests.CreateSupplierRequest) (*models.ModelSupplier, error)
	Result(id string) (*models.ModelSupplier, error)
	Results() (*[]models.ModelSupplier, error)
	Delete(id string) (*models.ModelSupplier, error)
	Update(input *requests.UpdateSupplierRequest) (*models.ModelSupplier, error)
}
