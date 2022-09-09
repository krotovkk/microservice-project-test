package services

import (
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/cartservice"
	"gitlab.ozon.dev/krotovkk/homework/internal/services/productservice"
)

type Options struct {
	Store ports.Store
	Cache ports.Cache
}

type AppService struct {
	productService ports.ProductService
	cartService    ports.CartService
}

func NewAppService(options *Options) *AppService {
	productService := productservice.NewProductService(&productservice.Options{Store: options.Store.Product(), Cache: options.Cache})
	cartService := cartservice.NewCartService(&cartservice.Options{Store: options.Store.Cart(), Cache: options.Cache})

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
