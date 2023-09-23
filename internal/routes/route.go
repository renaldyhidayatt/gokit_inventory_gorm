package routes

import "go_kit_inventory/internal/service"

type HandlerRoute struct {
	services *service.Service
}

func NewHandlerRoute(services *service.Service) *HandlerRoute {
	

	return &HandlerRoute{services: services}
}
