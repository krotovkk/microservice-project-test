-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts_products (
    cart_id INTEGER REFERENCES carts (id),
    product_id INTEGER REFERENCES products (id),
    CONSTRAINT carts_products_pk PRIMARY KEY (cart_id, product_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carts_products;
-- +goose StatementEnd
