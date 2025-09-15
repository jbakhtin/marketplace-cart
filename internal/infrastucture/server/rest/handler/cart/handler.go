package cart

import (
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/use_case"
)

type Handler struct {
	log     ports.Logger
	useCase use_case.CartUseCase
}

func NewOrderHandler(logger ports.Logger, useCase use_case.CartUseCase) (Handler, error) {
	return Handler{
		log:     logger,
		useCase: useCase,
	}, nil
}
