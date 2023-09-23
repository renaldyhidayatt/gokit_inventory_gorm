package service

import (
	"go_kit_inventory/internal/repository"
	"go_kit_inventory/pkg/security"
)

type Service struct {
	User          UserService
	Category      CategoryService
	Customer      CustomerService
	Product       ProductService
	ProductKeluar ProductKeluarService
	ProductMasuk  ProductMasukService
	Sale          SaleService
	Supplier      SupplierService
}

type Deps struct {
	Repository *repository.Repositories
	Hashing    security.Hashing
	Token      security.TokenManager
}

func NewService(deps Deps) *Service {
	return &Service{
		User:          NewServiceUser(deps.Repository.User, deps.Hashing, deps.Token),
		Category:      NewServiceCategory(deps.Repository.Category),
		Customer:      NewServiceCustomer(deps.Repository.Customer),
		Product:       NewServiceProduct(deps.Repository.Product),
		ProductKeluar: NewServiceProductKeluar(deps.Repository.ProductKeluar),
		ProductMasuk:  NewServiceProductMasuk(deps.Repository.ProductMasuk),
		Sale:          NewServiceSale(deps.Repository.SaleRepository),
		Supplier:      NewServiceSupplier(deps.Repository.SupplierRepository),
	}
}
