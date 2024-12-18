package domain

const CurrencyName = "EUR"

type Product struct {
	Sku      string
	Name     string
	Category string
	Price    int
	Currency string
}

type Discount int

type ProductWithDiscount struct {
	Product
	Discount
}

func AddDiscounts(products []*Product) []*ProductWithDiscount {
	// TODO: discounts should be a separate object and should be adjustable
	var productsWithDiscount []*ProductWithDiscount
	for _, product := range products {
		productWithDiscount := &ProductWithDiscount{Product: *product}
		if product.Category == "boots" {
			productWithDiscount.Discount = 30
		} else if product.Sku == "000003" {
			productWithDiscount.Discount = 15
		}
		productsWithDiscount = append(productsWithDiscount, productWithDiscount)
	}

	return productsWithDiscount
}
