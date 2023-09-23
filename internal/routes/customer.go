package routes

import (
	"go_kit_inventory/internal/endpoint"
	"go_kit_inventory/pkg/transport"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func (h *HandlerRoute) NewCustomerRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	endpoint := endpoint.NewCustomerEndpoints(h.services.Customer)

	router.Route(prefix, func(r chi.Router) {
		r.Method(http.MethodPost, "/create", httptransport.NewServer(
			endpoint.CreateEndpoint,
			transport.DecodeCreateCustomerRequest,
			transport.EncodeCustomerResponse,
		))

		r.Method(http.MethodGet, "/", httptransport.NewServer(
			endpoint.ResultsEndpoint,
			transport.DecodeResultsCustomerRequest,
			transport.EncodeCustomerResponse,
		))

		r.Method(http.MethodGet, "/{id}", httptransport.NewServer(
			endpoint.ResultEndpoint,
			transport.DecodeResultCustomerRequest,
			transport.EncodeCustomerResponse,
		))

		r.Method(http.MethodDelete, "/delete/{id}", httptransport.NewServer(
			endpoint.DeleteEndpoint,
			transport.DecodeDeleteCustomerRequest,
			transport.EncodeCustomerResponse,
		))

		r.Method(http.MethodPut, "/update", httptransport.NewServer(
			endpoint.UpdateEndpoint,
			transport.DecodeUpdateCustomerRequest,
			transport.EncodeCustomerResponse,
		))
	})

}
