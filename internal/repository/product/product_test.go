package product

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/Kroning/mytheresa/internal/database/postgresql"
	"github.com/Kroning/mytheresa/internal/domain"
)

func TestProductRepo_GetProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewRepo(postgresql.NewFromDB(sqlxDB))
	rows := []domain.Product{
		{"000001", "BV Lean leather ankle boots", "boots", 89000, "EUR"},
		{"000002", "BV Lean leather ankle boots", "boots", 99000, "EUR"},
		{"000004", "Naima embellished suede sandals", "sandals", 79500, "EUR"},
	}

	tests := []struct {
		name       string
		category   string
		price      int
		returnRows []domain.Product
		sql        string
		want       []*domain.Product
		wantErr    string
	}{
		{
			name:       "success",
			category:   "boots",
			price:      100000,
			returnRows: []domain.Product{rows[0], rows[1]},
			sql:        `SELECT sku, name, category, price FROM products WHERE category = \$1 AND price <= \$2 LIMIT 5`,
			want:       []*domain.Product{&rows[0], &rows[1]},
			wantErr:    "",
		},
		{
			name:       "no price",
			category:   "boots",
			price:      0,
			returnRows: []domain.Product{rows[0], rows[1]},
			sql:        `SELECT sku, name, category, price FROM products WHERE category = \$1 LIMIT 5`,
			want:       []*domain.Product{&rows[0], &rows[1]},
			wantErr:    "",
		},
		{
			name:       "sql error",
			category:   "boot",
			price:      100,
			returnRows: []domain.Product{rows[0], rows[1]},
			sql:        `SELECT sku, name, category, price FROM products WHERE category = \$1 AND price <= \$2 LIMIT 5`,
			want:       []*domain.Product{&rows[0], &rows[1]},
			wantErr:    "some error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRows := sqlmock.NewRows([]string{"sku", "name", "category", "price"})
			for _, row := range tt.returnRows {
				mockRows.AddRow(row.Sku, row.Name, row.Category, row.Price)
			}
			args := []driver.Value{tt.category}
			if tt.price > 0 {
				args = append(args, tt.price)
			}

			if tt.wantErr != "" {
				mock.ExpectQuery(tt.sql).WithArgs(args...).WillReturnError(errors.New(tt.wantErr))
			} else {
				mock.ExpectQuery(tt.sql).WithArgs(args...).WillReturnRows(mockRows)
			}

			products, err := repo.GetProducts(context.Background(), tt.category, tt.price)

			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, products)
		})
	}
}
