package repository

import (
	"context"

	"github.com/Kroning/mytheresa/internal/domain"
)

//go:generate go run go.uber.org/mock/mockgen -source=interface.go -destination=mocks/mock.go -typed

type ProductRepo interface {
	GetProducts(ctx context.Context, category string, price int) ([]*domain.Product, error)
}
