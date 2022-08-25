package postgresstore

import (
	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

type PostgresStore struct {
	conn *pgx.Conn

	productStore ports.ProductStore
	cartStore    ports.CartStore
}

func NewPostgresStore(conn *pgx.Conn) *PostgresStore {
	store := &PostgresStore{conn: conn}

	store.productStore = NewPostgresProductStore(store)
	store.cartStore = NewCartPostgresStore(store)

	return store
}

func (s *PostgresStore) Product() ports.ProductStore {
	return s.productStore
}

func (s *PostgresStore) Cart() ports.CartStore {
	return s.cartStore
}

func (s *PostgresStore) Conn() *pgx.Conn {
	return s.conn
}
