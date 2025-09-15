package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/server/rest/handler/cart"
	custommiddleware "github.com/jbakhtin/marketplace-cart/internal/infrastucture/server/rest/middleware"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/use_case"
)

type Config interface {
	GetAppKey() string
}

func NewRouter(
	cfg Config,
	logger ports.Logger,
	cartUseCase use_case.CartUseCase,
) (*chi.Mux, error) {
	cartHandler, err := cart.NewOrderHandler(logger, cartUseCase)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)

	// custom middleware
	authMiddleware := custommiddleware.NewAuthMiddleware(cfg)
	router.Use(authMiddleware.Auth)

	router.Route("/cart", func(r chi.Router) {
		r.Get("/list", cartHandler.List)
		r.Post("/checkout", cartHandler.Checkout)
		r.Post("/clear", cartHandler.Clear)

		r.Route("/items", func(r chi.Router) {
			r.Post("/add", cartHandler.AddItem)
			r.Post("/delete", cartHandler.Delete)
		})
	})

	return router, nil
}
