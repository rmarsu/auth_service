package service

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrAppNotFound           = errors.New("app not found")
	ErrPasswordIsNotValid    = errors.New("password is not valid")
	ErrWrongPassword         = errors.New("wrong password")
	ErrSomethingWentWrong    = errors.New("something went wrong")
	ErrIdAlreadyExists       = errors.New("id already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrTokenExpired          = errors.New("token expired")
	ErrTokenInvalid          = errors.New("invalid token")
	ErrInvalidApp            = errors.New("invalid app")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrEmailNotConfirmed     = errors.New("email not confirmed")
	ErrEmailNotVerified      = errors.New("email not verified")
	ErrInvalidToken          = errors.New("invalid token")
	ErrInvalidUsername       = errors.New("invalid username")
	ErrInvalidEmail          = errors.New("invalid email")
	ErrInvalidPassword       = errors.New("invalid password")
)
