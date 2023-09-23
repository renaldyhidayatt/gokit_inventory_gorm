package routes

import (
	"go_kit_inventory/internal/endpoint"
	"go_kit_inventory/pkg/transport"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func (h *HandlerRoute) NewCategoryRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	endpoint := endpoint.NewCategoryEndpoints(h.services.Category)

	router.Route(prefix, func(r chi.Router) {
		r.Method(http.MethodPost, "/create", httptransport.NewServer(
			endpoint.CreateEndpoint,
			transport.DecodeCreateCategoryRequest,
			transport.EncodeCategoryResponse,
		))

		r.Method(http.MethodGet, "/", httptransport.NewServer(
			endpoint.ResultsEndpoint,
			transport.DecodeResultsCategoryRequest,
			transport.EncodeCategoryResponse,
		))

		r.Method(http.MethodGet, "/{id}", httptransport.NewServer(
			endpoint.ResultEndpoint,
			transport.DecodeResultCategoryRequest,
			transport.EncodeCategoryResponse,
		))

		r.Method(http.MethodDelete, "/delete/{id}", httptransport.NewServer(
			endpoint.DeleteEndpoint,
			transport.DecodeDeleteCategoryRequest,
			transport.EncodeCategoryResponse,
		))

		r.Method(http.MethodPut, "/update", httptransport.NewServer(
			endpoint.UpdateEndpoint,
			transport.DecodeUpdateCategoryRequest,
			transport.EncodeCategoryResponse,
		))
	})
}
