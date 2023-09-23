package routes

import (
	"go_kit_inventory/internal/endpoint"
	"go_kit_inventory/pkg/transport"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func (h *HandlerRoute) NewProductKeluarRoutes(prefix string, db *gorm.DB, router *chi.Mux) {

	endpoints := endpoint.NewProductKeluarEndpoints(h.services.ProductKeluar)

	router.Route(prefix, func(r chi.Router) {
		r.Method(http.MethodPost, "/create", httptransport.NewServer(
			endpoints.CreateProductKeluarEndpoint,
			transport.DecodeCreateProductKeluarRequest,
			transport.EncodeProductKeluarResponse,
		))

		r.Method(http.MethodGet, "/results", httptransport.NewServer(
			endpoints.ResultsProductKeluarEndpoint,
			transport.DecodeResultsProductKeluarRequest,
			transport.EncodeProductKeluarResponse,
		))

		r.Method(http.MethodGet, "/result/{id}", httptransport.NewServer(
			endpoints.ResultProductKeluarEndpoint,
			transport.DecodeResultProductKeluarRequest,
			transport.EncodeProductKeluarResponse,
		))

		r.Method(http.MethodDelete, "/delete/{id}", httptransport.NewServer(
			endpoints.DeleteProductKeluarEndpoint,
			transport.DecodeDeleteProductKeluarRequest,
			transport.EncodeProductKeluarResponse,
		))

		r.Method(http.MethodPut, "/update", httptransport.NewServer(
			endpoints.UpdateProductKeluarEndpoint,
			transport.DecodeUpdateProductKeluarRequest,
			transport.EncodeProductKeluarResponse,
		))

	})
}
