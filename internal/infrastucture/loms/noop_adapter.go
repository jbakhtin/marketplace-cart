package loms

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
)

type NoOpAdapter struct{}

func (n NoOpAdapter) StockInfo(ctx context.Context, sku domain.SKU) (domain.StockInfo, error) {
	return domain.StockInfo{
		Count: domain.Count(5),
	}, nil
}
