package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/broker"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gitlab.ozon.dev/krotovkk/homework/config"
	"gitlab.ozon.dev/krotovkk/homework/internal/api/validationapi"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func init() {
	logrus.SetOutput(os.Stdout)
}

func main() {
	address := fmt.Sprintf("localhost:%d", config.ValidationGrpcServerPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	target := fmt.Sprintf(":%d", config.DataGrpcServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logrus.WithError(err).Fatal()
	}

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(config.Brokers, cfg)
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	brokerService := broker.NewBrokerService(producer)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServer(grpcServer, validationapi.NewProductServer(pb.NewProductClient(conn), brokerService.Product()))
	pb.RegisterCartServer(grpcServer, validationapi.NewCartServer(pb.NewCartClient(conn), brokerService.Cart()))

	if err := grpcServer.Serve(listener); err != nil {
		logrus.WithError(err).Fatal()
	}
}
