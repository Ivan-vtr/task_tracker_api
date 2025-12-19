-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
                       id BIGSERIAL PRIMARY KEY,
                       user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       title VARCHAR(200) NOT NULL,
                       description TEXT,
                       status VARCHAR(50) NOT NULL DEFAULT 'pending',
                       created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Создаём триггер для автообновления updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON tasks
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks CASCADE;
DROP FUNCTION IF EXISTS update_updated_at_column();
-- +goose StatementEnd
