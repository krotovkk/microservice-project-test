package productservice

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

func TestProductService_CreateProduct(t *testing.T) {
	t.Run("valid creation", func(t *testing.T) {
		//arrange
		fixture := setUpFixture(t)
		defer fixture.tearDown()
		product := model.Product{Id: 0, Price: 200, Name: "test"}

		fixture.productStore.EXPECT().CreateProduct(context.Background(), &product).Return(&product, nil)
		//act
		_, err := fixture.productService.CreateProduct(context.Background(), product.Name, product.Price)
		//assert
		require.NoError(t, err)
	})
}

func TestProductService_DeleteProduct(t *testing.T) {
	t.Run("valid delete", func(t *testing.T) {
		//arrange
		fixture := setUpFixture(t)
		defer fixture.tearDown()
		var productId uint = 1

		fixture.productStore.EXPECT().DeleteProduct(context.Background(), productId).Return(nil)
		//act
		err := fixture.productService.DeleteProduct(context.Background(), productId)
		//assert
		require.NoError(t, err)
	})
}

func TestProductService_GetAllProducts(t *testing.T) {
	t.Run("valid list", func(t *testing.T) {
		//arrange
		fixture := setUpFixture(t)
		defer fixture.tearDown()
		var limit uint64 = 100
		var offset uint64 = 0

		fixture.productStore.EXPECT().GetAllProducts(context.Background(), gomock.Any(), gomock.Any()).Return([]*model.Product{}, nil)
		//act
		products, err := fixture.productService.GetAllProducts(context.Background(), limit, offset)
		//assert
		require.NoError(t, err)
		assert.Equal(t, products, []*model.Product{})
	})
}

func TestProductService_UpdateProduct(t *testing.T) {
	t.Run("valid update", func(t *testing.T) {
		//arrange
		fixture := setUpFixture(t)
		defer fixture.tearDown()
		product := model.Product{Id: 1, Price: 200, Name: "test"}

		fixture.productStore.EXPECT().UpdateProduct(context.Background(), &product).Return(nil)
		//act
		err := fixture.productService.UpdateProduct(context.Background(), product.Name, product.Price, product.Id)
		//assert
		require.NoError(t, err)
	})
}
