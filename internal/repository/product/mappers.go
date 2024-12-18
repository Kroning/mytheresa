package product

import "github.com/Kroning/mytheresa/internal/domain"

func mapProductRowToProduct(row ProductRow) *domain.Product {
	return &domain.Product{
		Sku:      row.Sku,
		Name:     row.Name,
		Category: row.Category,
		Price:    row.Price,
		Currency: domain.CurrencyName,
	}
}
