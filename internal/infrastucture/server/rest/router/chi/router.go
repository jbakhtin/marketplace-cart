package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/server/rest/handler/cart"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/use_case"
)

type Config interface {
}

func NewRouter(
	cfg Config,
	logger ports.Logger,
	cartUseCase use_case.CartUseCase,
) (*chi.Mux, error) {
	cartHandler, err := cart.NewOrderHandler(cfg, logger, cartUseCase)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)

	router.Route("/orders", func(r chi.Router) {
		r.Post("/create", cartHandler.Create)

		r.Route("/{OrderID}", func(r chi.Router) {
			r.Get("/info", cartHandler.Info)
			r.Put("/pay", cartHandler.Pay)
			r.Put("/cancel", cartHandler.Cancel)
		})
	})

	router.Route("/stocks", func(r chi.Router) {
		r.Route("/{StockID}", func(r chi.Router) {
			r.Get("/info", cartHandler.Info)
		})
	})

	return router, nil
}
