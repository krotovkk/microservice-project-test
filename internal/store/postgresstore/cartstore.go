package postgresstore

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

type CartPostgresStore struct {
	*PostgresStore
}

func NewCartPostgresStore(store *PostgresStore) *CartPostgresStore {
	return &CartPostgresStore{PostgresStore: store}
}

func (cs *CartPostgresStore) CreateCart(ctx context.Context, cart *model.Cart) error {
	query, args, err := squirrel.Insert("carts").
		Columns("created_at").
		Values(cart.CreatedAt).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return fmt.Errorf("CartPostgresStore.CreateCart: to sql: %w", err)
	}

	_, err = cs.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("CartPostgresStore.CreateCart: insert: %w", err)
	}

	return nil
}

func (cs *CartPostgresStore) GetCartProducts(ctx context.Context, id int64) ([]*model.Product, error) {
	query, args, err := squirrel.Select("id", "name", "price").
		From("products").
		LeftJoin("carts_products ON products.id = carts_products.product_id").
		Where(squirrel.Eq{"carts_products.cart_id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return nil, fmt.Errorf("CartPostgresStore.GetCartProducts: to sql: %w", err)
	}

	var products []*model.Product

	err = pgxscan.Select(ctx, cs.conn, &products, query, args...)

	if err != nil {
		return nil, fmt.Errorf("CartPostgresStore.GetCartProducts: select: %w", err)
	}

	return products, nil
}

func (cs *CartPostgresStore) AddProductToCart(ctx context.Context, productId, cartId int64) error {
	query, args, err := squirrel.Insert("carts_products").
		Columns("cart_id").
		Columns("product_id").
		Values(cartId, productId).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return fmt.Errorf("CartPostgresStore.AddProductToCart: to sql: %w", err)
	}

	_, err = cs.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("CartPostgresStore.AddProductToCart: insert: %w", err)
	}

	return nil
}
