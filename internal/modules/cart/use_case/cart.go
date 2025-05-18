package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
)

type OrderUseCase struct {
	logger          ports.Logger
	orderRepository ports.OrderRepository
}

func NewOrderUseCase(
	logger ports.Logger,
	orderRepository ports.OrderRepository,
) (OrderUseCase, error) {
	return OrderUseCase{
		logger:          logger,
		orderRepository: orderRepository,
	}, nil
}

func (o *OrderUseCase) AddItem(ctx context.Context, ID domain.OrderID) error {
	return nil
}
