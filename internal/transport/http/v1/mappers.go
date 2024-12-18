package v1

import (
	"fmt"

	"github.com/Kroning/mytheresa/internal/domain"
)

func MapProductsWithDiscountResponse(products []*domain.ProductWithDiscount) []*ProductWithDiscountResponse {
	resp := make([]*ProductWithDiscountResponse, len(products))
	for i, product := range products {
		var discount *string
		finalPrice := product.Price
		if product.Discount > 0 {
			finalPrice = int(float64(product.Price) * (1 - float64(product.Discount)/100))
			discountStr := fmt.Sprintf("%d%%", product.Discount)
			discount = &discountStr
		}
		p := &ProductWithDiscountResponse{
			Sku:      product.Sku,
			Name:     product.Name,
			Category: product.Category,
			Price: Price{
				Original: product.Price,
				Final:    finalPrice,
				Discount: discount,
				Currency: product.Currency,
			},
		}
		resp[i] = p
	}

	return resp
}
