package service

import (
	"context"
	"time"

	"github.com/rmarsu/auth_service/internal/repository"
	"github.com/rmarsu/auth_service/pkg/hash"
	"github.com/rmarsu/auth_service/pkg/jwt"
)

type Services struct {
	Auth Auth
}

type Deps struct {
	Repo         *repository.Repository
	Hasher       *hash.SHA256Hasher
	TokenManager *jwt.Manager
	TTL          time.Duration
}

type Auth interface {
	RegisterUser(ctx context.Context, email, username, password string) (int64, error)
	Login(ctx context.Context, username, password string, appId int64) (string, error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

func NewServices(deps *Deps) Services {
	return Services{
		Auth: NewAuthService(deps.Repo, deps.Hasher, deps.TokenManager, deps.TTL),
	}
}
