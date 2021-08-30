-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION postgres_fdw;
CREATE SERVER postgres FOREIGN DATA WRAPPER postgres_fdw
    OPTIONS (host 'postgres://postgres:postgres@127.0.0.1:5432', dbname 'postgres');
CREATE USER MAPPING FOR alice SERVER postgres
    OPTIONS (user 'postgres_shard_alice');
IMPORT FOREIGN SCHEMA public LIMIT TO (certificate)
    FROM SERVER postgres_shard INTO public;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
