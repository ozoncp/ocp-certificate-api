-- +goose Up
-- +goose StatementBegin
CREATE TABLE certificate_user_id_100_199
    PARTITION OF certificate (
    user_id NOT NULL,
    CHECK (user_id > 99 AND user_id <= 199)
)
    FOR VALUES FROM (100) TO (200);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE certificate_user_id_100_199;
-- +goose StatementEnd
