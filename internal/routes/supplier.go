package routes

import (
	"go_kit_inventory/internal/endpoint"
	"go_kit_inventory/pkg/transport"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func (h *HandlerRoute) NewSupplierRoutes(prefix string, db *gorm.DB, router *chi.Mux) {

	endpoints := endpoint.NewSupplierEndpoints(h.services.Supplier)

	router.Route(prefix, func(r chi.Router) {
		r.Method(http.MethodPost, "/create", httptransport.NewServer(
			endpoints.CreateEndpoint,
			transport.DecodeCreateSupplierRequest,
			transport.EncodeSupplierResponse,
		))

		r.Method(http.MethodGet, "/results", httptransport.NewServer(
			endpoints.ResultsEndpoint,
			transport.DecodeResultsSupplierRequest,
			transport.EncodeSupplierResponse,
		))

		r.Method(http.MethodGet, "/result/{id}", httptransport.NewServer(
			endpoints.ResultEndpoint,
			transport.DecodeResultSupplierRequest,
			transport.EncodeSupplierResponse,
		))

		r.Method(http.MethodDelete, "/delete/{id}", httptransport.NewServer(
			endpoints.DeleteEndpoint,
			transport.DecodeDeleteSupplierRequest,
			transport.EncodeSupplierResponse,
		))

		r.Method(http.MethodPut, "/update", httptransport.NewServer(
			endpoints.UpdateEndpoint,
			transport.DecodeUpdateSupplierRequest,
			transport.EncodeSupplierResponse,
		))
	})
}
