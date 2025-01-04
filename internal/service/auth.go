package service

import (
	"context"

	"github.com/rmarsu/auth_service/internal/repository"
	"github.com/rmarsu/auth_service/pkg/hash"
)

type AuthService struct {
	repo   *repository.Repository
	hasher *hash.SHA256Hasher
}

func NewAuthService(repo *repository.Repository, hasher *hash.SHA256Hasher) *AuthService {
	return &AuthService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *AuthService) RegisterUser(ctx context.Context, username, password string) (int64, error) {
	panic("not implemented")
}

func (s *AuthService) Login(ctx context.Context, username, password string, appId int64) (string, error) {
	panic("not implemented")
}

func (s *AuthService) IsAdmin(ctx context.Context, token string) (bool, error) {
	panic("not implemented")
}
