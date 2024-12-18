package app

import (
	"context"

	"github.com/Kroning/mytheresa/internal/repository/product"
)

func (c *Container) GetProductRepo(ctx context.Context) *product.Repo {
	if c.productRepo == nil {
		c.productRepo = product.NewRepo(
			c.GetDb(ctx),
		)
	}

	return c.productRepo
}
