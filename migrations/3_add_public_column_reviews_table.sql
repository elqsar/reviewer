-- +goose Up
ALTER TABLE reviews ADD is_public BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE reviews DROP COLUMN is_public;