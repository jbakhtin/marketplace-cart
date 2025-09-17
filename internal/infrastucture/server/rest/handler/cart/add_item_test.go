package cart

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/mock/cart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/logger/noop"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/suite"
)

type SuiteAddItem struct {
	suite.Suite
	handler     Handler
	mockUseCase *cart.CartUseCaseInterface
}

func (s *SuiteAddItem) SetupSuite() {
	s.mockUseCase = new(cart.CartUseCaseInterface)

	s.handler, _ = NewHandler(noop.Logger{}, s.mockUseCase)
}

func (s *SuiteAddItem) TearDownSuite() {
	s.mockUseCase.AssertExpectations(s.T())
}

func (s *SuiteAddItem) TestAddItem_CheckValidation() {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	type useCaseArgs struct {
		userID domain.UserID
		item   domain.Item
	}

	type useCaseHits struct {
		err error
	}

	type useCase struct {
		needCall bool
		function string
		args     useCaseArgs
		hits     useCaseHits
	}

	for _, testCase := range []struct {
		name           string
		args           args
		useCase        useCase
		expectedStatus int
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					"POST",
					"/cart/items/add",
					strings.NewReader(`{"sku":2,"count":1}`),
				),
			},
			useCase: useCase{
				needCall: true,
				function: "AddItem",
				args: useCaseArgs{
					userID: domain.UserID(1),
					item: domain.Item{
						Sku:   domain.SKU(2),
						Count: domain.Count(1),
					},
				},
				hits: useCaseHits{
					err: nil,
				},
			},
			expectedStatus: http.StatusCreated,
		},
	} {
		s.T().Run(testCase.name, func(t *testing.T) {
			if testCase.useCase.needCall {
				s.mockUseCase.
					On(testCase.useCase.function, mock.Anything, testCase.useCase.args.userID, testCase.useCase.args.item).
					Return(testCase.useCase.hits.err).
					Once()
			}

			ctx := context.WithValue(testCase.args.r.Context(), "user_id", testCase.useCase.args.userID)
			testCase.args.r = testCase.args.r.WithContext(ctx)

			s.handler.AddItem(testCase.args.w, testCase.args.r)
			s.Equal(testCase.expectedStatus, testCase.args.w.(*httptest.ResponseRecorder).Code)
		})
	}
}

func TestAddItemSuite(t *testing.T) {
	suite.Run(t, new(SuiteAddItem))
}
