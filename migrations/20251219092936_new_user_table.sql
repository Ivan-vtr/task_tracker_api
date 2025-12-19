-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
       id BIGSERIAL PRIMARY KEY,
       email VARCHAR(100) NOT NULL UNIQUE,
       password_hash VARCHAR(255) NOT NULL,
       role VARCHAR(50) NOT NULL,
       created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
