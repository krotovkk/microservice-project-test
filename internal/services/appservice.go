package services

import (
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/cartservice"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/productservice"
)

type AppService struct {
	productService ports.ProductService
	cartService    ports.CartService
}

func NewAppService(store ports.Store) *AppService {
	productService := productservice.NewProductService(store.Product())
	cartService := cartservice.NewCartService(store.Cart())

	return &AppService{
		productService: productService,
		cartService:    cartService,
	}
}

func (as *AppService) Product() ports.ProductService {
	return as.productService
}

func (as *AppService) Cart() ports.CartService {
	return as.cartService
}
