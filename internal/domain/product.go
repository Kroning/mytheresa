package domain

const CurrencyName = "EUR"

type Product struct {
	Sku      string
	Name     string
	Category string
	Price    int
	Currency string
}
