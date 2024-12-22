package app

import (
	"github.com/Kroning/mytheresa/internal/database/postgresql"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/Kroning/mytheresa/internal/config"
	productrepo "github.com/Kroning/mytheresa/internal/repository/product"
	"github.com/Kroning/mytheresa/internal/service/product"
	v1 "github.com/Kroning/mytheresa/internal/transport/http/v1"
)

type Container struct {
	config *config.Config
	logger *zap.Logger

	// services
	productService *product.Service

	// infrastructure
	db *postgresql.Storage

	// repositories
	productRepo *productrepo.ProductRepo

	// api (http, grpc, graphql etc.)
	httpRouter       chi.Router
	httpApiHandlerV1 *v1.ApiHandler
}

func New(conf *config.Config, logger *zap.Logger) *Container {
	return &Container{
		config: conf,
		logger: logger,
	}
}

func (c *Container) AppName() string {
	return c.Config().App.Name
}

func (c *Container) Config() *config.Config {
	return c.config
}

func (c *Container) Logger() *zap.Logger {
	return c.logger
}

func (c *Container) Close() {
	// add closer for each service you need to close
	c.db.Close()
	c.Logger().Debug("container close")
}
