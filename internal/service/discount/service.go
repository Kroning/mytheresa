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

type Service struct {
	//productRepo productRepo
	logger *zap.Logger
}

func NewService(
	//productRepo productRepo,
	logger *zap.Logger,
) (*Service, error) {
	/*if productRepo == nil {
		return nil, fmt.Errorf("product repo is nil")
	}*/
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	return &Service{
		//productRepo: productRepo,
		logger: logger,
	}, nil
}

func (s *Service) GetDiscounts(ctx context.Context) (domain.Discounts, error) {
	return domain.Discounts{
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
	}, nil
}
