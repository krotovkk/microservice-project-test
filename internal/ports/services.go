//go:generate mockgen -source=./services.go -destination=../services/mocks/services.go -package=mock_services

package ports

import (
	"context"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

type Service interface {
	Product() ProductService
	Cart() CartService
}

type ProductService interface {
	GetAllProducts(ctx context.Context, limit uint64, offset uint64) ([]*model.Product, error)
	CreateProduct(ctx context.Context, name string, price float64) (*model.Product, error)
	UpdateProduct(ctx context.Context, name string, price float64, id uint) error
	DeleteProduct(ctx context.Context, id uint) error
}

type CartService interface {
	CreateCart(ctx context.Context) (*model.Cart, error)
	GetCartProducts(ctx context.Context, id int64) ([]*model.Product, error)
	AddProductToCart(ctx context.Context, productId, cartId int64) error
}
