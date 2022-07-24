package product_service

import (
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

func (ps *ProductService) Create(name string, price float64) error {
	product, err := model.NewProduct(0, name, price)
	if err != nil {
		return err
	}

	return ps.productStore.Add(product)
}

func (ps *ProductService) Update(id uint, name string, price float64) error {
	product, err := model.NewProduct(id, name, price)
	if err != nil {
		return err
	}

	return ps.productStore.Update(product)
}

func (ps *ProductService) Delete(id uint) error {
	return ps.productStore.Delete(id)
}

func (ps *ProductService) List() []*model.Product {
	return ps.productStore.List()
}
