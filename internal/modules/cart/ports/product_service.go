package ports

import (
	"context"
	"errors"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
)

var (
	ErrInternalProductService = errors.New("internal product service error")
	ErrProductNotFound        = errors.New("product not found")
)

type ProductService interface {
	GetProduct(ctx context.Context, sku domain.SKU) (domain.ProductInfo, error)
}
