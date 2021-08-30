-- +goose Up
-- +goose StatementBegin
CREATE TABLE certificate_user_id_1_99
    PARTITION OF certificate (
    user_id NOT NULL,
    CHECK (user_id > 0 AND user_id <= 99)
)
    FOR VALUES FROM (0) TO (100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE certificate_user_id_1_99;
-- +goose StatementEnd
