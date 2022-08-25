//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
	"gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func TestProductCreate(t *testing.T) {
	tests := []struct {
		name string
		args model.Product
		err  error
	}{
		{
			name: "create model 1",
			args: model.Product{Name: "test1", Price: 2000},
			err:  nil,
		},
		{
			name: "create model 2",
			args: model.Product{Name: "test2", Price: 20000},
			err:  nil,
		},
		{
			name: "create model 3",
			args: model.Product{Name: "test3", Price: 11000},
			err:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//arrange
			Db.SetUp(t)
			defer Db.TearDown()
			//act
			resp, err := ProductClient.ProductCreate(context.Background(), &api.ProductCreateRequest{Name: test.args.Name, Price: test.args.Price})
			//assert
			assert.Equal(t, err, test.err)
			assert.NotNil(t, resp)
		})
	}
}

func TestProductsList(t *testing.T) {
	t.Run("Products List", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()
		insertProducts := []*model.Product{
			{Price: 220, Name: "test1"},
			{Price: 100, Name: "test2"},
			{Price: 20000, Name: "test3"},
			{Price: 1111, Name: "test4"},
			{Price: 152, Name: "test5"},
		}

		resp, err := ProductClient.ProductList(context.Background(), &api.ProductListRequest{Limit: 1, Offset: 0})
		require.NoError(t, err)
		assert.NotNil(t, resp)

		for _, p := range insertProducts {
			_, err := ProductClient.ProductCreate(context.Background(), &api.ProductCreateRequest{Name: p.Name, Price: p.Price})
			assert.NoError(t, err)
		}

		resp, err = ProductClient.ProductList(context.Background(), &api.ProductListRequest{Limit: 1, Offset: 0})
		require.NoError(t, err)
	})
}

func TestProductsDelete(t *testing.T) {
	t.Run("Product Delete", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		p, err := ProductClient.ProductCreate(context.Background(), &api.ProductCreateRequest{Name: "test", Price: 1222})
		assert.NoError(t, err)

		resp, err := ProductClient.ProductDelete(context.Background(), &api.ProductDeleteRequest{Id: uint64(p.Id)})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
}

func TestProductsUpdate(t *testing.T) {
	t.Run("Product Update", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		p, err := ProductClient.ProductCreate(context.Background(), &api.ProductCreateRequest{Name: "test", Price: 1222})
		assert.NoError(t, err)

		resp, err := ProductClient.ProductUpdate(context.Background(), &api.ProductUpdateRequest{Id: uint64(p.Id), Name: "updated", Price: p.Price})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
}
