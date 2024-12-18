package product

import (
	"context"
	"fmt"

	"github.com/Kroning/mytheresa/internal/domain"
	"go.uber.org/zap"
)

type productRepo interface {
	GetProducts(ctx context.Context, category string, price int) ([]*domain.Product, error)
}

type Service struct {
	productRepo productRepo
	logger      *zap.Logger
}

func NewService(
	productRepo productRepo,
	logger *zap.Logger,
) (*Service, error) {
	if productRepo == nil {
		return nil, fmt.Errorf("product repo is nil")
	}
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	return &Service{
		productRepo: productRepo,
		logger:      logger,
	}, nil
}

func (s *Service) GetProductsWithDiscount(ctx context.Context, category string, price int) ([]*domain.ProductWithDiscount, error) {
	products, err := s.GetProducts(ctx, category, price)
	if err != nil {
		return nil, fmt.Errorf("cant get products from DB: %w", err)
	}
	
	// TODO: separate service for discounts
	productsWithDiscount := domain.AddDiscounts(products)

	return productsWithDiscount, nil
}

func (s *Service) GetProducts(ctx context.Context, category string, price int) ([]*domain.Product, error) {

	return s.productRepo.GetProducts(ctx, category, price)
}
