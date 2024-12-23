package app

import (
	"context"

	"github.com/Kroning/mytheresa/internal/service/discount"
	"github.com/Kroning/mytheresa/internal/service/product"
)

func (c *Container) GetProductService(ctx context.Context) *product.Service {
	if c.productService == nil {
		productService, err := product.NewService(
			c.GetProductRepo(ctx),
			c.Logger(),
		)
		if err != nil {
			panic(err)
		}

		c.productService = productService
	}

	return c.productService
}

func (c *Container) GetDiscountService(ctx context.Context) *discount.Service {
	if c.discountService == nil {
		discountService, err := discount.NewService(
			//c.GetProductRepo(ctx),
			c.Logger(),
		)
		if err != nil {
			panic(err)
		}

		c.discountService = discountService
	}

	return c.discountService
}
