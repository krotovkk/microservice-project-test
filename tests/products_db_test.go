//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/krotovkk/homework/internal/model"
)

func TestCreateProduct_Db(t *testing.T) {
	errFailedInsert := errors.New("PostgresProductStore.CreateProduct: insert: ERROR: new row for relation \"products\" violates check constraint \"price_chk\" (SQLSTATE 23514)")

	tests := []struct {
		name string
		args model.Product
		err  error
	}{
		{
			name: "create model 1",
			args: model.Product{Name: "test1", Price: 210},
			err:  nil,
		},
		{
			name: "create model 2",
			args: model.Product{Name: "test2", Price: 210},
			err:  nil,
		},
		{
			name: "create model 3",
			args: model.Product{Name: "test3", Price: 210},
			err:  nil,
		},
		{
			name: "create model failed",
			args: model.Product{},
			err:  errFailedInsert,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//arrange
			Db.SetUp(t)
			defer Db.TearDown()
			//act
			product, err := Db.Store.Product().CreateProduct(context.Background(), &test.args)
			//assert
			if test.err != nil {
				assert.Equal(t, err.Error(), test.err.Error())
			} else {
				require.Nil(t, err)
				assert.Equal(t, test.args, *product)
			}
		})
	}
}

func TestProductUpdate_Db(t *testing.T) {
	t.Run("Product Update", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		product, err := Db.Store.Product().CreateProduct(context.Background(), &model.Product{Price: 222, Name: "test"})
		require.NoError(t, err)

		err = Db.Store.Product().UpdateProduct(context.Background(), product)
		require.NoError(t, err)

		product.Price = 2000
		product.Name = "updated"
		err = Db.Store.Product().UpdateProduct(context.Background(), product)
		require.NoError(t, err)
	})
}

func TestProductDelete_Db(t *testing.T) {
	t.Run("Product Delete", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		product, err := Db.Store.Product().CreateProduct(context.Background(), &model.Product{Price: 222, Name: "test"})
		require.NoError(t, err)

		err = Db.Store.Product().DeleteProduct(context.Background(), product.Id)
		require.NoError(t, err)
	})
}

func TestProductList_Db(t *testing.T) {
	t.Run("Product List", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		insertProducts := []*model.Product{
			{Price: 220, Name: "test1"},
			{Price: 100, Name: "test2"},
			{Price: 20000, Name: "test3"},
			{Price: 1111, Name: "test4"},
			{Price: 152, Name: "test5"},
		}

		for _, product := range insertProducts {
			_, err := Db.Store.Product().CreateProduct(context.Background(), product)
			require.NoError(t, err)
		}

		products, err := Db.Store.Product().GetAllProducts(context.Background(), uint64(len(insertProducts)), 0)
		require.NoError(t, err)
		assert.Equal(t, insertProducts, products)

		products, err = Db.Store.Product().GetAllProducts(context.Background(), 1, 0)
		require.NoError(t, err)
		assert.NotEqual(t, insertProducts, products)
	})
}
