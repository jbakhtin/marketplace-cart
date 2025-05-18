package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
)

type CartUseCase struct {
	logger          ports.Logger
	orderRepository ports.CartRepository
}

func NewCartUseCase(
	logger ports.Logger,
	orderRepository ports.CartRepository,
) (CartUseCase, error) {
	return CartUseCase{
		logger:          logger,
		orderRepository: orderRepository,
	}, nil
}

func (o *CartUseCase) AddItem(ctx context.Context, item domain.Item) error {
	return nil
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
