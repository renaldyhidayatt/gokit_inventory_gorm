package routes

import (
	"go_kit_inventory/internal/endpoint"
	"go_kit_inventory/pkg/transport"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func (h *HandlerRoute) NewSaleRoutes(prefix string, db *gorm.DB, router *chi.Mux) {

	endpoints := endpoint.NewSaleEndpoints(h.services.Sale)

	router.Route(prefix, func(r chi.Router) {
		r.Method(http.MethodPost, "/create", httptransport.NewServer(
			endpoints.CreateEndpoint,
			transport.DecodeCreateSaleRequest,
			transport.EncodeSaleResponse,
		))

		r.Method(http.MethodGet, "/results", httptransport.NewServer(
			endpoints.ResultsEndpoint,
			transport.DecodeResultsSaleRequest,
			transport.EncodeSaleResponse,
		))

		r.Method(http.MethodGet, "/result/{id}", httptransport.NewServer(
			endpoints.ResultEndpoint,
			transport.DecodeResultSaleRequest,
			transport.EncodeSaleResponse,
		))

		r.Method(http.MethodDelete, "/delete/{id}", httptransport.NewServer(
			endpoints.DeleteEndpoint,
			transport.DecodeDeleteSaleRequest,
			transport.EncodeSaleResponse,
		))

		r.Method(http.MethodPut, "/update", httptransport.NewServer(
			endpoints.UpdateEndpoint,
			transport.DecodeUpdateSaleRequest,
			transport.EncodeSaleResponse,
		))
	})
}
