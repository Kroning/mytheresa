package product

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/Kroning/mytheresa/internal/domain"
)

func (r *ProductRepo) GetProducts(ctx context.Context, category string, price int) ([]*domain.Product, error) {
	var productRows []ProductRow
	query := sq.Select(
		"sku",
		"name",
		"category",
		"price ",
	).
		From("products").
		Where(sq.Eq{"category": category}).
		Limit(5).
		PlaceholderFormat(sq.Dollar)
	if price > 0 {
		query = query.Where(sq.LtOrEq{"price": price})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	err = r.db.Master.SelectContext(ctx, &productRows, sql, args...)
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
