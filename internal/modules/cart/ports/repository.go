package ports

import (
	"context"
	"errors"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
)

var (
	ErrInternalRepository = errors.New("cart repository error")
)

type CartRepository interface {
	AddItem(ctx context.Context, userID domain.UserID, sku domain.SKU, count domain.Count) error
	GetCartByUserID(ctx context.Context, userID domain.UserID) (domain.Cart, error)
	DeleteItemBySKU(ctx context.Context, userID domain.UserID) error
	DeleteItemsByUserID(ctx context.Context, userID domain.UserID) error
}
