package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/rmarsu/auth_service/internal/config"
	"github.com/rmarsu/auth_service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Run() error {
	ctx := context.Background()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GRPC.Port))
	if err != nil {
		logger.Errorf("failed to listen: %v", err)
	}
	defer l.Close()
	srv := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionAge:  60 * 60 * 24,
		MaxConnectionIdle: time.Second * 60,
		Time:              time.Second * 20,
		Timeout:           time.Second * 10,
	}),
		grpc.ChainUnaryInterceptor(
			grpcrecovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		),
	)
	go func() {
		logger.Infof("gRPC server listening on %s", l.Addr())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		logger.Errorf("ctx.Done: %v", done)
	}
	srv.GracefulStop()
	logger.Info("Server stopped gracefully")

	return nil
}
