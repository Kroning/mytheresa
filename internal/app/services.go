package app

import (
	"context"

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
