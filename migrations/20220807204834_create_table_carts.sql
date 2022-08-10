-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS identity,
    created_at INTEGER NOT NULL CONSTRAINT  price_chk CHECK ( created_at > 0 )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carts;
-- +goose StatementEnd
