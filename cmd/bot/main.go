package main

import (
	"log"

	"gitlab.ozon.dev/krotovkk/homework/internal/commander/botcommander"
	"gitlab.ozon.dev/krotovkk/homework/internal/handlers"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/productservice"
	"gitlab.ozon.dev/krotovkk/homework/internal/store/memorystore"
)

func main() {
	log.Println("Start application")

	store := memorystore.NewMemoryProductStore()
	service := productservice.NewProductService(store)

	runBot(service)
}

func runBot(service *productservice.ProductService) {
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
