-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
     id SERIAL PRIMARY KEY,
     email VARCHAR(255) NOT NULL UNIQUE,
     username VARCHAR(255) NOT NULL UNIQUE,
     password_hash BYTEA NOT NULL,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     is_admin BOOL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
