package httpTransoprt

import (
	"local/endpoint"

	"github.com/go-chi/chi/v5"
)

func handleRouter(r chi.Router, endpoints *endpoint.Endpoints) chi.Router {
	h := &handler{endpoints: endpoints}

	// Health check
	r.Get("/health", h.HealthCheck)

	// API routes
	r.Route("/api", func(r chi.Router) {
		// Product endpoints
		r.Get("/product", h.GetAllProducts)
		r.Get("/product/{id}", h.GetProductById)

		// Place an order
		r.Post("/order", h.PlaceAnOrder)

	})

	return r
}
