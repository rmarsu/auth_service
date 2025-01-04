package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rmarsu/auth_service/internal/config"
	delivery_grpc "github.com/rmarsu/auth_service/internal/delivery/grpc"
	"github.com/rmarsu/auth_service/internal/repository"
	"github.com/rmarsu/auth_service/internal/server"
	"github.com/rmarsu/auth_service/internal/service"
	"github.com/rmarsu/auth_service/pkg/hash"
	"github.com/rmarsu/auth_service/pkg/logger"
)

const (
	cfgPath = "configs/config.yaml"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad(cfgPath)

	repo := repository.NewRepository()

	services := service.NewServices(&service.Deps{
		Repo:   repo,
		Hasher: hash.NewSHA256Hasher(os.Getenv("salt")),
	})

	handlers := delivery_grpc.NewAuthHandlers(services)

	srv := server.New(cfg, handlers)

	go func() {
		srv.Serve()
	}()

	logger.Infof("GRPC Server started on port %d", cfg.GRPC.Port)

	// Gracefully shutdown.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		logger.Errorf("ctx.Done: %v", done)
	}
	srv.Stop()
	logger.Info("Server stopped gracefully")
}
