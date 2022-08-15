package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gitlab.ozon.dev/krotovkk/homework/config"
	"gitlab.ozon.dev/krotovkk/homework/internal/api/validationapi"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func main() {
	address := fmt.Sprintf("localhost:%d", config.ValidationGrpcServerPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	target := fmt.Sprintf(":%d", config.DataGrpcServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServer(grpcServer, validationapi.NewProductServer(pb.NewProductClient(conn)))
	pb.RegisterCartServer(grpcServer, validationapi.NewCartServer(pb.NewCartClient(conn)))

	if err := grpcServer.Serve(listener); err != nil {
		log.Panic(err)
	}
}
