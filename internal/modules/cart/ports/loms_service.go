package ports

import (
	"context"
	"errors"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
)

var (
	ErrInternalLomsService = errors.New("loms service error")
)

type LomsService interface {
	StockInfo(ctx context.Context, sku domain.SKU) (domain.StockInfo, error)
}
