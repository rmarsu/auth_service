package repository

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func ErrorCode(err error) *pgconn.PgError {
	var pgxErr *pgconn.PgError
	if errors.As(err, &pgxErr) {
		return pgxErr
	}
	return nil
}
