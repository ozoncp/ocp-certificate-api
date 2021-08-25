-- +goose Up
-- +goose StatementBegin
ALTER TABLE certificate ADD COLUMN is_deleted BOOLEAN NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE certificate DROP COLUMN is_deleted;
-- +goose StatementEnd
