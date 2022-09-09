package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/krotovkk/homework/internal/cache/rediscache"
	"gitlab.ozon.dev/krotovkk/homework/internal/commander/brokercommander"

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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		DB:       config.RedisDB,
		Password: config.RedisPassword,
	})

	redisPing := redisClient.Ping()

	if redisPing.Err() != nil {
		logrus.WithFields(logrus.Fields{"port": config.RedisPort, "host": config.RedisHost, "error": redisPing.Err()}).Fatal()
	}

	conn, err := pgx.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	store := postgresstore.NewPostgresStore(conn)
	cache := rediscache.NewRedisCache(redisClient)

	service := services.NewAppService(&services.Options{Store: store, Cache: cache})

	grpcCh := make(chan struct{})
	brokerCh := make(chan struct{})

	go datagrpccommander.RunGrpcServer(service, grpcCh)
	go brokercommander.Run(service, brokerCh)
	<-grpcCh
	<-brokerCh
}
