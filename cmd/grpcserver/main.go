package main

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/krotovkk/homework/internal/services"
	"log"

	"github.com/jackc/pgx/v4"

	"gitlab.ozon.dev/krotovkk/homework/config"
	"gitlab.ozon.dev/krotovkk/homework/internal/commander/grpcservercommander"
	"gitlab.ozon.dev/krotovkk/homework/internal/commander/restcommander"
	"gitlab.ozon.dev/krotovkk/homework/internal/store/postgresstore"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=%d sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName, config.ConnectTimeout)

	conn, err := pgx.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	store := postgresstore.NewPostgresStore(conn)
	service := services.NewAppService(store)

	go grpcservercommander.RunGrpcServer(service)
	restcommander.Run()
}
