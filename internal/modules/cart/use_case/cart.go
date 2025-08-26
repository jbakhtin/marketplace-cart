package use_case

import (
	"context"
	customContext "github.com/jbakhtin/marketplace-cart/internal/infrastucture/context"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
)

type CartUseCase struct {
	logger         ports.Logger
	cartRepository ports.CartRepository
}

func NewCartUseCase(
	logger ports.Logger,
	cartRepository ports.CartRepository,
) (CartUseCase, error) {
	return CartUseCase{
		logger:         logger,
		cartRepository: cartRepository,
	}, nil
}

func (o *CartUseCase) AddItem(ctx context.Context, item domain.Item) error {
	userID := ctx.Value(customContext.UserIDKey)
	return o.cartRepository.AddItem(ctx, userID.(domain.UserID), item.Sku, item.Count)
}

func (o *CartUseCase) DeleteItem(ctx context.Context, item domain.SKU) error {
	return nil
}

func (o *CartUseCase) List(ctx context.Context) (domain.Cart, error) {
	return domain.Cart{}, nil
}

func (o *CartUseCase) Clear(ctx context.Context) error {
	return nil
}

func (o *CartUseCase) Checkout(ctx context.Context) error {
	return nil
}
