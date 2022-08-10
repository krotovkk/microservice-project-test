package ports

import (
	"context"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

type Store interface {
	Product() ProductStore
	Cart() CartStore
}

type ProductStore interface {
	GetAllProducts(ctx context.Context, limit uint64, offset uint64) ([]*model.Product, error)
	CreateProduct(ctx context.Context, p *model.Product) error
	DeleteProduct(ctx context.Context, id uint) error
	UpdateProduct(ctx context.Context, p *model.Product) error
	GetProductOne(ctx context.Context, id int64) (*model.Product, error)
}

type CartStore interface {
	CreateCart(ctx context.Context, c *model.Cart) error
	GetCartProducts(ctx context.Context, id int64) ([]*model.Product, error)
	AddProductToCart(ctx context.Context, productId, cartId int64) error
}
