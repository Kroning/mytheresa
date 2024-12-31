package v1

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Kroning/mytheresa/internal/domain"
	"github.com/Kroning/mytheresa/internal/logger"
	mock_repository "github.com/Kroning/mytheresa/internal/repository/mocks"
	"github.com/Kroning/mytheresa/internal/service/discount"

	"github.com/Kroning/mytheresa/internal/service/product"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type mocks struct {
	productService  *product.ProductService
	discountService *discount.DiscountService
	productRepo     *mock_repository.MockProductRepo
}

func TestApiHandler_GetProducts(t *testing.T) {
	ctx := context.Background()
	product1 := &domain.Product{
		Sku:      "00001",
		Name:     "some boots",
		Category: "boots",
		Price:    9000,
		Currency: "EUR",
	}
	product2 := &domain.Product{
		Sku:      "00002",
		Name:     "some boots too",
		Category: "boots",
		Price:    10000,
		Currency: "EUR",
	}

	tests := []struct {
		name     string
		url      string
		mocks    func(mocks)
		expected string
		httpCode int
		err      string
	}{
		{
			name: "success",
			url:  "/api/v1/products?category=boots&priceLessThan=90000",
			mocks: func(m mocks) {
				gomock.InOrder(
					m.productRepo.EXPECT().GetProducts(ctx, "boots", 90000).Return(
						[]*domain.Product{product1, product2}, nil,
					).Times(1),
				)
			},
			expected: `[{"sku":"00001","name":"some boots","category":"boots","price":{"original":9000,"final":6300,"discount_percentage":"30%","currency":"EUR"}},{"sku":"00002","name":"some boots too","category":"boots","price":{"original":10000,"final":7000,"discount_percentage":"30%","currency":"EUR"}}]`,
			httpCode: http.StatusOK,
			err:      "",
		},
		{
			name: "zero products",
			url:  "/api/v1/products?category=sombreros&priceLessThan=90",
			mocks: func(m mocks) {
				gomock.InOrder(
					m.productRepo.EXPECT().GetProducts(ctx, "sombreros", 90).Return(
						[]*domain.Product{}, nil,
					).Times(1),
				)
			},
			expected: `[]`,
			httpCode: http.StatusOK,
			err:      "",
		},
		{
			name:     "no category",
			url:      "/api/v1/products?priceLessThan=90",
			mocks:    nil,
			expected: `[]`,
			httpCode: http.StatusBadRequest,
			err:      ErrNoCategory,
		},
		{
			name:     "invalid price",
			url:      "/api/v1/products?category=sombreros&priceLessThan=a90",
			mocks:    nil,
			expected: `[]`,
			httpCode: http.StatusBadRequest,
			err:      ErrInvalidPrice,
		},
		{
			name: "GetProducts error",
			url:  "/api/v1/products?category=boots&priceLessThan=90000",
			mocks: func(m mocks) {
				gomock.InOrder(
					m.productRepo.EXPECT().GetProducts(ctx, "boots", 90000).Return(
						[]*domain.Product{}, errors.New("some error"),
					).Times(1),
				)
			},
			expected: `[]`,
			httpCode: http.StatusInternalServerError,
			err:      ErrGetProducts,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			productRepo := mock_repository.NewMockProductRepo(ctrl)
			productService, _ := product.NewService(productRepo, logger.Logger())
			discountService, _ := discount.NewService(logger.Logger())
			//discountServiceMock := mock_service.NewMockDiscountService(ctrl)

			m := mocks{
				productService:  productService,
				discountService: discountService,
				productRepo:     productRepo,
			}
			if tt.mocks != nil {
				tt.mocks(m)
			}

			apiHandler := NewApiHandler(
				productService,
				discountService,
			)

			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			w := httptest.NewRecorder()
			handler := http.HandlerFunc(apiHandler.GetProducts)
			handler.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tt.httpCode)
			if w.Code != http.StatusOK {
				assert.Contains(t, w.Body.String(), tt.err)
				return
			}

			if strings.TrimSpace(w.Body.String()) != tt.expected {
				fmt.Println(w.Body)
				t.Errorf("unexpected body: got %v want %v", w.Body.String(), tt.expected)
			}
		})
	}
}
