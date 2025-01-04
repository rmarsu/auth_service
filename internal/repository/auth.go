package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rmarsu/auth_service/internal/domain"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(ctx context.Context, email, username string, passHash []byte) (int64, error) {
	query := `INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3) RETURNING id`
	row := r.db.QueryRow(ctx, query, email, username, passHash)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepo) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	query := `SELECT id, email, username, password_hash, created_at FROM users WHERE email = $1`
	var user domain.User
	err := r.db.QueryRow(ctx, query, email).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.CreatedAt)
	switch err {
	case nil:
		return user, nil
	case pgx.ErrNoRows:
		return domain.User{}, ErrUserNotFound
	default:
		return domain.User{}, err
	}
}

func (r *AuthRepo) GetAppById(ctx context.Context, id int64) (domain.App, error) {
	query := `SELECT id, name, secret FROM apps WHERE id = $1`
	var app domain.App
	err := r.db.QueryRow(ctx, query, id).Scan(&app.Id, &app.Name, &app.Secret)
	switch err {
	case nil:
		return app, nil
	case pgx.ErrNoRows:
		return domain.App{}, ErrAppNotFound
	default:
		return domain.App{}, err
	}
}

func (r *AuthRepo) IsAdmin(ctx context.Context, userId int64) (bool, error) {
	query := `SELECT is_admin FROM users WHERE id = $1`
	var isAdmin bool
	err := r.db.QueryRow(ctx, query, userId).Scan(&isAdmin)
	switch err {
	case nil:
		return isAdmin, nil
	case pgx.ErrNoRows:
		return false, ErrUserNotFound
	default:
		return false, err
	}
}
