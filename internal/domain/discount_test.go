package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDiscountsToProduct(t *testing.T) {
	var nilProductWithDiscount []*ProductWithDiscount
	tests := []struct {
		name      string
		products  []*Product
		discounts Discounts
		expect    []*ProductWithDiscount
	}{
		{
			name:      "product with category discount",
			products:  []*Product{defaultProduct()},
			discounts: getDiscounts(discountSandals(), discountBoots()),
			expect:    []*ProductWithDiscount{defaultProductWithDiscount()},
		},
		{
			name:      "product with sku discount",
			products:  []*Product{defaultProduct()},
			discounts: getDiscounts(discountSku()),
			expect:    []*ProductWithDiscount{productWithSKUDiscount()},
		},
		{
			name:      "no products",
			products:  []*Product{},
			discounts: Discounts{},
			expect:    nilProductWithDiscount,
		},
		{
			name:      "no discounts",
			products:  []*Product{defaultProduct()},
			discounts: Discounts{},
			expect:    []*ProductWithDiscount{defaultProductWithNoDiscount()},
		},
		{
			name:      "no matching discounts",
			products:  []*Product{defaultProduct()},
			discounts: getDiscounts(discountSandals()),
			expect:    []*ProductWithDiscount{defaultProductWithNoDiscount()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := AddDiscountsToProduct(tt.products, tt.discounts)

			assert.Equal(t, tt.expect, actual)
		})
	}
}

func defaultProduct() *Product {
	return &Product{
		Sku:      "00001",
		Name:     "Some name",
		Category: "boots",
		Price:    100,
		Currency: CurrencyName,
	}
}

func discountBoots() *Discount {
	return &Discount{
		TypeName:  "category",
		TypeValue: "boots",
		Amount:    20,
	}
}

func discountSandals() *Discount {
	return &Discount{
		TypeName:  "category",
		TypeValue: "sandals",
		Amount:    20,
	}
}

func discountSku() *Discount {
	return &Discount{
		TypeName:  "sku",
		TypeValue: "00001",
		Amount:    10,
	}
}

func getDiscounts(discounts ...*Discount) Discounts {
	discountMap := map[string][]*Discount{}
	for _, discount := range discounts {
		if discountMap[discount.TypeName] == nil {
			discountMap[discount.TypeName] = []*Discount{}
		}
		discountMap[discount.TypeName] = append(discountMap[discount.TypeName], discount)
	}

	return discountMap
}

func defaultProductWithDiscount() *ProductWithDiscount {
	return &ProductWithDiscount{
		Product:  *defaultProduct(),
		Discount: *discountBoots(),
	}
}

func productWithSKUDiscount() *ProductWithDiscount {
	return &ProductWithDiscount{
		Product:  *defaultProduct(),
		Discount: *discountSku(),
	}
}

func defaultProductWithNoDiscount() *ProductWithDiscount {
	return &ProductWithDiscount{
		Product:  *defaultProduct(),
		Discount: Discount{},
	}
}

func Test_maxDiscount(t *testing.T) {
	discountVar := discountBoots()
	tests := []struct {
		name        string
		curDiscount *Discount
		discount    *Discount
		expect      Discount
	}{
		{
			name:        "current discount is bigger",
			curDiscount: discountBoots(),
			discount:    discountSku(),
			expect:      *discountVar,
		},
		{
			name:        "current discount is smaller",
			curDiscount: discountSku(),
			discount:    discountBoots(),
			expect:      *discountVar,
		},
		{
			name:        "current discount is nil",
			curDiscount: nil,
			discount:    discountBoots(),
			expect:      *discountVar,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxDiscount(tt.curDiscount, tt.discount)

			assert.Equal(t, tt.expect, actual)
		})
	}
}
