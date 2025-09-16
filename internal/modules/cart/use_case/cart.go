package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
)

type CartUseCaseInterface interface {
	AddItem(ctx context.Context, userID domain.UserID, item domain.Item) error
	DeleteItem(ctx context.Context, userID domain.UserID, item domain.SKU) error
	List(ctx context.Context, userID domain.UserID) (domain.Cart, error)
	Clear(ctx context.Context, userID domain.UserID) error
	Checkout(ctx context.Context, userID domain.UserID) error
}

type CartUseCase struct {
	logger         ports.Logger
	cartRepository ports.CartRepository
	productService ports.ProductService
	lomsService    ports.LomsService
}

func NewCartUseCase(
	logger ports.Logger,
	cartRepository ports.CartRepository,
	productService ports.ProductService,
	lomsService ports.LomsService,
) (CartUseCase, error) {
	return CartUseCase{
		logger:         logger,
		cartRepository: cartRepository,
		productService: productService,
		lomsService:    lomsService,
	}, nil
}

func (c *CartUseCase) AddItem(ctx context.Context, userID domain.UserID, item domain.Item) error {
	_, err := c.productService.GetProduct(ctx, item.Sku)
	if err != nil {
		return err
	}

	stockInfo, err := c.lomsService.StockInfo(ctx, item.Sku)
	if err != nil {
		return err
	}

	if stockInfo.Count < item.Count {
		return domain.ErrNotEnoughStock
	}

	err = c.cartRepository.AddItem(ctx, userID, item.Sku, item.Count)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartUseCase) DeleteItem(ctx context.Context, userID domain.UserID, item domain.SKU) error {
	return nil
}

func (c *CartUseCase) List(ctx context.Context, userID domain.UserID) (domain.Cart, error) {
	return domain.Cart{}, nil
}

func (c *CartUseCase) Clear(ctx context.Context, userID domain.UserID) error {
	return nil
}

func (c *CartUseCase) Checkout(ctx context.Context, userID domain.UserID) error {
	return nil
}
