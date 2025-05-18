package cart

import (
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/use_case"
)

type Module struct {
	cartUseCase use_case.CartUseCase
}

func InitModule(logger ports.Logger, cartRepository ports.CartRepository) (Module, error) {
	cartUseCase, err := use_case.NewCartUseCase(logger, cartRepository)
	if err != nil {
		return Module{}, err
	}

	logger.Debug("cart use case successful initiated")

	return Module{
		cartUseCase: cartUseCase,
	}, nil
}

func (m Module) GetCartUseCase() use_case.CartUseCase {
	return m.cartUseCase
}
