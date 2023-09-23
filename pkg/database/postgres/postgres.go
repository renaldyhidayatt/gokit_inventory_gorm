package postgres

import (
	"fmt"
	"go_kit_inventory/internal/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed error connect to database")
	}

	err = db.AutoMigrate(&models.ModelCategory{}, &models.ModelCustomer{}, &models.ModelProduct{}, &models.ModelProductKeluar{}, &models.ModelProductMasuk{}, &models.ModelSale{}, &models.ModelSupplier{}, &models.ModelUser{})

	if err != nil {

		return nil, fmt.Errorf("faield connected to database")
	}

	return db, nil
}
