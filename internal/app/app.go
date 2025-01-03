package app

import (
	"github.com/rmarsu/auth_service/internal/config"
	"github.com/rmarsu/auth_service/internal/server"
)

const (
	cfgPath = "configs/config.yaml"
)

func Run() {
	cfg := config.MustLoad(cfgPath)
	
	srv := server.New(cfg)
	
	if err := srv.Run(); err != nil {
          panic(err)
     }
}
