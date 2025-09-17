package product

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
)

type NoOpAdapter struct{}

func (n NoOpAdapter) GetProduct(ctx context.Context, sku domain.SKU) (domain.ProductInfo, error) {
	return domain.ProductInfo{
		Name:  "Some Product",
		Price: domain.Price(100),
	}, nil
}
