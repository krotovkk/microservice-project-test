package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/krotovkk/homework/internal/commander/brokercommander"
	"log"
	"os"

	"github.com/jackc/pgx/v4"

	"gitlab.ozon.dev/krotovkk/homework/config"
	"gitlab.ozon.dev/krotovkk/homework/internal/commander/datagrpccommander"
	"gitlab.ozon.dev/krotovkk/homework/internal/services"
	"gitlab.ozon.dev/krotovkk/homework/internal/store/postgresstore"
)

func init() {
	logrus.SetOutput(os.Stdout)
}

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

	grpcCh := make(chan struct{})
	brokerCh := make(chan struct{})

	go datagrpccommander.RunGrpcServer(service, grpcCh)
	go brokercommander.Run(service, brokerCh)
	<-grpcCh
	<-brokerCh
}
