package datagrpccommander

import (
	"fmt"
	"log"
	"net"

	"gitlab.ozon.dev/krotovkk/homework/config"
	"gitlab.ozon.dev/krotovkk/homework/internal/api/dataapi"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"

	"google.golang.org/grpc"

	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func RunGrpcServer(service ports.Service) {
	address := fmt.Sprintf("localhost:%d", config.DataGrpcServerPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServer(grpcServer, dataapi.NewProductServer(service.Product()))
	pb.RegisterCartServer(grpcServer, dataapi.NewCartServer(service.Cart()))

	if err := grpcServer.Serve(listener); err != nil {
		log.Panic(err)
	}
}
