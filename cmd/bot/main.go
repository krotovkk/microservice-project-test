package main

import (
	"log"

	"gitlab.ozon.dev/krotovkk/homework/internal/commander"
	"gitlab.ozon.dev/krotovkk/homework/internal/handlers"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/product_service"
	"gitlab.ozon.dev/krotovkk/homework/internal/store/memorystore"
)

func main() {
	log.Println("Start application")

	cmd, err := commander.Init()
	if err != nil {
		log.Panic(err)
	}

	store := memorystore.NewMemoryProductStore()
	service := product_service.NewProductService(store)
	handler := handlers.NewBotHandler(service)

	cmd.RegisterHandlers(handler)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
