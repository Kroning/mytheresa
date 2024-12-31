package service

import (
	"context"

	"github.com/Kroning/mytheresa/internal/domain"
)

//go:generate go run go.uber.org/mock/mockgen -source=deps.go -destination=mocks/mock.go

type ProductService interface {
	GetProducts(ctx context.Context, category string, price int) ([]*domain.Product, error)
}

type DiscountService interface {
	GetDiscounts(ctx context.Context) (domain.Discounts, error)
}
