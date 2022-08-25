package dataapi

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	"gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func TestCartServer_CartAddProduct(t *testing.T) {
	t.Run("Valid cart product", func(t *testing.T) {
		//arrange
		fixture := setUpCartFixture(t)
		defer fixture.tearDown()

		var productId int64 = 1
		var cartId int64 = 1

		fixture.service.EXPECT().AddProductToCart(context.Background(), productId, cartId).Return(nil)

		//act
		resp, err := fixture.server.CartAddProduct(context.Background(), &api.CartAddProductRequest{ProductId: productId, CartId: cartId})

		//assert
		require.NoError(t, err)
		assert.Equal(t, resp, &api.CartAddProductResponse{})
	})

	t.Run("Not valid cart product", func(t *testing.T) {
		//arrange
		fixture := setUpCartFixture(t)
		defer fixture.tearDown()

		var productId int64 = 2
		var cartId int64 = 2
		returnErr := errors.New("not valid product")

		fixture.service.EXPECT().AddProductToCart(context.Background(), productId, cartId).Return(returnErr)

		//act
		resp, err := fixture.server.CartAddProduct(context.Background(), &api.CartAddProductRequest{ProductId: productId, CartId: cartId})

		//assert
		assert.Error(t, err, returnErr)
		assert.Nil(t, resp)
	})
}

func TestCartServer_CartCreate(t *testing.T) {
	t.Run("Valid cart creation", func(t *testing.T) {
		//arrange
		fixture := setUpCartFixture(t)
		defer fixture.tearDown()

		cart := &model.Cart{Id: 1, CreatedAt: time.Now().Unix()}

		fixture.service.EXPECT().CreateCart(context.Background()).Return(cart, nil)

		//act
		resp, err := fixture.server.CartCreate(context.Background(), &api.CartCreateRequest{})

		//assert
		require.NoError(t, err)
		assert.Equal(t, resp, &api.CartCreateResponse{Id: cart.Id, CreatedAt: cart.CreatedAt})
	})
}

func TestCartServer_CartGetProducts(t *testing.T) {
	t.Run("Valid products return", func(t *testing.T) {
		//arrange
		fixture := setUpCartFixture(t)
		defer fixture.tearDown()
		var id int64 = 1
		products := []*model.Product{
			{Id: 1, Price: 200, Name: "test1"},
			{Id: 26, Price: 440, Name: "test2"},
			{Id: 33, Price: 300, Name: "test3"},
			{Id: 4, Price: 210, Name: "test4"},
		}

		fixture.service.EXPECT().GetCartProducts(context.Background(), id).
			Return(products, nil)

		//act
		resp, err := fixture.server.CartGetProducts(context.Background(), &api.CartGetProductsRequest{Id: id})

		//assert
		require.NoError(t, err)
		for i, product := range products {
			assert.Equal(t, uint(resp.Products[i].Id), product.Id)
			assert.Equal(t, resp.Products[i].Price, product.Price)
			assert.Equal(t, resp.Products[i].Name, product.Name)
		}
	})
}
