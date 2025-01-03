package grpc

import (
	"context"

	"github.com/rmarsu/auth_service/internal/config"
	auth_service "github.com/rmarsu/auth_service/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	NoValue = 0
)

type AuthMicroservice struct {
	cfg  *config.Config
	auth Auth
}

func NewAuthMicroservice(cfg *config.Config) *AuthMicroservice {
	return &AuthMicroservice{cfg: cfg}
}

type Auth interface {
	RegisterUser(ctx context.Context, username, password string) (int64, error)
	Login(ctx context.Context, username, password string, appId int64) (string, error)
	IsAdmin(ctx context.Context, username string) (bool, error)
}

func (a *AuthMicroservice) Register(ctx context.Context,
	in *auth_service.RegisterRequest) (*auth_service.RegisterResponse, error) {

	if in.Username == "" {
		return nil, status.Error(codes.InvalidArgument, ErrUsernameIsRequired)
	}
	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, ErrPasswordIsRequired)
	}

	id, err := a.auth.RegisterUser(ctx, in.GetUsername(), in.GetPassword())
	if err != nil {
		return nil, err
	}
	return &auth_service.RegisterResponse{Id: id}, nil
}

func (a *AuthMicroservice) Login(ctx context.Context,
	in *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {

	if in.Username == "" {
		return nil, status.Error(codes.InvalidArgument, ErrUsernameIsRequired)
	}
	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, ErrPasswordIsRequired)
	}
	if in.AppId == NoValue {
		return nil, status.Error(codes.InvalidArgument, ErrAppIdIsRequired)
	}
	token, err := a.auth.Login(ctx, in.GetUsername(), in.GetPassword(), in.GetAppId())
	if err != nil {
		return nil, err
	}
	return &auth_service.LoginResponse{Token: token}, nil
}

func (a *AuthMicroservice) IsAdmin(ctx context.Context,
	in *auth_service.IsAdminRequest) (*auth_service.IsAdminResponse, error) {

	if in.Token == "" {
		return nil, status.Error(codes.InvalidArgument, ErrUsernameIsRequired)
	}
	isAdmin, err := a.auth.IsAdmin(ctx, in.GetToken())
	if err != nil {
		return nil, err
	}
	return &auth_service.IsAdminResponse{IsAdmin: isAdmin}, nil
}
