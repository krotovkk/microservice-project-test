package cartservice

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

func TestCartService_CreateCart(t *testing.T) {
	t.Run("Valid creation", func(t *testing.T) {
		//arrange
		fixture := setUpFixture(t)
		defer fixture.tearDown()
		cart := &model.Cart{CreatedAt: time.Now().UTC().Unix()}
		retCart := &model.Cart{Id: 1, CreatedAt: time.Now().UTC().Unix()}

		fixture.cartStore.EXPECT().CreateCart(context.Background(), cart).Return(retCart, nil)
		//act
		_, err := fixture.cartService.CreateCart(context.Background())

		//assert
		require.NoError(t, err)
	})
}

func TestCartService_AddProductToCart(t *testing.T) {
	t.Run("Valid addition", func(t *testing.T) {
		//arrange
		fixture := setUpFixture(t)
		defer fixture.tearDown()
		var productId int64 = 1
		var cartId int64 = 2

		fixture.cartStore.EXPECT().AddProductToCart(context.Background(), productId, cartId).Return(nil)
		//act
		err := fixture.cartService.AddProductToCart(context.Background(), productId, cartId)

		//assert
		require.NoError(t, err)
	})
}

func TestCartService_GetCartProducts(t *testing.T) {
	t.Run("Valid receiving", func(t *testing.T) {
		//arrange
		fixture := setUpFixture(t)
		defer fixture.tearDown()
		var cartId int64 = 1

		fixture.cartStore.EXPECT().GetCartProducts(context.Background(), cartId).Return([]*model.Product{}, nil)
		//act
		products, err := fixture.cartService.GetCartProducts(context.Background(), cartId)

		//assert
		require.NoError(t, err)
		require.Equal(t, products, []*model.Product{})
	})
}
