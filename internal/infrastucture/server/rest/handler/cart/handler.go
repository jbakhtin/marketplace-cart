package cart

import (
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/use_case"
)

type Config interface {
}

type Handler struct {
	cfg     Config
	log     ports.Logger
	useCase use_case.CartUseCase
}

func NewOrderHandler(cfg Config, logger ports.Logger, useCase use_case.CartUseCase) (Handler, error) {
	return Handler{
		cfg:     cfg,
		log:     logger,
		useCase: useCase,
	}, nil
}
