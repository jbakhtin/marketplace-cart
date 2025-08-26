package ports

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
)

type Product struct {
}

type ProductService interface {
	GetProduct(ctx context.Context, sku domain.SKU)
}
