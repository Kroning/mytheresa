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

type ProductService struct {
	productRepo productRepo
	logger      *zap.Logger
}

func NewService(
	productRepo productRepo,
	logger *zap.Logger,
) (*ProductService, error) {
	if productRepo == nil {
		return nil, fmt.Errorf("product repo is nil")
	}
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	return &ProductService{
		productRepo: productRepo,
		logger:      logger,
	}, nil
}

func (s *ProductService) GetProducts(ctx context.Context, category string, price int) ([]*domain.Product, error) {

	return s.productRepo.GetProducts(ctx, category, price)
}
