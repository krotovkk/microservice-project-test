//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func TestCartCreate(t *testing.T) {
	t.Run("Cart Create", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		c, err := CartClient.CartCreate(context.Background(), &api.CartCreateRequest{})
		assert.NoError(t, err)
		assert.NotNil(t, c)
	})
}

func TestCartGetProducts(t *testing.T) {
	t.Run("Cart Add And Get Products", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		c, err := CartClient.CartCreate(context.Background(), &api.CartCreateRequest{})
		require.NoError(t, err)
		require.NotNil(t, c)

		p, err := ProductClient.ProductCreate(context.Background(), &api.ProductCreateRequest{Name: "test", Price: 900})
		require.NoError(t, err)
		require.NotNil(t, p)

		resp, err := CartClient.CartAddProduct(context.Background(), &api.CartAddProductRequest{CartId: c.Id, ProductId: p.Id})
		require.NoError(t, err)
		assert.NotNil(t, resp)

		getResp, err := CartClient.CartGetProducts(context.Background(), &api.CartGetProductsRequest{Id: c.Id})
		require.NoError(t, err)
		assert.Equal(t, 1, len(getResp.Products))
		assert.Equal(t, uint64(p.Id), getResp.Products[0].Id)
		assert.Equal(t, p.Price, getResp.Products[0].Price)
		assert.Equal(t, p.Name, getResp.Products[0].Name)
	})
}
