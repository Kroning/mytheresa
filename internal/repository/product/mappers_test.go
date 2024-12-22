package product

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kroning/mytheresa/internal/domain"
)

func Test_mapProductRowToProduct(t *testing.T) {

	tests := []struct {
		name     string
		row      ProductRow
		expected *domain.Product
	}{
		{
			name:     "success",
			row:      ProductRow{"a", "b", "d", 100},
			expected: &domain.Product{"a", "b", "d", 100, domain.CurrencyName},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := mapProductRowToProduct(tt.row)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
