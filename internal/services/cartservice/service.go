package cartservice

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

type Options struct {
	Store ports.CartStore
	Cache ports.Cache
}

type CartService struct {
	cartStore ports.CartStore
	cache     ports.Cache
}

func NewCartService(store *Options) *CartService {
	return &CartService{
		cartStore: store.Store,
		cache:     store.Cache,
	}
}

func (cs *CartService) CreateCart(ctx context.Context) (*model.Cart, error) {
	cart := &model.Cart{CreatedAt: time.Now().UTC().Unix()}

	return cs.cartStore.CreateCart(ctx, cart)
}

func (cs *CartService) GetCartProducts(ctx context.Context, id int64) ([]*model.Product, error) {
	cachedProducts, err := cs.cache.Cart().GetCartProducts(ctx, id)

	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "get cart products", "operation": "get from cache"}).Error(err)
	} else {
		logrus.WithFields(logrus.Fields{"products": cachedProducts}).Infof("get cart products from cache successfully")
		return cachedProducts, nil
	}

	products, err := cs.cartStore.GetCartProducts(ctx, id)

	if err != nil {
		return nil, err
	}

	err = cs.cache.Cart().SetCartProducts(ctx, id, products)
	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "get cart products", "operation": "set to cache"}).Error(err)
	}

	return products, nil
}

func (cs *CartService) AddProductToCart(ctx context.Context, productId, cartId int64) error {
	err := cs.cache.Cart().ClearGetCartProducts(cartId)
	if err != nil {
		logrus.WithFields(logrus.Fields{"from": "add product to cart", "product": productId, "cartId": cartId, "operation": "clear cache"}).Error(err)
	}

	return cs.cartStore.AddProductToCart(ctx, productId, cartId)
}
