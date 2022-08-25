package postgres

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/krotovkk/homework/internal/store/postgresstore"
	"gitlab.ozon.dev/krotovkk/homework/tests/config"
)

type TDB struct {
	sync.Mutex
	Store *postgresstore.PostgresStore
}

func NewFromEnv() *TDB {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Panic(err)
	}
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=%d sslmode=disable",
		cfg.DbHost, cfg.DbPort, "user", "password", "products", 30)

	conn, err := pgx.Connect(context.Background(), psqlConn)
	if err != nil {
		log.Panic(err)
	}
	return &TDB{Store: postgresstore.NewPostgresStore(conn)}
}

func (d *TDB) SetUp(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	d.Lock()
	d.Truncate(ctx)
}

func (d *TDB) TearDown() {
	defer d.Unlock()
	d.Truncate(context.Background())
}

func (d *TDB) Truncate(ctx context.Context) {
	var tables []string

	err := pgxscan.Select(ctx, d.Store.Conn(), &tables, "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name != 'goose_db_version'")
	if err != nil {
		log.Panic(err)
	}
	if len(tables) == 0 {
		log.Panic(err)
	}
	q := fmt.Sprintf("Truncate table %s", strings.Join(tables, ","))
	if _, err := d.Store.Conn().Exec(ctx, q); err != nil {
		log.Panic(err)
	}
}
