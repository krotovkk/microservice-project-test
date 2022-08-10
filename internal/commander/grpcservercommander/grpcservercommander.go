package grpcservercommander

import (
	"gitlab.ozon.dev/krotovkk/homework/internal/services"
	"log"
	"net"

	"google.golang.org/grpc"

	"gitlab.ozon.dev/krotovkk/homework/internal/api"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func RunGrpcServer(service *services.AppService) {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServer(grpcServer, api.NewProductServer(service.Product()))
	pb.RegisterCartServer(grpcServer, api.NewCartServer(service.Cart()))

	if err := grpcServer.Serve(listener); err != nil {
		log.Panic(err)
	}
}
