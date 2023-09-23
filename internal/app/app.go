package app

import (
	"context"
	"fmt"
	"go_kit_inventory/internal/repository"
	"go_kit_inventory/internal/routes"
	"go_kit_inventory/internal/service"
	"go_kit_inventory/pkg/database/postgres"
	"go_kit_inventory/pkg/dotenv"
	"go_kit_inventory/pkg/logger"
	"go_kit_inventory/pkg/security"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func RunApp() {
	logger.Info("Running Server running")

	err := dotenv.Viper()

	if err != nil {
		logger.Fatal("Failed to listen: %v", zap.Error(err))
	}

	db, err := postgres.NewClient()

	if err != nil {
		logger.Error("error : ", zap.Error(err))
	}

	hashing := security.NewHashingPassword()
	repository := repository.NewRepository(db)
	token, err := security.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		logger.Error("error : ", zap.Error(err))
	}

	services := service.NewService(service.Deps{
		Repository: repository,
		Hashing:    *hashing,
		Token:      token,
	})

	handler := routes.NewHandlerRoute(services)

	r := chi.NewRouter()

	handler.NewAuthRoutes("/api/auth", db, r)
	handler.NewUserRoutes("/api/user", db, r)
	handler.NewCategoryRoutes("/api/category", db, r)
	handler.NewCustomerRoutes("/api/customer", db, r)
	handler.NewProductRoutes("/api/product", db, r)
	handler.NewProductKeluarRoutes("/api/productkeluar", db, r)
	handler.NewProductMasukRoutes("/api/productmasuk", db, r)
	handler.NewSaleRoutes("/api/sale", db, r)
	handler.NewSupplierRoutes("/api/supplier", db, r)

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("PORT")),
		WriteTimeout: time.Duration(viper.GetInt("WRITETIMEOUT")) * 10,
		ReadTimeout:  time.Duration(viper.GetInt("READTIMEOUT")) * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		err := serve.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Connected to port:", viper.GetString("PORT"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serve.Shutdown(ctx)
	os.Exit(0)
}
