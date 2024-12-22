package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kroning/mytheresa/internal/domain"
)

func TestMapProductsWithDiscountResponse(t *testing.T) {
	product1 := domain.Product{
		Sku:      "sku1",
		Name:     "name1",
		Category: "cat1",
		Price:    435,
		Currency: domain.CurrencyName,
	}
	productWithDiscount1 := &domain.ProductWithDiscount{Product: product1} // discount is 0
	product2 := domain.Product{
		Sku:      "sku2",
		Name:     "name2",
		Category: "cat2",
		Price:    500,
		Currency: domain.CurrencyName,
	}
	productWithDiscount2 := &domain.ProductWithDiscount{Product: product2, Discount: 20}
	response1 := &ProductWithDiscountResponse{
		Sku:      "sku1",
		Name:     "name1",
		Category: "cat1",
		Price: Price{
			Original: 435,
			Final:    435,
			Discount: nil,
			Currency: domain.CurrencyName,
		},
	}
	twenty := "20%"
	response2 := &ProductWithDiscountResponse{
		Sku:      "sku2",
		Name:     "name2",
		Category: "cat2",
		Price: Price{
			Original: 500,
			Final:    400,
			Discount: &twenty,
			Currency: domain.CurrencyName,
		},
	}

	tests := []struct {
		name     string
		product  []*domain.ProductWithDiscount
		response []*ProductWithDiscountResponse
	}{
		{
			name:     "product with no discount",
			product:  []*domain.ProductWithDiscount{productWithDiscount1},
			response: []*ProductWithDiscountResponse{response1},
		},
		{
			name:     "product with discount",
			product:  []*domain.ProductWithDiscount{productWithDiscount2},
			response: []*ProductWithDiscountResponse{response2},
		},
		{
			name:     "both products",
			product:  []*domain.ProductWithDiscount{productWithDiscount2, productWithDiscount1},
			response: []*ProductWithDiscountResponse{response2, response1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := MapProductsWithDiscountResponse(tt.product)

			assert.Equal(t, tt.response, resp)
		})
	}
}
