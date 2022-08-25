package cartservice

import (
	"context"
	"time"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	"gitlab.ozon.dev/krotovkk/homework/internal/ports"
)

type CartService struct {
	cartStore ports.CartStore
}

func NewCartService(store ports.CartStore) *CartService {
	return &CartService{cartStore: store}
}

func (cs *CartService) CreateCart(ctx context.Context) (*model.Cart, error) {
	cart := &model.Cart{CreatedAt: time.Now().UTC().Unix()}

	return cs.cartStore.CreateCart(ctx, cart)
}

func (cs *CartService) GetCartProducts(ctx context.Context, id int64) ([]*model.Product, error) {
	return cs.cartStore.GetCartProducts(ctx, id)
}

func (cs *CartService) AddProductToCart(ctx context.Context, productId, cartId int64) error {
	return cs.cartStore.AddProductToCart(ctx, productId, cartId)
}
