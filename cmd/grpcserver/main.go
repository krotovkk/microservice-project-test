package main

import (
	"gitlab.ozon.dev/krotovkk/homework/internal/commander/grpcservercommander"
	"gitlab.ozon.dev/krotovkk/homework/internal/commander/restcommander"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/product_service"
	"gitlab.ozon.dev/krotovkk/homework/internal/store/memorystore"
)

func main() {
	store := memorystore.NewMemoryProductStore()
	service := product_service.NewProductService(store)

	go grpcservercommander.RunGrpcServer(service)
	restcommander.Run()
}
