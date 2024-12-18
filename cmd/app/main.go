package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Kroning/mytheresa/internal/app"
	"github.com/Kroning/mytheresa/internal/config"
	"github.com/Kroning/mytheresa/internal/logger"
	httpServ "github.com/Kroning/mytheresa/internal/transport/http"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	defer func() {
		_ = logger.Logger().Sync()
	}()

	if err := run(); err != nil {
		logger.Fatal(context.Background(), "incentives web server start / shutdown problem", zap.Error(err))
	}
}

func run() error {
	logger.SetLevel("debug")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// TODO: enrich logger with ctx data
	cfg, err := config.NewConfig(logger.Logger())
	if err != nil {
		logger.Fatal(ctx, "cant get config", zap.Error(err))
	}

	logger.SetLevel(cfg.App.LogLevel)

	container := app.New(cfg, logger.Logger())
	defer container.Close()

	// TODO metrics, tracer, debugserver

	g, _ := errgroup.WithContext(ctx)

	httpRouter := container.GetHttpRouter(ctx)
	httpServer, err := httpServ.NewServer(cfg.Server.HTTP, httpRouter)
	if err != nil {
		return fmt.Errorf("cant create http server %w", err)
	}

	g.Go(func() error {
		go func() {
			if err := httpServer.Start(); err != nil {
				logger.Error(ctx, "error starting http server", zap.Error(err))
				cancel()
			}
		}()

		logger.Info(ctx, "http server started", zap.Int("port", cfg.Server.HTTP.Port))

		return nil
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("cant start service %w", err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer signal.Stop(interrupt)

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		logger.Info(ctx, "Context Done")
		break
	}

	logger.Info(ctx, "shutting down...")
	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err = httpServer.Stop(shutdownCtx); err != nil {
		logger.Error(ctx, "problem while shutting down http server", zap.Error(err))
	}

	return nil
}
