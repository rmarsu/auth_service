package delivery_grpc

import (
	"context"

	"github.com/rmarsu/auth_service/internal/domain"
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
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrEmailIsRequired)
	}
	if in.Username == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrUsernameIsRequired)
	}
	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrPasswordIsRequired)
	}

	id, err := a.services.Auth.RegisterUser(ctx, in.GetEmail(), in.GetUsername(), in.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, domain.ErrSomethingWentWrong)
	}
	return &auth_service.RegisterResponse{Id: id}, nil
}

func (a *AuthHandlers) Login(ctx context.Context,
	in *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {

	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrEmailIsRequired)
	}
	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrPasswordIsRequired)
	}
	if in.AppId == NoValue {
		return nil, status.Error(codes.InvalidArgument, domain.ErrAppIdIsRequired)
	}
	token, err := a.services.Auth.Login(ctx, in.GetEmail(), in.GetPassword(), in.GetAppId())
	if err != nil {
		return nil, status.Error(codes.Internal, domain.ErrSomethingWentWrong)
	}
	return &auth_service.LoginResponse{Token: token}, nil
}

func (a *AuthHandlers) IsAdmin(ctx context.Context,
	in *auth_service.IsAdminRequest) (*auth_service.IsAdminResponse, error) {

	if in.Id == NoValue {
		return nil, status.Error(codes.InvalidArgument, domain.ErrIdIsRequired)
	}
	isAdmin, err := a.services.Auth.IsAdmin(ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, domain.ErrSomethingWentWrong)
	}
	return &auth_service.IsAdminResponse{IsAdmin: isAdmin}, nil
}
