package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/logger/noop"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/mocks"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SuiteUseCaseAddItem struct {
	suite.Suite

	useCase            CartUseCase
	mockCartRepository *mocks.CartRepository
	mockProductService *mocks.ProductService
	mockLomsService    *mocks.LomsService
}

func (s *SuiteUseCaseAddItem) SetupSuite() {
	s.mockCartRepository = mocks.NewCartRepository(s.T())
	s.mockProductService = mocks.NewProductService(s.T())
	s.mockLomsService = mocks.NewLomsService(s.T())

	s.useCase, _ = NewCartUseCase(noop.Logger{}, s.mockCartRepository, s.mockProductService, s.mockLomsService)
}

func (s *SuiteUseCaseAddItem) TearDownSuite() {
	s.mockCartRepository.AssertExpectations(s.T())
	s.mockProductService.AssertExpectations(s.T())
	s.mockLomsService.AssertExpectations(s.T())
}

func (s *SuiteUseCaseAddItem) TestAddItem() {
	type args struct {
		userID domain.UserID
		item   domain.Item
	}

	type productServiceArgs struct {
		SKU domain.SKU
	}

	type productServiceHits struct {
		productInfo domain.ProductInfo
		err         error
	}

	type productService struct {
		needCall bool
		function string
		args     productServiceArgs
		hits     productServiceHits
	}

	type lomsServiceArgs struct {
		sku domain.SKU
	}

	type lomsServiceHits struct {
		stockInfo domain.StockInfo
		err       error
	}

	type lomsService struct {
		needCall bool
		function string
		args     lomsServiceArgs
		hits     lomsServiceHits
	}

	type cartRepositoryArgs struct {
		userID domain.UserID
		sku    domain.SKU
		count  domain.Count
	}

	type cartRepositoryHits struct {
		err error
	}

	type cartRepository struct {
		needCall bool
		function string
		args     cartRepositoryArgs
		hits     cartRepositoryHits
	}

	for _, testCase := range []struct {
		name           string
		args           args
		cartRepository cartRepository
		productService productService
		lomsService    lomsService
		expectedError  error
	}{
		{
			name: "success",
			args: args{
				userID: domain.UserID(1),
				item: domain.Item{
					Sku:   domain.SKU(1),
					Count: domain.Count(1),
				},
			},
			productService: productService{
				needCall: true,
				function: "GetProduct",
				args: productServiceArgs{
					SKU: domain.SKU(1),
				},
				hits: productServiceHits{
					productInfo: domain.ProductInfo{
						Name:  domain.Name("First Product"),
						Price: domain.Price(100),
					},
					err: nil,
				},
			},
			lomsService: lomsService{
				needCall: true,
				function: "StockInfo",
				args: lomsServiceArgs{
					sku: domain.SKU(1),
				},
				hits: lomsServiceHits{
					stockInfo: domain.StockInfo{
						Count: domain.Count(10),
					},
					err: nil,
				},
			},
			cartRepository: cartRepository{
				needCall: true,
				function: "AddItem",
				args: cartRepositoryArgs{
					userID: domain.UserID(1),
					sku:    domain.SKU(1),
					count:  domain.Count(1),
				},
				hits: cartRepositoryHits{
					err: nil,
				},
			},
			expectedError: nil,
		},
		{
			name: "product not found",
			args: args{
				userID: domain.UserID(1),
				item: domain.Item{
					Sku:   domain.SKU(1),
					Count: domain.Count(1),
				},
			},
			productService: productService{
				needCall: true,
				function: "GetProduct",
				args: productServiceArgs{
					SKU: domain.SKU(1),
				},
				hits: productServiceHits{
					productInfo: domain.ProductInfo{},
					err:         ports.ErrProductNotFound,
				},
			},
			lomsService: lomsService{
				needCall: false,
			},
			cartRepository: cartRepository{
				needCall: false,
			},
			expectedError: ports.ErrProductNotFound,
		},
	} {
		s.T().Run(testCase.name, func(t *testing.T) {
			if testCase.cartRepository.needCall {
				s.mockCartRepository.
					On(testCase.cartRepository.function,
						mock.Anything,
						testCase.cartRepository.args.userID,
						testCase.cartRepository.args.sku,
						testCase.cartRepository.args.count).
					Return(testCase.cartRepository.hits.err).
					Once()
			}

			if testCase.productService.needCall {
				s.mockProductService.
					On(testCase.productService.function,
						mock.Anything,
						testCase.productService.args.SKU).
					Return(testCase.productService.hits.productInfo, testCase.productService.hits.err).
					Once()
			}

			if testCase.lomsService.needCall {
				s.mockLomsService.
					On(testCase.lomsService.function,
						mock.Anything,
						testCase.lomsService.args.sku).
					Return(testCase.lomsService.hits.stockInfo, testCase.lomsService.hits.err).
					Once()
			}

			err := s.useCase.AddItem(context.TODO(), testCase.args.userID, testCase.args.item)
			s.Equal(err, testCase.expectedError)
		})
	}
}

func TestAddItemSuite(t *testing.T) {
	suite.Run(t, new(SuiteUseCaseAddItem))
}
