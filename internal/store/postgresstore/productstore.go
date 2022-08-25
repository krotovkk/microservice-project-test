package postgresstore

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

type PostgresProductStore struct {
	*PostgresStore
}

func NewPostgresProductStore(store *PostgresStore) *PostgresProductStore {
	return &PostgresProductStore{PostgresStore: store}
}

func (ps *PostgresProductStore) GetAllProducts(ctx context.Context, limit uint64, offset uint64) ([]*model.Product, error) {
	query, args, err := squirrel.Select("id", "name", "price").
		From("products").
		OrderBy("id").
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return nil, fmt.Errorf("PostgresProductStore.GetAllProducts: to sql: %w", err)
	}

	var products []*model.Product
	err = pgxscan.Select(ctx, ps.conn, &products, query, args...)

	if err != nil {
		return nil, fmt.Errorf("PostgresProductStore.GetAllProducts: select: %w", err)
	}

	return products, nil
}

func (ps *PostgresProductStore) CreateProduct(ctx context.Context, p *model.Product) (*model.Product, error) {
	query, args, err := squirrel.Insert("products").
		Columns("name", "price").
		Values(p.GetName(), p.GetPrice()).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return nil, fmt.Errorf("PostgresProductStore.CreateProduct: to sql: %w", err)
	}

	err = ps.conn.QueryRow(ctx, query, args...).Scan(&p.Id)
	if err != nil {
		return nil, fmt.Errorf("PostgresProductStore.CreateProduct: insert: %w", err)
	}

	return p, nil
}

func (ps *PostgresProductStore) DeleteProduct(ctx context.Context, id uint) error {
	query, args, err := squirrel.Delete("products").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return fmt.Errorf("PostgresProductStore.DeleteProduct: to sql: %w", err)
	}

	_, err = ps.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("PostgresProductStore.DeleteProduct: delete: %w", err)
	}

	return nil
}

func (ps *PostgresProductStore) UpdateProduct(ctx context.Context, p *model.Product) error {
	query, args, err := squirrel.Update("products").
		Set("name", p.GetName()).
		Set("price", p.GetPrice()).
		Where(squirrel.Eq{"id": p.GetId()}).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return fmt.Errorf("PostgresProductStore.UpdateProduct: to sql: %w", err)
	}

	_, err = ps.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("PostgresProductStore.UpdateProduct: update: %w", err)
	}

	return nil
}

func (ps *PostgresProductStore) GetProductOne(ctx context.Context, id int64) (*model.Product, error) {
	query, args, err := squirrel.Select("id", "name", "price").
		From("products").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return nil, fmt.Errorf("PostgresProductStore.GetProductOne: to sql: %w", err)
	}

	var products []*model.Product
	err = pgxscan.Select(ctx, ps.conn, &products, query, args...)

	if err != nil {
		return nil, fmt.Errorf("PostgresProductStore.GetProductOne: select: %w", err)
	}

	if len(products) != 1 {
		return nil, fmt.Errorf("PostgresProductStore.GetProductOne: count of products more than one")
	}

	return products[0], nil
}
