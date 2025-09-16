package cart

import (
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/use_case"
)

type Handler struct {
	log     ports.Logger
	useCase use_case.CartUseCaseInterface
}

func NewHandler(logger ports.Logger, useCase use_case.CartUseCaseInterface) (Handler, error) {
	return Handler{
		log:     logger,
		useCase: useCase,
	}, nil
}
