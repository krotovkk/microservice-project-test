package productservice

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

type Options struct {
	Store ports.ProductStore
	Cache ports.Cache
}

type ProductService struct {
	productStore ports.ProductStore
	cache        ports.Cache
}

func NewProductService(options *Options) *ProductService {
	return &ProductService{
		productStore: options.Store,
		cache:        options.Cache,
	}
}

func (ps *ProductService) CreateProduct(ctx context.Context, name string, price float64) (*model.Product, error) {
	product, err := model.NewProduct(0, name, price)
	if err != nil {
		return nil, err
	}

	err = ps.cache.Product().ClearGetAllProducts()

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "create product", "operation": "clear cache"}).Error(err)
	}

	return ps.productStore.CreateProduct(ctx, product)
}

func (ps *ProductService) UpdateProduct(ctx context.Context, name string, price float64, id uint) error {
	product, err := model.NewProduct(id, name, price)
	if err != nil {
		return err
	}

	err = ps.cache.Product().ClearGetAllProducts()

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "update product", "operation": "clear cache"}).Error(err)
	}

	err = ps.cache.Cart().ClearAllCartProducts()

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "update product", "operation": "clear cache"}).Error(err)
	}

	return ps.productStore.UpdateProduct(ctx, product)
}

func (ps *ProductService) DeleteProduct(ctx context.Context, id uint) error {
	err := ps.cache.Product().ClearGetAllProducts()

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "delete product", "operation": "clear cache"}).Error(err)
	}

	return ps.productStore.DeleteProduct(ctx, id)
}

func (ps *ProductService) GetAllProducts(ctx context.Context, limit uint64, offset uint64) ([]*model.Product, error) {
	cachedProducts, err := ps.cache.Product().GetAllProducts(ctx, limit, offset)

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "get all products", "operation": "get from cache"}).Error(err)
	} else {
		logrus.WithFields(logrus.Fields{"products": cachedProducts}).Infof("get products from cache successfully")
		return cachedProducts, nil
	}

	products, err := ps.productStore.GetAllProducts(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	err = ps.cache.Product().SetAllProducts(ctx, limit, offset, products)

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "get all products", "operation": "set to cache"}).Error(err)
	}

	return products, nil
}
