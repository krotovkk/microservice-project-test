-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS identity,
    name TEXT NOT NULL CONSTRAINT name_chk CHECK ( char_length(name) <= 30 ),
    price FLOAT NOT NULL CONSTRAINT  price_chk CHECK ( price > 0 )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
