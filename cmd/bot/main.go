package main

import (
	"log"

	"gitlab.ozon.dev/krotovkk/homework/internal/commander/botcommander"
	"gitlab.ozon.dev/krotovkk/homework/internal/handlers"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/product_service"
	"gitlab.ozon.dev/krotovkk/homework/internal/store/memorystore"
)

func main() {
	log.Println("Start application")

	store := memorystore.NewMemoryProductStore()
	service := product_service.NewProductService(store)

	runBot(service)
}

func runBot(service *product_service.ProductService) {
	cmd, err := botcommander.Init()
	if err != nil {
		log.Panic(err)
	}

	handler := handlers.NewBotHandler(service)

	cmd.RegisterHandlers(handler)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
