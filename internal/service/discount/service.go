package discount

import (
	"context"
	"fmt"

	"github.com/Kroning/mytheresa/internal/domain"
	"go.uber.org/zap"
)

type discountRepo interface {
	GetProducts(ctx context.Context, category string, price int) ([]*domain.Product, error)
}

type DiscountService struct {
	//productRepo productRepo
	logger *zap.Logger
}

func NewService(
	//productRepo productRepo,
	logger *zap.Logger,
) (*DiscountService, error) {
	/*if productRepo == nil {
		return nil, fmt.Errorf("product repo is nil")
	}*/
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	return &DiscountService{
		//productRepo: productRepo,
		logger: logger,
	}, nil
}

var defaultDiscounts = domain.Discounts{
	"category": {
		{
			TypeName:  "category",
			TypeValue: "boots",
			Amount:    30,
		},
	},
	"sky": {
		{
			TypeName:  "sku",
			TypeValue: "000003",
			Amount:    15,
		},
	},
}

func (s *DiscountService) GetDiscounts(ctx context.Context) (domain.Discounts, error) {
	return defaultDiscounts, nil
}
