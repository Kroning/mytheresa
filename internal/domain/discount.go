package domain

type Discount struct {
	TypeName  string
	TypeValue string
	Amount    int
}

type Discounts map[string][]*Discount

/*func AddDiscounts(products []*Product) []*ProductWithDiscount {
	// TODO: discounts should be a separate object and should be adjustable
	var productsWithDiscount []*ProductWithDiscount
	for _, product := range products {
		productWithDiscount := &ProductWithDiscount{Product: *product}
		if product.Category == "boots" {
			productWithDiscount.DiscountAmount = 30
		} else if product.Sku == "000003" {
			productWithDiscount.DiscountAmount = 15
		}
		productsWithDiscount = append(productsWithDiscount, productWithDiscount)
	}

	return productsWithDiscount
}*/

//type DiscountAmount int

type ProductWithDiscount struct {
	Product
	Discount
}

func AddDiscountsToProduct(products []*Product, discounts Discounts) []*ProductWithDiscount {
	// TODO: discounts should be a separate object and should be adjustable
	var productsWithDiscount []*ProductWithDiscount
	for _, product := range products {
		productWithDiscount := &ProductWithDiscount{Product: *product}
		var curDiscount *Discount

		for _, discount := range discounts["category"] {
			if product.Category == discount.TypeValue {
				productWithDiscount.Discount = maxDiscount(curDiscount, discount)
			}
		}

		for _, discount := range discounts["sku"] {
			if product.Sku == discount.TypeValue {
				productWithDiscount.Discount = maxDiscount(curDiscount, discount)
			}
		}

		productsWithDiscount = append(productsWithDiscount, productWithDiscount)
	}

	return productsWithDiscount
}

func maxDiscount(curDiscount, discount *Discount) Discount {
	if curDiscount != nil && curDiscount.Amount > discount.Amount {
		return *curDiscount
	}

	return *discount
}
