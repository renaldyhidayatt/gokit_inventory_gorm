package routes

import (
	"go_kit_inventory/internal/endpoint"
	"go_kit_inventory/pkg/transport"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func (h *HandlerRoute) NewProductMasukRoutes(prefix string, db *gorm.DB, router *chi.Mux) {

	endpoints := endpoint.NewProductMasukEndpoints(h.services.ProductMasuk)

	router.Route(prefix, func(r chi.Router) {
		r.Method(http.MethodPost, "/create", httptransport.NewServer(
			endpoints.CreateProductMasukEndpoint,
			transport.DecodeCreateProductKeluarRequest,
			transport.EncodeProductMasukResponse,
		))

		r.Method(http.MethodGet, "/results", httptransport.NewServer(
			endpoints.ResultsProductMasukEndpoint,
			transport.DecodeResultsProductMasukRequest,
			transport.EncodeProductMasukResponse,
		))

		r.Method(http.MethodGet, "/result/{id}", httptransport.NewServer(
			endpoints.ResultProductMasukEndpoint,
			transport.DecodeResultProductMasukRequest,
			transport.EncodeProductMasukResponse,
		))

		r.Method(http.MethodDelete, "/delete/{id}", httptransport.NewServer(
			endpoints.DeleteProductMasukEndpoint,
			transport.DecodeDeleteProductMasukRequest,
			transport.EncodeProductMasukResponse,
		))

		r.Method(http.MethodPut, "/update", httptransport.NewServer(
			endpoints.UpdateProductMasukEndpoint,
			transport.DecodeUpdateProductMasukRequest,
			transport.EncodeProductMasukResponse,
		))

	})
}
