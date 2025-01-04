package delivery_grpc

import (
	"context"

	auth_service "github.com/rmarsu/auth_service/internal/proto"
	"github.com/rmarsu/auth_service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	NoValue = 0
)

type AuthHandlers struct {
	auth_service.UnimplementedAuthServiceServer
	services service.Services
}

func NewAuthHandlers(services service.Services) *AuthHandlers {
	return &AuthHandlers{
		services: services,
	}
}

func (a *AuthHandlers) Register(ctx context.Context,
	in *auth_service.RegisterRequest) (*auth_service.RegisterResponse, error) {

	if in.Username == "" {
		return nil, status.Error(codes.InvalidArgument, ErrUsernameIsRequired)
	}
	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, ErrPasswordIsRequired)
	}

	id, err := a.services.Auth.RegisterUser(ctx, in.GetUsername(), in.GetPassword())
	if err != nil {
		return nil, err
	}
	return &auth_service.RegisterResponse{Id: id}, nil
}

func (a *AuthHandlers) Login(ctx context.Context,
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
	token, err := a.services.Auth.Login(ctx, in.GetUsername(), in.GetPassword(), in.GetAppId())
	if err != nil {
		return nil, err
	}
	return &auth_service.LoginResponse{Token: token}, nil
}

func (a *AuthHandlers) IsAdmin(ctx context.Context,
	in *auth_service.IsAdminRequest) (*auth_service.IsAdminResponse, error) {

	if in.Token == "" {
		return nil, status.Error(codes.InvalidArgument, ErrUsernameIsRequired)
	}
	isAdmin, err := a.services.Auth.IsAdmin(ctx, in.GetToken())
	if err != nil {
		return nil, err
	}
	return &auth_service.IsAdminResponse{IsAdmin: isAdmin}, nil
}
