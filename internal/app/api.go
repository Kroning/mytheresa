package app

import (
	"context"

	v1 "github.com/Kroning/mytheresa/internal/transport/http/v1"
	"github.com/go-chi/chi/v5"

	httpServ "github.com/Kroning/mytheresa/internal/transport/http"
)

func (c *Container) GetHttpRouter(ctx context.Context) chi.Router {
	if c.httpRouter == nil {
		c.httpRouter = httpServ.NewRouter(
			c.GetHTTPApiHandlerV1(ctx),
		)
	}

	return c.httpRouter
}

func (c *Container) GetHTTPApiHandlerV1(ctx context.Context) *v1.ApiHandler {
	if c.httpApiHandlerV1 == nil {
		c.httpApiHandlerV1 = v1.NewApiHandler(
			c.GetProductService(ctx),
			c.GetDiscountService(ctx),
		)
	}

	return c.httpApiHandlerV1
}
