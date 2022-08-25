package dataapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	"gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func TestProductServer_ProductCreate(t *testing.T) {
	t.Run("valid product creation", func(t *testing.T) {
		//arrange
		fixture := setUpProductFixture(t)
		defer fixture.tearDown()

		product := &model.Product{Id: 1, Name: "test", Price: 2100}

		fixture.service.EXPECT().CreateProduct(context.Background(), product.Name, product.Price).Return(product, nil)

		//act
		resp, err := fixture.server.ProductCreate(context.Background(), &api.ProductCreateRequest{Price: product.Price, Name: product.Name})

		//assert
		require.NoError(t, err)
		assert.Equal(t, resp, &api.ProductCreateResponse{Id: int64(product.Id), Name: product.Name, Price: product.Price})
	})
}

func TestProductServer_ProductList(t *testing.T) {
	t.Run("valid products return", func(t *testing.T) {
		//arrange
		//act
		//assert
	})
}

func TestProductServer_ProductUpdate(t *testing.T) {
	t.Run("valid product update", func(t *testing.T) {
		//arrange
		fixture := setUpProductFixture(t)
		defer fixture.tearDown()

		var price float64 = 200
		var name = "test name"
		var id uint = 1

		fixture.service.EXPECT().UpdateProduct(context.Background(), name, price, id).Return(nil)

		//act
		resp, err := fixture.server.ProductUpdate(context.Background(), &api.ProductUpdateRequest{Price: price, Name: name, Id: uint64(id)})

		//assert
		require.NoError(t, err)
		assert.Equal(t, resp, &api.ProductUpdateResponse{})
	})
}

func TestProductServer_ProductDelete(t *testing.T) {
	t.Run("valid product delete", func(t *testing.T) {
		//arrange
		fixture := setUpProductFixture(t)
		defer fixture.tearDown()

		var id uint = 1

		fixture.service.EXPECT().DeleteProduct(context.Background(), id).Return(nil)

		//act
		resp, err := fixture.server.ProductDelete(context.Background(), &api.ProductDeleteRequest{Id: uint64(id)})

		//assert
		require.NoError(t, err)
		assert.Equal(t, resp, &api.ProductDeleteResponse{})
	})
}
