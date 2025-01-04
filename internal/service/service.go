package service

import (
	"context"

	"github.com/rmarsu/auth_service/internal/repository"
	"github.com/rmarsu/auth_service/pkg/hash"
)

type Services struct {
	Auth Auth
}

type Deps struct {
	Repo   *repository.Repository
	Hasher *hash.SHA256Hasher
}

type Auth interface {
	RegisterUser(ctx context.Context, username, password string) (int64, error)
	Login(ctx context.Context, username, password string, appId int64) (string, error)
	IsAdmin(ctx context.Context, token string) (bool, error)
}

func NewServices(deps *Deps) Services {
	return Services{
		Auth: NewAuthService(deps.Repo, deps.Hasher),
	}
}
