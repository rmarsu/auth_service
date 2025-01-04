package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rmarsu/auth_service/internal/domain"
)

type Repository struct {
	Auth Auth
}

type Auth interface {
	CreateUser(ctx context.Context, email string, username string, passHash []byte) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetAppById(ctx context.Context, id int64) (domain.App, error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
	}
}
