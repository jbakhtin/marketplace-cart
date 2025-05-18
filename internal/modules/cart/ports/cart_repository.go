package ports

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, order domain.Order) (domain.Order, error)
	SetStatus(ctx context.Context, ID int, status string) (domain.Order, error)
	GetByID(ctx context.Context, ID int) (domain.Order, error)
}
