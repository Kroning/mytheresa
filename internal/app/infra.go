package app

import (
	"context"

	"github.com/Kroning/mytheresa/internal/database/postgresql"
)

func (c *Container) GetDb(ctx context.Context) *postgresql.Storage {
	if c.db == nil {
		cfg := c.Config()
		db, err := postgresql.New(postgresql.Config{
			Master: postgresql.NodeConfig{
				Host:           cfg.DB.Master.Host,
				Port:           cfg.DB.Master.Port,
				User:           cfg.DB.Master.User,
				Password:       cfg.DB.Master.Password,
				Database:       cfg.DB.Master.Database,
				MaxOpen:        cfg.DB.Master.MaxOpen,
				Timeout:        cfg.DB.Master.Timeout,
				MigrationsPath: cfg.DB.Master.MigrationsPath,
			},
		})
		if err != nil {
			panic(err)
		}

		c.db = db
	}

	return c.db
}
