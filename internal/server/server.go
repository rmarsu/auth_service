package server

import (
	"fmt"
	"net"
	"os"

	"github.com/rmarsu/auth_service/internal/config"
	auth_service "github.com/rmarsu/auth_service/internal/proto"
	"github.com/rmarsu/auth_service/pkg/logger"
	"google.golang.org/grpc"
)

type Server struct {
	srv      *grpc.Server
	cfg      *config.Config
	handlers []auth_service.AuthServiceServer
}

func New(cfg *config.Config, handlers ...auth_service.AuthServiceServer) *Server {
	return &Server{
		cfg:      cfg,
		handlers: handlers,
	}
}

func (s *Server) Serve() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		logger.Errorf("failed to listen: %v", err)
	}
	defer l.Close()
	s.srv = grpc.NewServer()

	for _, handler := range s.handlers {
		auth_service.RegisterAuthServiceServer(s.srv, handler)
	}

	if err := s.srv.Serve(l); err != nil {
		logger.Errorf("failed to serve: %v", err)
	}
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}
