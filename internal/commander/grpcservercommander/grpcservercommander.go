package grpcservercommander

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"gitlab.ozon.dev/krotovkk/homework/internal/api"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/product_service"
	pb "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func RunGrpcServer(service *product_service.ProductService) {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, api.NewServer(service))

	if err := grpcServer.Serve(listener); err != nil {
		log.Panic(err)
	}
}
