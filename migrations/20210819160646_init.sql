-- +goose Up
-- +goose StatementBegin
CREATE TABLE certificate
(
    id         SERIAL PRIMARY KEY,
    user_id    INT,
    created    TIMESTAMP(0) WITH TIME ZONE,
    link       VARCHAR(255)
)
PARTITION BY RANGE (id);
CREATE INDEX idx_certificate_user_id_created ON certificate (user_id, created);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE certificate;
-- +goose StatementEnd