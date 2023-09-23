package routes

import (
	"go_kit_inventory/internal/endpoint"
	"go_kit_inventory/pkg/transport"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func (h *HandlerRoute) NewProductRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	endpoint := endpoint.NewProductEndpoints(h.services.Product)

	router.Route(prefix, func(r chi.Router) {
		r.Method(http.MethodPost, "/create", httptransport.NewServer(
			endpoint.CreateEndpoint,
			transport.DecodeCreateProductRequest,
			transport.EncodeProductResponse,
		))

		r.Method(http.MethodGet, "/results", httptransport.NewServer(
			endpoint.ResultsEndpoint,
			transport.DecodeResultsProductRequest,
			transport.EncodeProductResponse,
		))

		r.Method(http.MethodGet, "/result/{id}", httptransport.NewServer(
			endpoint.ResultEndpoint,
			transport.DecodeResultProductRequest,
			transport.EncodeProductResponse,
		))

		r.Method(http.MethodDelete, "/delete/{id}", httptransport.NewServer(
			endpoint.DeleteEndpoint,
			transport.DecodeDeleteProductRequest,
			transport.EncodeProductResponse,
		))

		r.Method(http.MethodPut, "/update", httptransport.NewServer(
			endpoint.UpdateEndpoint,
			transport.DecodeUpdateProductRequest,
			transport.EncodeProductResponse,
		))
	})

}
