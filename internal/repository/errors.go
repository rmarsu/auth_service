package repository

import "fmt"

var (
	ErrUserNotFound   = fmt.Errorf("user not found")
	ErrAppNotFound    = fmt.Errorf("app not found")
	ErrTokenNotFound  = fmt.Errorf("token not found")
	ErrEmailTaken     = fmt.Errorf("email already taken")
	ErrUsernameTaken  = fmt.Errorf("username already taken")
	ErrInvalidToken   = fmt.Errorf("invalid token")
	ErrExpiredToken   = fmt.Errorf("expired token")
	ErrWrongPassword  = fmt.Errorf("wrong password")
	ErrNoAppsFound    = fmt.Errorf("no apps found")
	ErrNoUsersFound   = fmt.Errorf("no users found")
	ErrNoTokensFound  = fmt.Errorf("no tokens found")
	ErrNoPermissions  = fmt.Errorf("insufficient permissions")
)
