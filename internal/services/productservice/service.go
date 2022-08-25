package productservice

import (
	"context"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

type ProductService struct {
	productStore ports.ProductStore
}

func NewProductService(store ports.ProductStore) *ProductService {
	return &ProductService{
		productStore: store,
	}
}

func (ps *ProductService) CreateProduct(ctx context.Context, name string, price float64) (*model.Product, error) {
	product, err := model.NewProduct(0, name, price)
	if err != nil {
		return nil, err
	}

	return ps.productStore.CreateProduct(ctx, product)
}

func (ps *ProductService) UpdateProduct(ctx context.Context, name string, price float64, id uint) error {
	product, err := model.NewProduct(id, name, price)
	if err != nil {
		return err
	}

	return ps.productStore.UpdateProduct(ctx, product)
}

func (ps *ProductService) DeleteProduct(ctx context.Context, id uint) error {
	return ps.productStore.DeleteProduct(ctx, id)
}

func (ps *ProductService) GetAllProducts(ctx context.Context, limit uint64, offset uint64) ([]*model.Product, error) {
	return ps.productStore.GetAllProducts(ctx, limit, offset)
}
