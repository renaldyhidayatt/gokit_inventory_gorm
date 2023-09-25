package routes

import (
	"go_kit_inventory/internal/endpoint"
	"go_kit_inventory/internal/middleware"
	"go_kit_inventory/pkg/transport"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func (h *HandlerRoute) NewUserRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	endpoint := endpoint.UserEndpoints(h.services.User)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)

		r.Method(http.MethodPost, "/register", httptransport.NewServer(
			endpoint.RegisterEndpoint,
			transport.RegisterRequestDecoder,
			transport.RegisterResponseEncoder,
		))
		r.Method(http.MethodPost, "/login", httptransport.NewServer(
			endpoint.LoginEndpoint,
			transport.LoginRequestDecoder,
			transport.LoginResponseEncoder,
		))
		r.Method(http.MethodGet, "/", httptransport.NewServer(
			endpoint.ResultsEndpoint,
			transport.ResultsRequestDecoder,
			transport.ResultsResponseEncoder,
		))
		r.Method(http.MethodGet, "/{id}", httptransport.NewServer(
			endpoint.ResultEndpoint,
			transport.ResultRequestDecoder,
			transport.ResultResponseEncoder,
		))
		r.Method(http.MethodDelete, "/{id}", httptransport.NewServer(
			endpoint.DeleteEndpoint,
			transport.DeleteRequestDecoder,
			transport.DeleteResponseEncoder,
		))
		r.Method(http.MethodPut, "/{id}", httptransport.NewServer(
			endpoint.UpdateEndpoint,
			transport.UpdateRequestDecoder,
			transport.UpdateResponseEncoder,
		))
	})
}
