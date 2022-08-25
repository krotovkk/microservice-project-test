//go:build integration
// +build integration

package tests

import (
	"log"

	"gitlab.ozon.dev/krotovkk/homework/tests/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gitlab.ozon.dev/krotovkk/homework/pkg/api"
	"gitlab.ozon.dev/krotovkk/homework/tests/config"
)

var (
	ProductClient api.ProductClient
	CartClient    api.CartClient
	Db            *postgres.TDB
)

func init() {
	cfg, err := config.FromEnv()

	conn, err := grpc.Dial(cfg.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	ProductClient = api.NewProductClient(conn)
	CartClient = api.NewCartClient(conn)

	Db = postgres.NewFromEnv()
}
