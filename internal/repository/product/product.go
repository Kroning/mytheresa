package product

import (
	"context"
	"fmt"

	"github.com/Kroning/mytheresa/internal/domain"
)

func (r *Repo) GetProducts(ctx context.Context, category string, price int) ([]*domain.Product, error) {
	var productRows []ProductRow
	query := "SELECT " +
		"sku," +
		"name," +
		"category," +
		"price " +
		"FROM products " +
		"WHERE category = $1 "
	if price > 0 {
		query = query + "AND price <= $2 "
	}
	query += "LIMIT 5;"

	var err error
	// TODO: use squirrel to construct SQL and args
	if price > 0 {
		err = r.db.Master.SelectContext(ctx, &productRows, query, category, price)
	} else {
		err = r.db.Master.SelectContext(ctx, &productRows, query, category)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to make select: %w", err)
	}

	products := make([]*domain.Product, 0, len(productRows))
	for _, row := range productRows {
		product := mapProductRowToProduct(row)
		products = append(products, product)
	}

	return products, nil
}
