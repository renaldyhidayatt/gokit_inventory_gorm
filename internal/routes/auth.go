package routes

import (
	"go_kit_inventory/internal/endpoint"
	"go_kit_inventory/pkg/transport"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func (h *HandlerRoute) NewAuthRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	endpoint := endpoint.UserEndpoints(h.services.User)

	router.Route(prefix, func(r chi.Router) {
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

	})
}
