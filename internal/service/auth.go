package service

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/rmarsu/auth_service/internal/domain"
	"github.com/rmarsu/auth_service/internal/repository"
	"github.com/rmarsu/auth_service/pkg/hash"
	"github.com/rmarsu/auth_service/pkg/jwt"
	"github.com/rmarsu/auth_service/pkg/logger"
)

type AuthService struct {
	repo   *repository.Repository
	hasher *hash.SHA256Hasher
	tokMgr *jwt.Manager
	ttl    time.Duration
}

func NewAuthService(repo *repository.Repository,
	hasher *hash.SHA256Hasher, tokenManager *jwt.Manager, ttl time.Duration) *AuthService {
	return &AuthService{
		repo:   repo,
		hasher: hasher,
		tokMgr: tokenManager,
		ttl:    ttl,
	}
}

func (s *AuthService) RegisterUser(ctx context.Context, email, username, password string) (int64, error) {
	logger.Info("Attemping to register user...")
	if !checkPasswordValid(password) {
		return 0, errors.New(domain.ErrPasswordIsNotValid)
	}
	hashedPassword, err := s.hasher.Hash(password)
	if err != nil {
		logger.Error("Failed to hash password")
		return 0, err
	}
	return s.repo.Auth.CreateUser(ctx, email, username, hashedPassword)

}

func (s *AuthService) Login(ctx context.Context, email, password string, appId int64) (string, error) {
	logger.Info("Attemping to login user...")
	user, err := s.repo.Auth.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			logger.Warnf("User %s not found", email)
			return "", errors.New(domain.ErrUserNotFound)
		}
		logger.Error("Failed to get user by email")
		return "", errors.New(domain.ErrSomethingWentWrong)
	}
	if !s.hasher.Verify(user.Password, password) {
		logger.Warnf("Invalid password for user %s", email)
		return "", errors.New(domain.ErrWrongPassword)
	}

	app, err := s.repo.Auth.GetAppById(ctx, appId)
	if err != nil {
		if errors.Is(err, repository.ErrAppNotFound) {
			logger.Warnf("Failed to get app by id %d", appId)
			return "", errors.New(domain.ErrAppNotFound)
		}
		logger.Errorf("Failed to get app by id %d", appId)
	}

	logger.Info("User logged in successfully")

	token, err := s.tokMgr.NewJWT(user.Id, app.Id, s.ttl)
	if err != nil {
		logger.Error("Failed to generate JWT")
		return "", err
	}
	return token, nil
}

func (s *AuthService) IsAdmin(ctx context.Context, userId int64) (bool, error) {
	logger.Infof("Checking user with ID %d for admin privileges", userId)
	isAdmin, err := s.repo.Auth.IsAdmin(ctx, userId)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			logger.Warnf("User %d not found", userId)
			return false, errors.New(domain.ErrUserNotFound)
		}
		logger.Error("Failed to check admin privileges")
		return false, err
	}
	logger.Infof("User with ID %d admin = %t", userId, isAdmin)
	return isAdmin, nil
}

func checkPasswordValid(password string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+-=]+$`)
	return re.MatchString(password)
}
