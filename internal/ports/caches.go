package ports

import (
	"context"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

type Cache interface {
	Product() ProductCache
	Cart() CartCache
}

type ProductCache interface {
	GetAllProducts(ctx context.Context, limit uint64, offset uint64) ([]*model.Product, error)
	SetAllProducts(ctx context.Context, limit uint64, offset uint64, products []*model.Product) error
	ClearGetAllProducts() error
	SendProductsToChanel(data []*model.Product) error
}

type CartCache interface {
	GetCartProducts(ctx context.Context, id int64) ([]*model.Product, error)
	SetCartProducts(ctx context.Context, id int64, products []*model.Product) error
	ClearGetCartProducts(id int64) error
	ClearAllCartProducts() error
	SendProductsToChanel(data []*model.Product) error
}
