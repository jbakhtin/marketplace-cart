package postgres

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
)

type CartStorage struct {
}

func NewCartStorage() (CartStorage, error) {
	return CartStorage{}, nil
}

func (c *CartStorage) AddItem(
	ctx context.Context,
	userID domain.UserID,
	sku domain.SKU,
	count domain.Count,
) error {
	return nil
}

func (c *CartStorage) GetCartByUserID(ctx context.Context, userID domain.UserID) (domain.Cart, error) {
	return domain.Cart{}, nil
}

func (c *CartStorage) DeleteItemBySKU(ctx context.Context, userID domain.UserID) error {
	return nil
}
func (c *CartStorage) DeleteItemsByUserID(ctx context.Context, userID domain.UserID) error {
	return nil
}
